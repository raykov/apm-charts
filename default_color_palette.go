package apmcharts

import (
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

// DefaultColorPalette represents the default palette.
var defaultColorPalette DefaultColorPalette

type DefaultColorPalette struct{}

// BackgroundColor returns background color
func (dp DefaultColorPalette) BackgroundColor() drawing.Color {
	return chart.ColorWhite
}

// BackgroundStrokeColor returns background stroke color
func (dp DefaultColorPalette) BackgroundStrokeColor() drawing.Color {
	return chart.ColorWhite
}

// CanvasColor returns canvas color
func (dp DefaultColorPalette) CanvasColor() drawing.Color {
	return chart.ColorWhite
}

// CanvasStrokeColor returns canvas stroke color
func (dp DefaultColorPalette) CanvasStrokeColor() drawing.Color {
	return chart.ColorWhite
}

// AxisStrokeColor returns axis stroke color
func (dp DefaultColorPalette) AxisStrokeColor() drawing.Color {
	return chart.ColorBlack
}

// TextColor returns text color
func (dp DefaultColorPalette) TextColor() drawing.Color {
	return chart.ColorBlack
}

// GetSeriesColor returns a color from the colorSchema list by index.
func (dp DefaultColorPalette) GetSeriesColor(index int) drawing.Color {
	finalIndex := index % len(colorSchema)
	return colorSchema[finalIndex]
}
