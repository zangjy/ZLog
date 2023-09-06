package utils

import "sync"

// KeyPair 包含公钥和共享密钥。
type KeyPair struct {
	PublicKey string
	SharedKey string
}

// GlobalMap 包含全局的 Map 和一个互斥锁用于并发控制。
type GlobalMap struct {
	data map[string]KeyPair
	mu   sync.Mutex
}

var globalMap = &GlobalMap{
	data: make(map[string]KeyPair),
}

// Put 将键值对插入全局 Map 中。
func Put(key, publicKey, sharedSecret string) {
	globalMap.mu.Lock()
	defer globalMap.mu.Unlock()
	globalMap.data[key] = KeyPair{
		PublicKey: publicKey,
		SharedKey: sharedSecret,
	}
}

// Get 从全局 Map 中根据键检索值。
func Get(key string) (KeyPair, bool) {
	globalMap.mu.Lock()
	defer globalMap.mu.Unlock()
	value, ok := globalMap.data[key]
	return value, ok
}

// Delete 从全局 Map 中删除键值对。
func Delete(key string) {
	globalMap.mu.Lock()
	defer globalMap.mu.Unlock()
	delete(globalMap.data, key)
}

// Size 返回全局 Map 中键值对的数量。
func Size() int {
	globalMap.mu.Lock()
	defer globalMap.mu.Unlock()
	return len(globalMap.data)
}
