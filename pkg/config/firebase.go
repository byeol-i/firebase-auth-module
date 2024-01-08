package config

import (
	"flag"
	"os"

	"github.com/byeol-i/firebase-auth-module/pkg/logger"
)


var (
	local = flag.Bool("local", false, "using local cred")
	credFilePath = flag.String("firebaseCredFilePath", "/run/secrets/firebase-key", "cred path")
	firebaseProjectID = flag.String("firebaseProjectID", "worker-1234", "firebaseProjectID")
)

type FirebaseConfigImpl interface {
	GetFirebaseCredFilePath() (string)
	GetFirebaseProjectID() (string)
	GetApiKey() (string)
	GetAppID() (string)
}

type FirebaseConfig struct {
	FirebaseConfigImpl
}


func (c FirebaseConfig) GetFirebaseCredFilePath() string {
	if *local {
		logger.Info("Using local conf")
		return "firebase/key.json"
	}

	return *credFilePath
}

func (c FirebaseConfig) GetFirebaseProjectID() string {
	return *firebaseProjectID
}

func (c FirebaseConfig) GetApiKey() string {
	key := os.Getenv("API_KEY")

	return key
}

func (c FirebaseConfig) GetAppID() string {
	appID := os.Getenv("APP_ID")

	return appID
}
