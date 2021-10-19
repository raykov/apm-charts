package apmcharts

import (
	"io"
	"time"

	"github.com/wcharczuk/go-chart/v2"
)

// RenderLatency renders latency chart
func RenderLatency(series [][]float64, timestamps []float64, w io.Writer, options Options) error {
	times := make([]time.Time, 0, len(timestamps))

	timeSeries := make([]chart.Series, 0, len(series))

	for _, timestamp := range timestamps {
		times = append(times, time.Unix(int64(timestamp/1000), 0))
	}

	for index, values := range series {
		timeSeries = append(
			timeSeries,
			chart.TimeSeries{
				Style: chart.Style{
					StrokeColor: colorSchema[index],
					StrokeWidth: 1,
				},
				XValues: times,
				YValues: values,
			},
		)
	}

	graph := chart.Chart{
		Height: options.Height,
		Width:  options.Width,
		XAxis: chart.XAxis{
			ValueFormatter: chart.TimeValueFormatterWithFormat("15:04"),
		},
		YAxis:  chart.YAxis{},
		Series: timeSeries,
	}

	return graph.Render(chart.PNG, w)
}
