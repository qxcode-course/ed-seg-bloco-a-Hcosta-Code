package main

import (
	"fmt"
	"math"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func desenhar(pen *Pen, x, y, tamanho float64) {
	if tamanho < 5 {
		return
	}

	pen.SetPosition(x, y)

	for i := 0; i < 6; i++ {

		anguloRad := float64(i) * 60 * (math.Pi / 180)
		novoX := x + tamanho*math.Cos(anguloRad)
		novoY := y - tamanho*math.Sin(anguloRad)

		desenhar(pen, novoX, novoY, tamanho/3)

		pen.SetPosition(novoX, novoY)
		pen.Walk(tamanho / 3)
		pen.Walk(-tamanho / 3)
		pen.Right(60)
	}
}

func main() {
	pen := NewPen(1000, 1000)
	pen.SetRGB(0, 0, 0)

	desenhar(pen, 500, 500, 250)

	pen.SavePNG("gelo.png")
	fmt.Println("Fractal gerado com sucesso em fractal_square.png")
}
