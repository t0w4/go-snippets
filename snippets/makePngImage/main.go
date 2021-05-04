package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	// 画像の大きさを定義
	const width, height = 256, 256

	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// RGBAはuint8で0~255の値をとる
			// 255との論理積をとることで256以上の値にならないようにしている
			// RGBの各要素をxに比例するようにすると、縦方向の色が同じになる
			// RGBの各要素をyに比例するようにすると、横方向の色が同じになる
			// G, Bをビットシフトさせているのは各値が同じだとグレースケールになるため
			img.Set(x, y, color.NRGBA{
				R: uint8(x & 255),      // 赤
				G: uint8(x << 2 & 255), // 緑
				B: uint8(x << 3 & 255), // 青
				A: 255,                 // 透明度(大きいほど不透明に)
			})
		}
	}

	f, err := os.Create("image.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := f.Close(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
