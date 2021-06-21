package main

import (
	"crypto/rand"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/big"
	"os"
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

func main() {
	// 画像の大きさを定義
	const width, height = 256, 256
	// グラフの原点が左下からx, y軸でどれだけずれているか
	const padding = 0

	graph := NewGraph(width, height, padding)
	graph.FillWithWhite()
	for i := 0; i < 100; i++ {
		x, _ := rand.Int(rand.Reader, big.NewInt(width))
		y, _ := rand.Int(rand.Reader, big.NewInt(height))
		graph.Plot(
			int(x.Int64()),
			int(y.Int64()),
			color.Black,
		)
	}

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
