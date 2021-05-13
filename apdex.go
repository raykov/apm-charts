package apmcharts

import (
	"io"
	"time"

	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

// RenderApdex renders Apdex chart
func RenderApdex(values, timestamps []float64, w io.Writer, options Options) error {
	times := make([]time.Time, 0, len(timestamps))
	yValues5000 := make([]float64, 0, len(timestamps))
	yValues7500 := make([]float64, 0, len(timestamps))
	yValues8750 := make([]float64, 0, len(timestamps))
	yValues9375 := make([]float64, 0, len(timestamps))
	yValues1000 := make([]float64, 0, len(timestamps))

	for _, timestamp := range timestamps {
		times = append(times, time.Unix(int64(timestamp/1000), 0))
		yValues5000 = append(yValues5000, 0.5)
		yValues7500 = append(yValues7500, 0.75)
		yValues8750 = append(yValues8750, 0.875)
		yValues9375 = append(yValues9375, 0.9375)
		yValues1000 = append(yValues1000, 1.0)
	}

	graph := chart.Chart{
		Height: options.Height,
		Width:  options.Width,
		XAxis: chart.XAxis{
			ValueFormatter: chart.TimeValueFormatterWithFormat("15:04"),
		},
		YAxis: chart.YAxis{
			Ticks: []chart.Tick{
				{0.0, "0.0"},
				{0.25, "0.25"},
				{0.5, "0.5"},
				{0.75, "0.75"},
				{1.0, "1.0"},
				{1.1, ""},
			},
		},
		Series: []chart.Series{
			chart.TimeSeries{
				Style: chart.Style{
					StrokeColor: drawing.Color{R: 217, G: 249, B: 249, A: 255},
					FillColor:   drawing.Color{R: 217, G: 249, B: 249, A: 255},
					StrokeWidth: 0.2,
				},
				XValues: times,
				YValues: yValues1000,
			},
			chart.TimeSeries{
				Style: chart.Style{
					StrokeColor: drawing.Color{R: 0, G: 217, B: 217, A: 255},
					FillColor:   drawing.Color{R: 233, G: 249, B: 233, A: 255},
					StrokeWidth: 0.5,
				},
				XValues: times,
				YValues: yValues9375,
			},
			chart.TimeSeries{
				Style: chart.Style{
					StrokeColor: drawing.Color{R: 106, G: 216, B: 110, A: 255},
					FillColor:   drawing.Color{R: 248, G: 247, B: 226, A: 255},
					StrokeWidth: 0.5,
				},
				XValues: times,
				YValues: yValues8750,
			},
			chart.TimeSeries{
				Style: chart.Style{
					StrokeColor: drawing.Color{R: 206, G: 199, B: 63, A: 255},
					FillColor:   drawing.Color{R: 254, G: 243, B: 243, A: 255},
					StrokeWidth: 1,
				},
				XValues: times,
				YValues: yValues7500,
			},
			chart.TimeSeries{
				Style: chart.Style{
					StrokeColor: drawing.Color{R: 219, G: 141, B: 140, A: 255},
					FillColor:   drawing.Color{R: 237, G: 237, B: 237, A: 255},
					StrokeWidth: 0.5,
				},
				XValues: times,
				YValues: yValues5000,
			},
			chart.TimeSeries{
				Style: chart.Style{
					StrokeColor: drawing.Color{R: 105, G: 203, B: 226, A: 255},
					StrokeWidth: 3.0,
				},
				XValues: times,
				YValues: values,
			},
		},
	}

	return graph.Render(chart.PNG, w)
}
