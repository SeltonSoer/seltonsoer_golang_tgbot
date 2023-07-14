package environments

import (
	"errors"
	"os"
)

func GetEnvironments() (string, error) {
	var tgKey string

	appEnv := os.Getenv("APP_ENV")

	if appEnv == "local" {
		tgKey = os.Getenv("TG_DEV_KEY")
		if tgKey == "" {
			return "", errors.New("переменная окружения TG_DEV_KEY не задана")
		}
		return tgKey, nil
	} else if appEnv == "prod" {
		tgKey = os.Getenv("TG_PROD_KEY")
		if tgKey == "" {
			return "", errors.New("переменная окружения TG_PROD_KEY не задана")
		}
		return tgKey, nil
	} else {
		return "", errors.New("переменная окружения APP_ENV не задана")
	}

}
