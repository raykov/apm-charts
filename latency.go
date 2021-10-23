package apmcharts

import (
	"io"
	"time"

	"github.com/pkg/errors"
	"github.com/wcharczuk/go-chart/v2"
)

// RenderLatency renders latency chart
func RenderLatency(series, timestamps [][]float64, w io.Writer, options Options) error {
	if len(series) != len(timestamps) {
		return errors.New("RenderLatency: amount of series and timestamps should be equal")
	}

	timeSeries := make([]chart.Series, 0, len(series))

	for index, values := range series {
		times := make([]time.Time, 0, len(timestamps[index]))

		for _, timestamp := range timestamps[index] {
			times = append(times, time.Unix(int64(timestamp/1000), 0))
		}

		timeSeries = append(
			timeSeries,
			chart.TimeSeries{
				Name: getLegend(options.Legend, index),
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
		Height:       options.GetHeight(),
		Width:        options.GetWidth(),
		ColorPalette: options.GetColorPalette(),
		Title:        options.GetTitle(),
		TitleStyle:   options.GetTitleStyle(),

		XAxis: chart.XAxis{
			ValueFormatter: options.GetTimeFormatter(),
		},
		YAxis:  chart.YAxis{},
		Series: timeSeries,
	}

	if len(options.Legend) > 0 {
		graph.Elements = []chart.Renderable{
			chart.LegendLeft(&graph),
		}
	}

	return graph.Render(chart.PNG, w)
}
