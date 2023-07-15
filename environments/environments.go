package environments

import (
	"errors"
	"os"
)

func GetEnvironments() (string, error) {
	var tgKey string

	tgKey = os.Getenv("TG_KEY")
	if tgKey == "" {
		return "", errors.New("переменная окружения TG_KEY не задана")
	}
	return tgKey, nil

}
