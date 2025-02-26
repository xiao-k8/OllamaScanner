package util

import (
	"bytes"
)

func CheckResposeBody(respose []byte) (bool, error) {
	ollamaTrait := "Ollama is running"
	if bytes.Contains(respose, []byte(ollamaTrait)) == true {
		return true, nil
	} else {
		return false, nil
	}
}

func CheckApiTags(respose []byte) (bool, error) {
	ollamaTrait := "models"
	if bytes.Contains(respose, []byte(ollamaTrait)) == true {
		return true, nil
	} else {
		return false, nil
	}
}
