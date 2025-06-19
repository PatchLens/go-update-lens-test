package pkg

import (
	"testing"

	"github.com/go-analyze/charts"
	"github.com/stretchr/testify/require"
)

func TestRenderLineChartBasic(t *testing.T) {
	values, xAxisLabels, _, err := PopulateData()
	require.NoError(t, err)

	p := charts.NewPainter(charts.PainterOptions{
		OutputFormat: charts.ChartOutputSVG,
		Width:        600,
		Height:       400,
	})

	err = RenderLineChart(p, values, xAxisLabels)
	require.NoError(t, err)
}

func TestRenderCustomTextBasic(t *testing.T) {
	_, _, labels, err := PopulateData()
	require.NoError(t, err)

	p := charts.NewPainter(charts.PainterOptions{
		OutputFormat: charts.ChartOutputSVG,
		Width:        600,
		Height:       400,
	})

	RenderCustomText(p, charts.GetDefaultTheme(), labels)
}
