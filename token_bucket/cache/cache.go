package cache

type CacheI interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	Delete(key string)
}
