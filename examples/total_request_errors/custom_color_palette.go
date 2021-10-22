package main

import (
	apmcharts "github.com/raykov/apm-charts"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

// Color Schema
var colorSchema = []drawing.Color{
	{R: 105, G: 203, B: 226, A: 255},
	{R: 227, G: 26, B: 28, A: 255},
}

// CustomColorPalette represents the custom palette.
var CustomColorPalette customColorPalette

type customColorPalette struct {
	apmcharts.DefaultColorPalette
}

// GetSeriesColor returns a color from the colorSchema list by index.
func (dp customColorPalette) GetSeriesColor(index int) drawing.Color {
	finalIndex := index % len(colorSchema)
	return colorSchema[finalIndex]
}
