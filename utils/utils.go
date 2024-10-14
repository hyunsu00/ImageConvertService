package utils

import (
	"encoding/base64"
	"fmt"
	"os"
)

// ReadImageFile 이미지 파일을 읽고 base64로 인코딩된 이미지 데이터를 반환합니다.
func ReadImageFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("파일 읽기 오류: %v", err)
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

// WriteImageFile base64로 인코딩된 이미지 데이터를 이미지 파일에 씁니다.
func WriteImageFile(filePath string, data string) error {
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return fmt.Errorf("이미지 데이터 디코딩 오류: %v", err)
	}
	return os.WriteFile(filePath, decodedData, 0644)
}
