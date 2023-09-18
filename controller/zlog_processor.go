package controller

import (
	"ZLog/models"
	"ZLog/utils"
	"ZLog/zlog"
	"archive/zip"
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// ZLogProcessorInstance 是一个全局的 ZLogProcessor 实例。
var ZLogProcessorInstance *zLogProcessor

// ZLogProcessor 结构体
type zLogProcessor struct {
	wg        sync.WaitGroup
	baseDir   string
	taskQueue chan string
	stopChan  chan struct{}
}

// InitZLogProcessor 初始化 ZLogProcessor 实例
func InitZLogProcessor(baseDir string) {
	ZLogProcessorInstance = &zLogProcessor{
		baseDir:   baseDir,
		taskQueue: make(chan string, 50),
		stopChan:  make(chan struct{}),
	}
}

// AddTask 添加任务到任务队列
func (p *zLogProcessor) AddTask(zipFilePath string) {
	p.taskQueue <- zipFilePath
}

// Start 启动任务处理协程
func (p *zLogProcessor) Start() {
	p.wg.Add(1)
	go p.processTasks()
}

// Stop 停止任务处理协程
func (p *zLogProcessor) Stop() {
	close(p.taskQueue)
	p.wg.Wait()
	close(p.stopChan)
}

func (p *zLogProcessor) processTasks() {
	defer p.wg.Done()
	defer close(p.stopChan) //关闭 stopChan

	for {
		select {
		case zipFilePath, ok := <-p.taskQueue:
			if !ok {
				//任务队列已关闭
				return
			}

			if err := p.processZipAndLogFiles(zipFilePath); err != nil {
				fmt.Println(err)
			}
		case <-p.stopChan:
			//收到停止信号
			return
		}
	}
}

func (p *zLogProcessor) processZipAndLogFiles(zipFilePath string) error {
	//文件名为任务Id
	taskId := strings.TrimSuffix(filepath.Base(zipFilePath), filepath.Ext(zipFilePath))

	//根据任务Id获取对应的SessionId
	sessionId, err := models.GetSessionId(taskId)
	if err != nil {
		return err
	}

	//根据Session获取密钥对
	keyPair, err := models.GetKeyPairBySessionId(sessionId)
	if err != nil {
		return err
	}

	zipFile, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return err
	}

	logFilesDir := filepath.Join(p.baseDir, taskId+"_log_files")
	//先清空目录
	_ = utils.DeleteDirectory(logFilesDir)
	//再创建新的目录
	if err := os.MkdirAll(logFilesDir, os.ModePerm); err != nil {
		return err
	}

	//解压ZIP文件
	if err := p.extractAndSaveFiles(zipFile, logFilesDir); err != nil {
		return err
	}

	//关闭ZIP文件
	_ = zipFile.Close()

	//解析日志文件
	if err := p.processLogFiles(taskId, logFilesDir, keyPair); err != nil {
		return err
	}

	//删除ZIP文件
	if err := os.Remove(zipFilePath); err != nil {
		return err
	}

	//删除解压后的目录及其内容
	_ = utils.DeleteDirectory(logFilesDir)

	return nil
}

func (p *zLogProcessor) extractAndSaveFiles(zipFile *zip.ReadCloser, logFilesDir string) error {
	for _, file := range zipFile.File {
		entryReader, err := file.Open()
		if err != nil {
			return err
		}

		logFileName := filepath.Join(logFilesDir, file.Name)
		logFile, err := os.Create(logFileName)
		if err != nil {
			return err
		}

		_, err = io.Copy(logFile, entryReader)
		if err != nil {
			return err
		}

		_ = entryReader.Close()
		_ = logFile.Close()
	}

	return nil
}

func (p *zLogProcessor) processLogFiles(taskId string, logFilesDir string, keyPair utils.KeyPair) error {
	fileList, err := ioutil.ReadDir(logFilesDir)
	if err != nil {
		return err
	}

	for _, fileInfo := range fileList {
		filePath := filepath.Join(logFilesDir, fileInfo.Name())
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}

		if err := p.processLogFile(taskId, file, keyPair); err != nil {
			return err
		}

		_ = file.Close()
	}

	return nil
}

func (p *zLogProcessor) processLogFile(taskId string, logFile *os.File, keyPair utils.KeyPair) error {
	const maxBatchSize = 10 //定义最大批量大小

	var logData []*models.OfflineLog //用于存储要写入数据库的数据

	reader := bufio.NewReader(logFile)
	for {
		//读取文件的前6个字节
		headerData := make([]byte, 6)
		_, err := io.ReadFull(reader, headerData)
		if err != nil {
			if err == io.EOF {
				//如果到达文件末尾，退出循环
				break
			}
			return err //返回非EOF错误
		}

		encryptionFlag := headerData[0]
		compressionFlag := headerData[1]
		dataLength := binary.BigEndian.Uint32(headerData[2:6])

		//读取后续数据块
		dataBlock := make([]byte, dataLength)
		_, err = io.ReadFull(reader, dataBlock)
		if err != nil {
			return err //返回读取数据错误
		}

		//根据加密标志进行解密
		if encryptionFlag == 1 && len(keyPair.SharedKey) > 0 {
			decryptedBytes, err := utils.DecryptBytes(dataBlock, keyPair.SharedKey)
			if err != nil {
				continue //继续下一个循环，丢弃原始数据
			}
			dataBlock = decryptedBytes
		}

		//根据压缩标志进行解压缩
		if compressionFlag == 1 {
			decompressedBytes, err := utils.DecompressBytes(dataBlock)
			if err != nil {
				continue //继续下一个循环，丢弃原始数据
			}
			dataBlock = decompressedBytes
		}

		log := &zlog.Log{}
		err = proto.Unmarshal(dataBlock, log)

		if err == nil {
			//添加到待写入数据
			logData = append(logData, &models.OfflineLog{
				TaskId:        taskId,
				Sequence:      log.Sequence,
				SystemVersion: log.SystemVersion,
				AppVersion:    log.AppVersion,
				TimeStamp:     log.Timestamp,
				LogLevel:      int(log.LogLevel),
				Identify:      log.Identify,
				Tag:           log.Tag,
				Msg:           log.Msg,
			})
		}

		//检查是否达到最大批量大小，如果达到则写入数据库并清空切片
		if len(logData) >= maxBatchSize {
			p.writeToDatabase(logData)
			logData = nil
		}
	}

	//处理剩余的数据
	if len(logData) > 0 {
		p.writeToDatabase(logData)
	}

	return nil
}

func (p *zLogProcessor) writeToDatabase(log []*models.OfflineLog) {
	_ = models.WriteOfflineLogs(log)
}
