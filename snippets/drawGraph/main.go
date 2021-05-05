package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

type Graph struct {
	width   int          // 画像の幅
	height  int          // 画像の高さ
	padding int          // グラフの原点が左下からx, y軸でどれだけずれているか
	img     *image.NRGBA // 描画用のイメージ
}

func NewGraph(width, height, padding int) Graph {
	return Graph{
		width:   width,
		height:  height,
		padding: padding,
		img:     image.NewNRGBA(image.Rect(0, 0, width, height)),
	}
}

func (g Graph) GetImg() *image.NRGBA {
	return g.img
}

// FillWithWhite imgを白で塗りつぶす
func (g Graph) FillWithWhite() {
	for y := 0; y < g.img.Rect.Dy(); y++ {
		for x := 0; x < g.img.Rect.Dx(); x++ {
			g.img.Set(x, y, color.White)
		}
	}
}

// Plot x, y座標で指定された点にプロットする
func (g Graph) Plot(x, y int, c color.Color) {
	g.img.Set(
		x+g.padding,
		g.height-y-g.padding,
		c,
	)
}

// DrawXAxis x軸を描画する
func (g Graph) DrawXAxis() {
	for x := 0; x < g.width; x++ {
		yVal := g.height - g.padding
		g.img.Set(x, yVal, color.Gray{Y: 150})
	}
}

// DrawYAxis y軸を描画する
func (g Graph) DrawYAxis() {
	for y := 0; y < g.height; y++ {
		g.img.Set(g.padding, y, color.Gray{Y: 150})
	}
}

func (g Graph) DrawOrigin() {
	// 原点の0を描画
	d := &font.Drawer{
		Dst:  g.img,
		Src:  image.NewUniform(color.Black),
		Face: basicfont.Face7x13,
		Dot: fixed.Point26_6{
			X: fixed.Int26_6((g.padding - 6) * 64), // 6は表示位置がちょうどよかっただけで特別な意味はない
			Y: fixed.Int26_6((g.height + 1) * 64),  // 1は表示位置がちょうどよかっただけで特別な意味はない
		},
	}
	d.DrawString("0")
}

func (g Graph) DrawLinearFunction(f func(int) int) {
	// 対象の関数を描画
	for x := -1 * g.padding; x < g.width-g.padding; x++ {
		g.Plot(x, f(x), color.Black)
	}
}

func main() {
	// 画像の大きさを定義
	const width, height = 256, 256
	// グラフの原点が左下からx, y軸でどれだけずれているか
	const padding = 8

	graph := NewGraph(width, height, padding)
	graph.FillWithWhite()
	graph.DrawXAxis()
	graph.DrawYAxis()
	graph.DrawOrigin()

	// 対象の関数を描画
	graph.DrawLinearFunction(calcY)

	f, err := os.Create("image.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	if err := png.Encode(f, graph.GetImg()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := f.Close(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// calcY 関数f(x)、入力xから計算結果のyを返す
func calcY(x int) int {
	return x
}
