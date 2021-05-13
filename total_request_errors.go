package apmcharts

import (
	"fmt"
	"io"
	"time"

	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

// RenderTotalRequestErrors renders throughput with errors
func RenderTotalRequestErrors(values [][]float64, timestamps []float64, w io.Writer, options Options) error {
	times := make([]time.Time, 0, len(timestamps))

	for _, timestamp := range timestamps {
		times = append(times, time.Unix(int64(timestamp/1000), 0))
	}

	max := Max(values[0])
	mx := max + max*0.1

	graph := chart.Chart{
		Height: options.Height,
		Width:  options.Width,
		XAxis: chart.XAxis{
			ValueFormatter: chart.TimeValueFormatterWithFormat("15:04"),
		},
		YAxis: chart.YAxis{
			Ticks: []chart.Tick{
				{0.0, "0k"},
				{mx * 0.125, fmt.Sprintf("%.0fk", (mx*0.125)/1000)},
				{mx * 0.25, fmt.Sprintf("%.0fk", mx/4000)},
				{mx * 0.375, fmt.Sprintf("%.0fk", (mx*0.375)/1000)},
				{mx * 0.5, fmt.Sprintf("%.0fk", mx/2000)},
				{mx * 0.625, fmt.Sprintf("%.0fk", (mx*0.625)/1000)},
				{mx * 0.75, fmt.Sprintf("%.0fk", (mx*0.75)/1000)},
				{mx * 0.865, fmt.Sprintf("%.0fk", (mx*0.865)/1000)},
				{mx, fmt.Sprintf("%.0fk", mx/1000)},
			},
		},
		Series: []chart.Series{
			chart.TimeSeries{
				Style: chart.Style{
					StrokeColor: drawing.Color{R: 51, G: 153, B: 204, A: 255},
					FillColor:   drawing.Color{R: 51, G: 153, B: 204, A: 255},
					StrokeWidth: 3.0,
				},
				XValues: times,
				YValues: values[0],
			},
			chart.TimeSeries{
				Style: chart.Style{
					StrokeColor: drawing.Color{R: 227, G: 27, B: 27, A: 255},
					FillColor:   drawing.Color{R: 227, G: 27, B: 27, A: 255},
					StrokeWidth: 3.0,
				},
				XValues: times,
				YValues: values[1],
			},
		},
	}

	return graph.Render(chart.PNG, w)
}

func Max(values []float64) (max float64) {
	for _, val := range values {
		if val > max {
			max = val
		}
	}

	return max
}
