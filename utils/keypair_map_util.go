package utils

import (
	"sync"
	"time"
)

const defaultExpireAfter = 1 * time.Hour

type KeyPair struct {
	PublicKey string
	SharedKey string
	ExpireAt  time.Time
}

type globalMap struct {
	data map[string]KeyPair
	mu   sync.Mutex
}

var globalMapInstance = &globalMap{
	data: make(map[string]KeyPair),
}

func Put(key, publicKey, sharedSecret string, expireAfter ...time.Duration) {
	globalMapInstance.mu.Lock()
	defer globalMapInstance.mu.Unlock()

	expireDuration := defaultExpireAfter

	if len(expireAfter) > 0 {
		expireDuration = expireAfter[0]
	}

	expireAt := time.Now().Add(expireDuration)

	globalMapInstance.data[key] = KeyPair{
		PublicKey: publicKey,
		SharedKey: sharedSecret,
		ExpireAt:  expireAt,
	}

	//启动定时器，在过期后删除KeyPair
	go func() {
		<-time.After(expireDuration)
		Delete(key) //调用删除方法
	}()
}

func Get(key string) (KeyPair, bool) {
	globalMapInstance.mu.Lock()
	defer globalMapInstance.mu.Unlock()
	value, ok := globalMapInstance.data[key]

	//如果KeyPair过期，删除它
	if ok && time.Now().After(value.ExpireAt) {
		Delete(key) //调用删除方法
		return KeyPair{}, false
	}

	return value, ok
}

func Delete(key string) {
	globalMapInstance.mu.Lock()
	defer globalMapInstance.mu.Unlock()
	delete(globalMapInstance.data, key)
}

func Size() int {
	globalMapInstance.mu.Lock()
	defer globalMapInstance.mu.Unlock()
	return len(globalMapInstance.data)
}
