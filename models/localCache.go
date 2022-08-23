package models

import "time"

type Operations string

const (
	SetOperations    Operations = "set"
	GetOperations    Operations = "get"
	DeleteOperations Operations = "delete"
	ExitOperations   Operations = "exit"
)

type LocalCacheStruct struct {
	LocalCacheData interface{}
	ExpiryTime     map[string]interface{}
}

type ExpiryTime struct {
	Key  string
	Time time.Time
}

type KeyValueStruct struct {
	Key         string
	Value       interface{}
	TimeSeconds int
}

//type
