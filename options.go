package apmcharts

import (
	"github.com/wcharczuk/go-chart/v2"
)

const (
	defaultTimeFormat = "15:04"
)

// Options used to set chart options
type Options struct {
	Width  int
	Height int

	Legend []string

	TimeFormat string

	ColorPalette chart.ColorPalette

	Title      string
	TitleStyle *chart.Style
}

// GetColorPalette returns custom or default color palette
func (o Options) GetColorPalette() chart.ColorPalette {
	if o.ColorPalette != nil {
		return o.ColorPalette
	}

	return defaultColorPalette
}

// GetTitleStyle returns custom or default title style
func (o Options) GetTitleStyle() chart.Style {
	if o.TitleStyle != nil {
		return *o.TitleStyle
	}

	return chart.StyleTextDefaults()
}

// GetTimeFormatter returns custom or default time formatter
func (o Options) GetTimeFormatter() chart.ValueFormatter {
	if o.TimeFormat != "" {
		return chart.TimeValueFormatterWithFormat(o.TimeFormat)
	}

	return chart.TimeValueFormatterWithFormat(defaultTimeFormat)
}

// GetHeight returns height
func (o Options) GetHeight() int {
	return o.Height
}

// GetWidth returns width
func (o Options) GetWidth() int {
	return o.Width
}

// GetTitle returns title
func (o Options) GetTitle() string {
	return o.Title
}
