package apmcharts

import (
	"io"
	"sort"
	"time"

	"github.com/wcharczuk/go-chart/v2"
)

// DataWithLegend used to sort series to put less time consuming services down in chart
type DataWithLegend struct {
	series [][]float64
	legend []string
}

// Len returns length
func (d DataWithLegend) Len() int {
	return len(d.series)
}

// Less checks if one item less than another
func (d DataWithLegend) Less(i, j int) bool {
	return d.series[j][0] > d.series[i][0]
}

// Swap swaps items
func (d DataWithLegend) Swap(i, j int) {
	d.series[i], d.series[j] = d.series[j], d.series[i]
	if len(d.legend) == len(d.series) {
		d.legend[i], d.legend[j] = d.legend[j], d.legend[i]
	}
}

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

		name := ""
		if len(dataWithLegend.legend) == len(dataWithLegend.series) {
			name = dataWithLegend.legend[index]
		}

		timeSeries = append(
			[]chart.Series{
				chart.TimeSeries{
					Name: name,
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
		Height: options.Height,
		Width:  options.Width,
		XAxis: chart.XAxis{
			ValueFormatter: chart.TimeValueFormatterWithFormat("15:04"),
		},
		YAxis: chart.YAxis{},
		Series: timeSeries,
	}

	graph.Elements = []chart.Renderable{
		chart.LegendLeft(&graph),
	}

	return graph.Render(chart.PNG, w)
}
