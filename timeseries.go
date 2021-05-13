package apmcharts

import (
	"io"
	"time"

	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

// RenderTimeseries renders time series
func RenderTimeseries(values, timestamps []float64, w io.Writer, options Options) error {
	times := make([]time.Time, 0, len(timestamps))

	for _, timestamp := range timestamps {
		times = append(times, time.Unix(int64(timestamp/1000), 0))
	}

	if options.StrokeColor.IsZero() {
		options.StrokeColor = drawing.ColorBlue
	}

	if options.FillColor.IsZero() {
		options.FillColor = drawing.ColorBlue.WithAlpha(20)
	}

	graph := chart.Chart{
		Height: options.Height,
		Width:  options.Width,
		XAxis: chart.XAxis{
			ValueFormatter: chart.TimeValueFormatterWithFormat("15:04"),
		},
		Series: []chart.Series{
			chart.TimeSeries{
				Style: chart.Style{
					StrokeColor: options.StrokeColor,
					FillColor:   options.FillColor,
					StrokeWidth: 2.0,
				},
				XValues: times,
				YValues: values,
			},
		},
	}

	return graph.Render(chart.PNG, w)
}
