package apmcharts

import (
	"io"
	"time"

	"github.com/pkg/errors"
	"github.com/wcharczuk/go-chart/v2"
)

// RenderTimeseries renders time series
func RenderTimeseries(values, timestamps []float64, w io.Writer, options Options) error {
	if len(values) != len(timestamps) {
		return errors.New("RenderTimeseries: amount of values and timestamps should be equal")
	}

	times := make([]time.Time, 0, len(timestamps))

	for _, timestamp := range timestamps {
		times = append(times, time.Unix(int64(timestamp/1000), 0))
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
		Series: []chart.Series{
			chart.TimeSeries{
				Style: chart.Style{
					StrokeColor: options.GetColorPalette().GetSeriesColor(0),
					FillColor:   options.GetColorPalette().GetSeriesColor(0).WithAlpha(20),
					StrokeWidth: 2.0,
				},
				XValues: times,
				YValues: values,
			},
		},
	}

	return graph.Render(chart.PNG, w)
}
