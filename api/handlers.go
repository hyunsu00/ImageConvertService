package api

import (
	"encoding/json"
	"imgconv/service"
	"net/http"
)

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
