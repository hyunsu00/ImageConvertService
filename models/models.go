package models

// Img2ImgRequest는 이미지 변환 요청의 구조체입니다.
type Img2ImgRequest struct {
	SourceFile   string `json:"sourceFile"`
	TargetPath   string `json:"targetPath"`
	OutputFormat string `json:"outputFormat"`
}
