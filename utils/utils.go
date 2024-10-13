package utils

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func ReadImageFile(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("파일 읽기 오류: %v", err)
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func WriteImageFile(filePath string, data string) error {
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return fmt.Errorf("이미지 데이터 디코딩 오류: %v", err)
	}
	return ioutil.WriteFile(filePath, decodedData, 0644)
}
