package pkg

import (
	"os"

	"github.com/go-analyze/charts"
)

func ptr[T any](x T) *T {
	return &x
}

func RenderLineChart(p *charts.Painter, values [][]float64, xAxisLabels []string) error {
	opt := charts.NewLineChartOptionWithData(values)
	opt.Padding = charts.NewBoxEqual(20)
	opt.Title.Text = "Canon RF Zoom Lenses"
	opt.Title.Offset = charts.OffsetCenter
	opt.Title.FontStyle.FontSize = 16

	opt.XAxis.Labels = xAxisLabels
	opt.XAxis.Unit = 40
	opt.XAxis.LabelCount = 10
	opt.XAxis.LabelRotation = charts.DegreesToRadians(45)
	opt.XAxis.BoundaryGap = ptr(true)
	opt.XAxis.FontStyle = charts.NewFontStyleWithSize(6.0)
	opt.YAxis = []charts.YAxisOption{
		{
			Show:          ptr(false), // disabling in favor of manually printed y-values
			Min:           ptr(1.4),
			Max:           ptr(8.0),
			LabelCount:    4,
			SpineLineShow: ptr(true),
			FontStyle:     charts.NewFontStyleWithSize(8.0),
		},
	}
	opt.Legend.Show = ptr(false)
	opt.Symbol = charts.SymbolNone
	opt.LineStrokeWidth = 1.5

	return p.LineChart(opt)
}

func RenderCustomText(p *charts.Painter, theme charts.ColorPalette, labels []string) {
	fontStyle := charts.FontStyle{
		FontSize:  12,
		FontColor: charts.ColorBlack,
		Font:      charts.GetDefaultFont(),
	}

	fontStyle.FontColor = theme.GetSeriesColor(0)
	p.Text(labels[0], 420, 84, 0, fontStyle)

	fontStyle.FontColor = theme.GetSeriesColor(1)
	p.Text(labels[1], 45, 284, 0, fontStyle)

	fontStyle.FontColor = theme.GetSeriesColor(2)
	p.Text(labels[2], 140, 230, 0, fontStyle)

	fontStyle.FontColor = theme.GetSeriesColor(3)
	p.Text(labels[3], 160, 155, 0, fontStyle)

	fontStyle.FontSize = 8
	fontStyle.FontColor = theme.GetSeriesColor(0)
	p.Text("f/4.5", 42, 220, 0, fontStyle)
	p.Text("f/5.0", 105, 196, 0, fontStyle)
	p.Text("f/6.3", 370, 137, 0, fontStyle)
	p.Text("f/7.1", 570, 100, 0, fontStyle)

	fontStyle.FontColor = theme.GetSeriesColor(1)
	p.Text("f/2.8", 5, 298, 0, fontStyle)

	fontStyle.FontColor = theme.GetSeriesColor(2)
	p.Text("f/4.0", 40, 244, 0, fontStyle)

	fontStyle.FontColor = theme.GetSeriesColor(3)
	p.Text("f/5.6", 92, 168, 0, fontStyle)
}

func WritePainter(p *charts.Painter, filename string) error {
	buf, err := p.Bytes()
	if err != nil {
		return err
	}
	return os.WriteFile(filename, buf, 0600)
}
