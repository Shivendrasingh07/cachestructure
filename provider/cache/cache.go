package cache

import (
	"encoding/json"
	"example.com/m/provider"
	"fmt"
	"time"

	"example.com/m/models"
	"github.com/sirupsen/logrus"
)

//var LocalCachedData models.LocalCacheStruct

type cacheProvider struct {
	LocalCachedData models.LocalCacheStruct
	//cacheProvider provider.cacheProvider
}

func NewCacheProvider() provider.CacheProviderInterface {
	var LocalCached models.LocalCacheStruct
	return &cacheProvider{
		LocalCachedData: LocalCached,
	}
}

func (c *cacheProvider) Set(data models.KeyValueStruct) error {

	tempData := map[string]interface{}{
		data.Key: data.Value,
	}

	expTimeData := time.Now().Add(time.Duration(data.TimeSeconds))
	tempTimeData := map[string]interface{}{
		data.Key: expTimeData,
	}
	metaData, err := json.Marshal(tempData)
	if err != nil {
		logrus.Errorf("%v", err)
	}
	c.LocalCachedData.LocalCacheData = metaData
	c.LocalCachedData.ExpiryTime = tempTimeData

	go func() {
		time.Sleep(time.Second * time.Duration(data.TimeSeconds))
		err := c.DeleteKeyValue(data.Key)
		if err != nil {
			fmt.Println("unable to delete data")
			return
		}
	}()

	return nil
}

func (c *cacheProvider) Get(key string) interface{} {
	tempData := make(map[string]interface{})
	err := json.Unmarshal(c.LocalCachedData.LocalCacheData.([]byte), &tempData)
	if err != nil {
		fmt.Println("unable to unmarshal data")
		return err
	}
	return tempData[key]
}

func (c *cacheProvider) DeleteKeyValue(key string) error {
	tempData := make(map[string]interface{})
	err := json.Unmarshal(c.LocalCachedData.LocalCacheData.([]byte), &tempData)
	if err != nil {
		fmt.Println("unable to unmarshal data")
		return err
	}

	tempData[key] = nil

	metaData, err := json.Marshal(tempData)
	if err != nil {
		logrus.Errorf("%v", err)
	}
	c.LocalCachedData.LocalCacheData = metaData
	return nil
}
