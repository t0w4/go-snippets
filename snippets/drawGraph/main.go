package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"golang.org/x/image/font/basicfont"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func main() {
	// 画像の大きさを定義
	const width, height = 256, 256

	// グラフの原点が左下からx, y軸でどれだけずれているか
	const padding = 8

	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	imgFilledWithWhite := fillWithWhite(img)

	// x軸を描画
	for x := 0; x < width; x++ {
		yVal := height - padding
		imgFilledWithWhite.Set(x, yVal, color.Gray{Y: 150})
	}

	// y軸を描画
	for y := 0; y < height; y++ {
		xVal := padding
		imgFilledWithWhite.Set(xVal, y, color.Gray{Y: 150})
	}

	// 原点の0を描画
	d := &font.Drawer{
		Dst:  imgFilledWithWhite,
		Src:  image.NewUniform(color.Black),
		Face: basicfont.Face7x13,
		Dot: fixed.Point26_6{
			X: (padding - 6) * 64, // 6は表示位置がちょうどよかっただけで特別な意味はない
			Y: (height + 1) * 64,  // 1は表示位置がちょうどよかっただけで特別な意味はない
		},
	}
	d.DrawString("0")

	// 対象の関数を描画
	for x := 0; x < width; x++ {
		xVal := x - padding
		yVal := height - calcY(x) + padding // paddingの分だけ上にずらす

		imgFilledWithWhite.Set(xVal, yVal, color.Black)
	}

	f, err := os.Create("image.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	if err := png.Encode(f, imgFilledWithWhite); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := f.Close(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// fillWithWhite imgを白で塗りつぶす
func fillWithWhite(img *image.NRGBA) *image.NRGBA {
	for y := 0; y < img.Rect.Dy(); y++ {
		for x := 0; x < img.Rect.Dx(); x++ {
			img.Set(x, y, color.White)
		}
	}
	return img
}

// calcY 関数f(x)、入力xから計算結果のyを返す
func calcY(x int) int {
	return x + 25
}
