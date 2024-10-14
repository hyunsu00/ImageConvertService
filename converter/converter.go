package converter

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"

	"github.com/srwiley/oksvg"   // SVG 이미지를 디코딩하고 인코딩
	"github.com/srwiley/rasterx" // 비트맵 이미지를 스캔하고 처리
	"golang.org/x/image/tiff"    // TIFF 이미지를 디코딩하고 인코딩
)

// Converter는 이미지 변환을 수행하는 구조체입니다.
type Converter struct{}

// Converter 생성자
func NewConverter() *Converter {
	return &Converter{}
}

// Convert 메서드는 입력 이미지 데이터를 디코딩하고, 지정된 출력 형식으로 인코딩하여 변환합니다.
func (c *Converter) Convert(inputFormat, outputFormat string, imageData []byte) ([]byte, error) {
	img, err := c.decode(inputFormat, imageData)
	if err != nil {
		return nil, err
	}

	return c.encode(outputFormat, img)
}

func (c *Converter) decode(format string, data []byte) (image.Image, error) {
	reader := bytes.NewReader(data)
	switch format {
	case "jpeg", "jpg":
		return jpeg.Decode(reader)
	case "png":
		return png.Decode(reader)
	case "gif":
		return gif.Decode(reader)
	case "tiff":
		return tiff.Decode(reader)
	case "svg":
		icon, _ := oksvg.ReadIconStream(reader)
		w, h := int(icon.ViewBox.W), int(icon.ViewBox.H)
		img := image.NewRGBA(image.Rect(0, 0, w, h))
		scanner := rasterx.NewScannerGV(w, h, img, img.Bounds())
		raster := rasterx.NewDasher(w, h, scanner)
		icon.Draw(raster, 1.0)
		return img, nil
	default:
		return nil, fmt.Errorf("지원하지 않는 입력 형식: %s", format)
	}
}

func (c *Converter) encode(format string, img image.Image) ([]byte, error) {
	var buf bytes.Buffer
	var err error

	switch format {
	case "jpeg", "jpg":
		err = jpeg.Encode(&buf, img, nil)
	case "png":
		err = png.Encode(&buf, img)
	case "gif":
		err = gif.Encode(&buf, img, nil)
	case "tiff":
		err = tiff.Encode(&buf, img, nil)
	default:
		return nil, fmt.Errorf("지원하지 않는 출력 형식: %s", format)
	}

	if err != nil {
		return nil, fmt.Errorf("이미지 인코딩 오류: %v", err)
	}

	return buf.Bytes(), nil
}
