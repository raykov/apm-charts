package apmcharts

import (
	"fmt"
	"io"
	"time"

	"github.com/pkg/errors"
	"github.com/wcharczuk/go-chart/v2"
)

// RenderTotalRequestErrors renders throughput with errors
func RenderTotalRequestErrors(values, timestamps [][]float64, w io.Writer, options Options) error {
	if len(values) != 2 {
		return errors.New("RenderTotalRequestErrors: values should be a slice of 2 slices")
	}
	if len(values) != len(timestamps) {
		return errors.New("RenderTotalRequestErrors: amount of values and timestamps should be equal")
	}

	times := make([][]time.Time, 0, len(timestamps))

	for index, subTimestamps := range timestamps {
		times = append(times, make([]time.Time, 0, len(subTimestamps)))

		for _, timestamp := range subTimestamps {
			times[index] = append(times[index], time.Unix(int64(timestamp/1000), 0))
		}

	}

	max := Max(values[0])
	mx := max + max*0.1

	graph := chart.Chart{
		Height:       options.GetHeight(),
		Width:        options.GetWidth(),
		ColorPalette: options.GetColorPalette(),
		Title:        options.GetTitle(),
		TitleStyle:   options.GetTitleStyle(),

		XAxis: chart.XAxis{
			ValueFormatter: options.GetTimeFormatter(),
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
					StrokeColor: options.GetColorPalette().GetSeriesColor(0),
					FillColor:   options.GetColorPalette().GetSeriesColor(0),
					StrokeWidth: 3.0,
				},
				XValues: times[0],
				YValues: values[0],
			},
			chart.TimeSeries{
				Style: chart.Style{
					StrokeColor: options.GetColorPalette().GetSeriesColor(1),
					FillColor:   options.GetColorPalette().GetSeriesColor(1),
					StrokeWidth: 3.0,
				},
				XValues: times[1],
				YValues: values[1],
			},
		},
	}

	return graph.Render(chart.PNG, w)
}
