package main

import (
	"github.com/go-analyze/charts"

	"test/pkg"
)

func main() {
	values, xAxisLabels, labels, err := pkg.PopulateData()
	if err != nil {
		panic(err)
	}

	theme := charts.GetTheme(charts.ThemeAnt)
	p := charts.NewPainter(charts.PainterOptions{
		OutputFormat: charts.ChartOutputPNG,
		Width:        600,
		Height:       400,
		Theme:        theme,
	})

	if err := pkg.RenderLineChart(p, values, xAxisLabels); err != nil {
		panic(err)
	}

	pkg.RenderCustomText(p, theme, labels)

	if err := pkg.WritePainter(p, "line-chart.png"); err != nil {
		panic(err)
	}
}
