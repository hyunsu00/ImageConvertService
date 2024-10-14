package api

import (
	"encoding/json"
	"fmt"
	"imgconv/models"
	"imgconv/service"
	"net/http"
)

// ConvertHandler 이미지 변환 요청을 처리하는 핸들러 함수입니다.
func ConvertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "허용되지 않는 메서드", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		InputFormat  string `json:"input_format"`
		OutputFormat string `json:"output_format"`
		ImageData    string `json:"image_data"` // Base64 인코딩된 이미지 데이터
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "잘못된 요청 형식", http.StatusBadRequest)
		return
	}

	imageService := service.NewImageService()
	result, err := imageService.ConvertImage(request.InputFormat, request.OutputFormat, request.ImageData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"converted_image": result})
}

// Img2ImgHandler 이미지 파일을 다른 형식으로 변환하는 핸들러 함수입니다.
func Img2ImgHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "허용되지 않는 메서드", http.StatusMethodNotAllowed)
		return
	}

	// 요청 바디를 읽고 Img2ImgRequest 구조체로 디코딩합니다.
	var req models.Img2ImgRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("잘못된 요청 형식: %v", err), http.StatusBadRequest)
		return
	}

	// 요청 파라미터를 출력합니다.
	fmt.Printf("SourceFile: %s\n", req.SourceFile)
	fmt.Printf("TargetPath: %s\n", req.TargetPath)
	fmt.Printf("OutputFormat: %s\n", req.OutputFormat)

	// ImageService 인스턴스 생성
	imageService := service.NewImageService()

	// 이미지 변환
	_, err := imageService.Img2Img(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("이미지 변환 오류: %v", err), http.StatusInternalServerError)
		return
	}

	// 성공 응답을 반환합니다.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("이미지 변환 성공"))
}
