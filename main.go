package main

import (
	"imgconv/api"
	"log"
	"net/http"
	"strings"
)

// ContextPathHandler는 컨텍스트 경로를 처리하는 사용자 정의 핸들러입니다.
type ContextPathHandler struct {
	contextPath string
	handler     http.Handler
}

// ServeHTTP는 HTTP 요청을 처리합니다.
func (c *ContextPathHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch c.contextPath {
	case "/", "":
		// 컨텍스트 경로가 "/" 또는 빈 문자열("")인 경우 경로를 수정하지 않고 바로 핸들러 호출
		c.handler.ServeHTTP(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, c.contextPath) {
		// 컨텍스트 경로를 제거하고 요청을 핸들러에 전달
		r.URL.Path = strings.TrimPrefix(r.URL.Path, c.contextPath)
		c.handler.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/convert", api.ConvertHandler)

	// 컨텍스트 경로를 설정하여 멀티플렉서 래핑
	contextPath := ""
	contextPathHandler := &ContextPathHandler{
		contextPath: contextPath,
		handler:     mux,
	}

	log.Println("서버 시작: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", contextPathHandler))
}
