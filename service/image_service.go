package service

import (
	"encoding/base64"
	"fmt"
	"imgconv/converter"
	"imgconv/models"
	"os"
	"path/filepath"
	"strings"
)

// ImageService 이미지 변환 서비스를 제공합니다.
type ImageService struct {
	converter *converter.Converter
}

// NewImageService returns a new ImageService instance.
func NewImageService() *ImageService {
	return &ImageService{
		converter: converter.NewConverter(),
	}
}

// ConvertImage 변환된 이미지 데이터를 base64 인코딩으로 반환합니다.
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

	convertedData, err := s.converter.Convert(inputFormat, outputFormat, decodedData)
	if err != nil {
		return "", fmt.Errorf("이미지 변환 오류: %v", err)
	}

	return base64.StdEncoding.EncodeToString(convertedData), nil
}

// Img2Img 이미지 파일을 다른 형식으로 변환하여 대상 경로에 저장합니다.
func (s *ImageService) Img2Img(req models.Img2ImgRequest) (string, error) {
	// 파일 읽기
	inputData, err := os.ReadFile(req.SourceFile)
	if err != nil {
		return "", fmt.Errorf("소스 파일 읽기 오류: %v", err)
	}

	// 입력 형식 추론 (예: 파일 확장자 사용)
	inputFormat := getFileExtension(req.SourceFile)
	if !converter.IsSupportedInputFormat(inputFormat) {
		return "", fmt.Errorf("지원하지 않는 입력 형식: %s", inputFormat)
	}

	// 출력 형식 처리 (jpeg와 jpg를 동일하게 취급)
	outputFormat := strings.ToLower(req.OutputFormat)
	if !converter.IsSupportedOutputFormat(req.OutputFormat) {
		return "", fmt.Errorf("지원하지 않는 출력 형식: %s", req.OutputFormat)
	}

	// 이미지 변환
	convertedData, err := s.converter.Convert(inputFormat, outputFormat, inputData)
	if err != nil {
		return "", fmt.Errorf("이미지 변환 오류: %v", err)
	}

	// TargetPath 마지막에 경로 구분자 추가
	targetPath := req.TargetPath
	if !strings.HasSuffix(targetPath, string(filepath.Separator)) {
		targetPath += string(filepath.Separator)
	}

	// 변환된 파일 경로 생성
	fileName := strings.TrimSuffix(filepath.Base(req.SourceFile), filepath.Ext(req.SourceFile))
	outputFilePath := targetPath + fmt.Sprintf("%s.%s", fileName, outputFormat)

	// 변환된 데이터 파일로 저장
	if err := os.WriteFile(outputFilePath, convertedData, os.ModePerm); err != nil {
		return "", fmt.Errorf("타겟 파일 쓰기 오류: %v", err)
	}

	return base64.StdEncoding.EncodeToString(convertedData), nil
}

// getFileExtension은 파일 경로에서 확장자를 추출합니다.
func getFileExtension(filePath string) string {
	ext := filepath.Ext(filePath)
	if len(ext) > 0 {
		return ext[1:] // 확장자에서 '.'을 제거
	}
	return ""
}
