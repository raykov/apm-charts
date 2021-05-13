package apmcharts

import (
	"github.com/wcharczuk/go-chart/v2/drawing"
)

// Options used to set chart options
type Options struct {
	Width  int
	Height int

	Legend []string

	TimeFormat  string
	StrokeColor drawing.Color
	FillColor   drawing.Color
}
