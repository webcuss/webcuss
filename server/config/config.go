package config

import "os"

func GetSecret() []byte {
	secret := os.Getenv("APP_SECRET")
	if secret != "" {
		return []byte(secret)
	} else {
		panic("App secret not found")
	}
}
