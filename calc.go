package main

import (
	"github.com/go-echarts/go-echarts/charts"
	"log"
	"math"
	"os"
)

const e  = 2.718281828459

func Mod1(x float64) float64 {
	switch {
		case x < 25:
			return 10
		case x >90:
			return 100
	default:
		p := x / 7  *  math.Log2(x)
		c := math.Floor(p / 5) * 5
		return c
	}
}


func Mod2(x float64) float64 {
	switch {
	case x < 25:
		return 10
	case x >90:
		return 100
	default:
		p :=  x / 4.3  *  math.Log(x)
		c := math.Floor(p / 5) * 5
		return c
	}
}

func image(xAxis []float64 ,yAxis []float64)  {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "转速-温度图"})
	line.AddXAxis(xAxis).AddYAxis("转速百分比%", yAxis)
	f, err := os.Create("line.html")
	if err != nil {
		log.Println(err)
	}
	line.Render(f)
}
