package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"

	apmcharts "github.com/raykov/apm-charts"
)

func main() {
	values := [][]float64{
		{
			300,
			299,
			299,
			307,
			250,
			299,
			290,
			310,
			311,
			280,
		},
		{
			300 - 50,
			299 - 49,
			299 - 52,
			307 - 50,
			250 - 50,
			299 - 54,
			290 - 50,
			310 - 45,
			311 - 50,
			280 - 20,
		},
		{
			300 - 150,
			299 - 149,
			299 - 152,
			307 - 150,
			250 - 150,
			299 - 154,
			290 - 150,
			310 - 145,
			311 - 150,
			280 - 120,
		},
	}

	times := []float64{
		float64(time.Now().AddDate(0, 0, -9).Unix()),
		float64(time.Now().AddDate(0, 0, -8).Unix()),
		float64(time.Now().AddDate(0, 0, -7).Unix()),
		float64(time.Now().AddDate(0, 0, -6).Unix()),
		float64(time.Now().AddDate(0, 0, -5).Unix()),
		float64(time.Now().AddDate(0, 0, -4).Unix()),
		float64(time.Now().AddDate(0, 0, -3).Unix()),
		float64(time.Now().AddDate(0, 0, -2).Unix()),
		float64(time.Now().AddDate(0, 0, -1).Unix()),
		float64(time.Now().Unix()),
	}

	legend := []string{
		"max",
		"p99",
		"p95",
	}

	options := apmcharts.Options{
		Width:  1024,
		Height: 400,

		Title:  "Latency",
		Legend: legend,
	}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Can't get current file name")
	}

	png := path.Join(path.Dir(filename), "latency.png")
	fmt.Printf("Writing chart to %s\n", png)
	file, err := os.Create(png)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	apmcharts.RenderLatency(
		values,
		times,
		file,
		options,
	)

}
