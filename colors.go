package apmcharts

import (
	"github.com/wcharczuk/go-chart/v2/drawing"
)

// Color Schema
var colorSchema = []drawing.Color{
	{51, 153, 204, 255},
	{146, 127, 185, 255},
	{255, 204, 0, 255},
	{87, 183, 154, 255},
	{190, 83, 187, 255},
	{221, 132, 81, 255},
	{57, 105, 179, 255},
	{190, 208, 23, 255},
	{137, 52, 164, 255},
	{59, 203, 203, 255},
	{110, 105, 204, 255},
	{80, 147, 31, 255},
	{200, 107, 116, 255},
	{252, 175, 43, 255},
	{46, 176, 222, 255},
	{198, 140, 205, 255},
	{69, 117, 87, 255},
	{204, 60, 113, 255},
	{152, 80, 131, 255},
	{167, 179, 66, 255},
	{119, 74, 164, 255},
	{95, 59, 131, 255},
	{102, 101, 104, 255},
	{148, 145, 150, 255},
	{52, 156, 80, 255},
	{65, 196, 100, 255},
	{229, 162, 38, 255},
	{255, 181, 43, 255},
	{188, 43, 60, 255},
	{235, 54, 76, 255},
	{255, 0, 153, 255},
	{212, 236, 249, 255},
	{234, 246, 252, 255},
	{9, 83, 191, 255},
	{0, 107, 194, 255},
}

// colorSchemaByStatusCode color Schema
var colorSchemaByStatusCode = map[uint32]drawing.Color{
	200: {255, 237, 160, 255},
	201: {254, 217, 118, 255},
	204: {254, 178, 76, 255},
	403: {253, 141, 60, 255},
	426: {252, 78, 42, 255},
	500: {227, 26, 28, 255},
	503: {189, 0, 38, 255},
	504: {128, 0, 38, 255},
}

// Red stroke color
var Red = drawing.Color{R: 235, G: 54, B: 75, A: 255}

// RedFill color
var RedFill = drawing.Color{R: 235, G: 54, B: 75, A: 20}

// Green stroke color
var Green = drawing.Color{R: 64, G: 196, B: 99, A: 255}

// GreenFill color
var GreenFill = drawing.Color{R: 64, G: 196, B: 99, A: 20}

// Blue stroke color
var Blue = drawing.Color{R: 0, G: 185, B: 217, A: 255}

// BlueFill color
var BlueFill = drawing.Color{R: 0, G: 185, B: 217, A: 20}

// Yellow stroke color
var Yellow = drawing.Color{R: 255, G: 204, B: 0, A: 255}

// YellowFill color
var YellowFill = drawing.Color{R: 255, G: 204, B: 0, A: 20}
