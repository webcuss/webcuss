package config

import "os"

func GetSecret() []byte {
	secret := os.Getenv("APP_SECRET")
	if secret != "" {
		return []byte(secret)
	} else {
		return []byte("ajfahfiAUHFUOfqjfjkahfusahfYAS98*AUV*asu8")
	}
}
