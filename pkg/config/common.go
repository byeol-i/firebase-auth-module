package config

import (
	"flag"
	"sync"
)

var lock = &sync.Mutex{}

type configManager struct {
	FirebaseConfig
	GrpcConfig
}

var configManagerInstance *configManager

func GetInstance() *configManager {
	if configManagerInstance == nil {
		lock.Lock()

		flag.Parse()
		
		if configManagerInstance == nil {
			configManagerInstance = &configManager{}
		} 

		defer lock.Unlock()
	}

	return configManagerInstance
}