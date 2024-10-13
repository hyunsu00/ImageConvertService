package service

import (
	"encoding/base64"
	"fmt"
	"imgconv/converter"
)

type ImageService struct {
	converter *converter.Converter
	semaphore chan struct{}
}

func NewImageService() *ImageService {
	return &ImageService{
		converter: converter.NewConverter(),
		semaphore: make(chan struct{}, 10), // 동시에 처리할 수 있는 최대 요청 수
	}
}

func (s *ImageService) ConvertImage(inputFormat, outputFormat, imageData string) (string, error) {
	if !converter.IsSupportedInputFormat(inputFormat) {
		return "", fmt.Errorf("지원하지 않는 입력 형식: %s", inputFormat)
	}
	if !converter.IsSupportedOutputFormat(outputFormat) {
		return "", fmt.Errorf("지원하지 않는 출력 형식: %s", outputFormat)
	}

	decodedData, err := base64.StdEncoding.DecodeString(imageData)
	if err != nil {
		return "", fmt.Errorf("이미지 데이터 디코딩 오류: %v", err)
	}

	s.semaphore <- struct{}{}        // 세마포어 획득
	defer func() { <-s.semaphore }() // 세마포어 해제

	convertedData, err := s.converter.Convert(inputFormat, outputFormat, decodedData)
	if err != nil {
		return "", fmt.Errorf("이미지 변환 오류: %v", err)
	}

	return base64.StdEncoding.EncodeToString(convertedData), nil
}
