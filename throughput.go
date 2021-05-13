package apmcharts

import (
	"fmt"
	"io"
	"time"

	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

// RenderThroughput renders throughput chart
func RenderThroughput(values, timestamps []float64, w io.Writer, options Options) error {
	times := make([]time.Time, 0, len(timestamps))

	for _, timestamp := range timestamps {
		times = append(times, time.Unix(int64(timestamp/1000), 0))
	}

	if options.TimeFormat == "" {
		options.TimeFormat = "15:04"
	}

	max := Max(values)
	mx := max + max*0.1

	graph := chart.Chart{
		Height: options.Height,
		Width:  options.Width,
		XAxis: chart.XAxis{
			ValueFormatter: chart.TimeValueFormatterWithFormat(options.TimeFormat),
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
					StrokeColor: drawing.Color{R: 0, G: 185, B: 217, A: 255},
					StrokeWidth: 3.0,
				},
				XValues: times,
				YValues: values,
			},
		},
	}

	return graph.Render(chart.PNG, w)
}
