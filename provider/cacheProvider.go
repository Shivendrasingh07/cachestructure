package provider

import "example.com/m/models"

type CacheProviderInterface interface {
	Set(data models.KeyValueStruct) error
	Get(key string) interface{}
	DeleteKeyValue(key string) error
}
