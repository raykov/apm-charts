package apmcharts

// DataWithLegend used to sort series to put less time consuming services down in chart
type DataWithLegend struct {
	series     [][]float64
	timestamps [][]float64
	legend     []string
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
	d.timestamps[i], d.timestamps[j] = d.timestamps[j], d.timestamps[i]

	if len(d.legend) == len(d.series) {
		d.legend[i], d.legend[j] = d.legend[j], d.legend[i]
	}
}
