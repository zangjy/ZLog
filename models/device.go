package models

import (
	"ZLog/dao"
	"ZLog/utils"
	"gorm.io/gorm"
)

type Device struct {
	Id           uint `gorm:"primaryKey"`
	AppId        string
	DeviceType   int
	DeviceName   string
	DeviceId     string `gorm:"unique"`
	ClientPubKey string
	SharedKey    string
	SessionID    string `gorm:"unique"`
}

//
// DeviceRegister
//  @Description: 设备注册
//  @param appId
//  @param deviceType
//  @param deviceName
//  @param deviceId
//  @param clientPubKey
//  @param sharedKey
//  @param TmpSessionId
//  @return string
//  @return bool
//
func DeviceRegister(appId string, deviceType int, deviceName string, deviceId string, clientPubKey string, sharedKey string, TmpSessionId string) (string, bool) {
	//在数据库中查找记录
	var existingDevice Device
	result := dao.DB.Table("device").Where("device_id = ?", deviceId).First(&existingDevice)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			//记录不存在，创建新记录
			newDevice := Device{
				AppId:        appId,
				DeviceType:   deviceType,
				DeviceName:   deviceName,
				DeviceId:     deviceId,
				ClientPubKey: clientPubKey,
				SharedKey:    sharedKey,
				SessionID:    TmpSessionId, //使用传入的 TmpSessionId
			}
			if err := dao.DB.Table("device").Create(&newDevice).Error; err != nil {
				return "", false
			}
			return TmpSessionId, true
		}
		//其他错误
		return "", false
	}

	//记录已存在，更新记录的 clientPubKey 和 sharedKey
	existingDevice.ClientPubKey = clientPubKey
	existingDevice.SharedKey = sharedKey

	if err := dao.DB.Table("device").Save(&existingDevice).Error; err != nil {
		return "", false
	}

	return existingDevice.SessionID, true
}

//
// GetKeyPairBySessionId
//  @Description: 通过SessionId获取密钥对
//  @param sessionId
//  @return utils.KeyPair
//  @return error
//
func GetKeyPairBySessionId(sessionId string) (utils.KeyPair, error) {
	if keyPair, state := utils.Get(sessionId); state {
		return keyPair, nil
	}
	var device Device
	if err := dao.DB.Table("device").Where("session_id = ?", sessionId).Select("client_pub_key, shared_key").First(&device).Error; err != nil {
		return utils.KeyPair{}, err
	}
	//记录到Map中
	utils.Put(sessionId, device.ClientPubKey, device.SharedKey)
	return utils.KeyPair{
		PublicKey: device.ClientPubKey,
		SharedKey: device.SharedKey,
	}, nil
}
