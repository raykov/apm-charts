package apmcharts

import (
	"io"
	"sort"
	"time"

	"github.com/wcharczuk/go-chart/v2"
)

// RenderResponseTime renders response time chart, aka. service sublayer
func RenderResponseTime(series [][]float64, timestamps []float64, w io.Writer, options Options) error {
	times := make([]time.Time, 0, len(timestamps))
	converted := make([][]float64, len(series))

	timeSeries := make([]chart.Series, 0, len(series))

	dataWithLegend := DataWithLegend{
		series: series,
		legend: options.Legend,
	}

	sort.Sort(dataWithLegend)

	for index, values := range dataWithLegend.series {
		converted[index] = make([]float64, 0, len(values))

		for subIndex, value := range values {
			if index == 0 {
				times = append(times, time.Unix(int64(timestamps[subIndex]/1000), 0))
				converted[index] = append(converted[index], value)
			} else {
				converted[index] = append(converted[index], value+converted[index-1][subIndex])
			}
		}

		timeSeries = append(
			[]chart.Series{
				chart.TimeSeries{
					Name: getLegend(dataWithLegend.legend, index),
					Style: chart.Style{
						StrokeColor: colorSchema[index],
						FillColor:   colorSchema[index],
						StrokeWidth: 1,
					},
					XValues: times,
					YValues: converted[index],
				},
			},
			timeSeries...,
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
