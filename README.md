# ImageConvertService


## Go 환경

```bash
# 일반적인 설치 경로
$ which go
# Go 설치 경로 확인
$ go env GOROOT
# Go 환경 변수 확인
$ go env
```

### Go 모듈 초기화

`$ go mod init imgconv`

### Go 패키지 설치

```Bash
# 패키지 설치
$ go get github.com/srwiley/oksvg github.com/srwiley/rasterx golang.org/x/image/tiff

# 버전 고정 패키지 설치
$ go get github.com/srwiley/oksvg@v0.0.0-20221011165216-be6e8873101c
$ go get github.com/srwiley/rasterx@v0.0.0-20220730225603-2ab79fcdd4ef
$ go get golang.org/x/image@v0.21.0
$ go get golang.org/x/net@v0.0.0-20211118161319-6a13c67c3ce4
$ go get golang.org/x/text@v0.19.0

# go.mod 파일에 명시된 패키지를 설치하고, 필요 없는 패키지를 정리
$ go mod tidy
```

### Go 실행

`$ go run main.go`

### Go 빌드

`$ go build -o hellowolrd`

### 모든 캐쉬 지우기

`go clean -cache -modcache -testcache`

### Go 디버깅

```bash
# 주의사항
# go-dlv는 브레이크 포인트 설정시 심볼릭링크경로에는 붙지 않는다.
# 따라서 실제 경로에서 프로젝트에서 열어야만 디버깅이 가능하다.

go install github.com/go-delve/delve/cmd/dlv@latest
which dlv
```
