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
	values := []float64{
		100_000,
		102_000,
		130_000,
		150_000,
		170_000,
		173_000,
		178_000,
		170_000,
		175_000,
		174_000,
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

	options := apmcharts.Options{
		Width:  1024,
		Height: 400,

		Title: "Timeseries",
	}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Can't get current file name")
	}

	png := path.Join(path.Dir(filename), "timeseries.png")
	fmt.Printf("Writing chart to %s\n", png)
	file, err := os.Create(png)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	apmcharts.RenderTimeseries(
		values,
		times,
		file,
		options,
	)

}
