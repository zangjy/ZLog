package utils

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

// WorkerInstance 是一个全局的 Worker 实例，用于生成唯一 ID。
var (
	WorkerInstance, _ = NewWorker(0)
)

// 常量定义
const (
	workerBits  uint8 = 10                      // 每台机器(节点)的ID位数，最多1024个节点
	numberBits  uint8 = 12                      // 每毫秒可生成的id序号的二进制位数，最多4095个唯一ID
	workerMax   int64 = -1 ^ (-1 << workerBits) // 节点ID的最大值，用于防止溢出
	numberMax   int64 = -1 ^ (-1 << numberBits) // 用来表示生成id序号的最大值
	timeShift   uint8 = workerBits + numberBits // 时间戳向左的偏移量
	workerShift uint8 = numberBits              // 节点ID向左的偏移量
	epoch       int64 = 1525705533000           // 起始时间戳 (毫秒) - 2018年5月7日
)

// Worker 包含了生成唯一ID所需的基本参数
type Worker struct {
	mu        sync.Mutex // 互斥锁，确保并发安全
	timestamp int64      // 记录时间戳
	workerId  int64      // 节点的唯一ID
	number    int64      // 当前毫秒已生成的id序列号，每毫秒最多生成4096个唯一ID
}

// NewWorker 创建一个Worker实例，需要传入节点的唯一ID。
// 节点的ID必须在0到1023之间，否则会返回错误。
func NewWorker(workerId int64) (*Worker, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("Worker ID超出范围")
	}
	return &Worker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

// GetId 生成并返回唯一ID。
func (w *Worker) GetId() string {
	w.mu.Lock()
	defer w.mu.Unlock()

	now := time.Now().UnixNano() / 1e6 // 将纳秒转换为毫秒
	if w.timestamp == now {
		w.number++
		if w.number > numberMax {
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.number = 0
		w.timestamp = now
	}

	ID := strconv.FormatInt((now-epoch)<<timeShift|(w.workerId<<workerShift)|(w.number), 10)
	return ID
}
