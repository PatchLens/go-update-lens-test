package pkg

import (
	"testing"

	"github.com/go-analyze/charts"
	"github.com/stretchr/testify/assert"
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

func TestRenderLineChartSizeFail(t *testing.T) {
	values, xAxisLabels, _, err := PopulateData()
	require.NoError(t, err)

	p := charts.NewPainter(charts.PainterOptions{
		OutputFormat: charts.ChartOutputSVG,
		Width:        600,
		Height:       400,
	})

	err = RenderLineChart(p, values, xAxisLabels)
	require.NoError(t, err)

	b, err := p.Bytes()
	require.NoError(t, err)

	assert.Equal(t, 13081, len(b))
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

func TestRenderCustomTextSizeFail(t *testing.T) {
	_, _, labels, err := PopulateData()
	require.NoError(t, err)

	p := charts.NewPainter(charts.PainterOptions{
		OutputFormat: charts.ChartOutputSVG,
		Width:        600,
		Height:       400,
	})

	RenderCustomText(p, charts.GetDefaultTheme(), labels)

	b, err := p.Bytes()
	require.NoError(t, err)

	assert.Equal(t, 1608, len(b))
}
