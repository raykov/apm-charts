package apmcharts

import (
	"io"
	"sort"
	"time"

	"github.com/pkg/errors"
	"github.com/wcharczuk/go-chart/v2"
)

// RenderResponseTime renders response time chart, aka. service sublayer
func RenderResponseTime(series, timestamps [][]float64, w io.Writer, options Options) error {
	if len(series) != len(timestamps) {
		return errors.New("RenderResponseTime: amount of series and timestamps should be equal")
	}

	convertedIndexed := map[float64]float64{}

	timeSeries := make([]chart.Series, 0, len(series))

	dataWithLegend := DataWithLegend{
		series:     series,
		legend:     options.Legend,
		timestamps: timestamps,
	}

	sort.Sort(dataWithLegend)

	for index, values := range dataWithLegend.series {
		times := make([]time.Time, 0, len(dataWithLegend.timestamps[index]))
		converted := make([]float64, 0, len(values))

		for subIndex, value := range values {
			t := dataWithLegend.timestamps[index][subIndex]
			if v, ok := convertedIndexed[t]; ok {
				convertedIndexed[t] = value + v
				converted = append(converted, value+v)
			} else {
				convertedIndexed[t] = value
				converted = append(converted, value)
			}

			times = append(times, time.Unix(int64(dataWithLegend.timestamps[index][subIndex]/1000), 0))
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
					YValues: converted,
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
