package services

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func GeneratePortfolioChart(dates []string, values []float64, outputPath string) error {
	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Portfolio Performance"}),
		charts.WithXAxisOpts(opts.XAxis{Type: "category"}),
	)
	line.SetXAxis(dates).
		AddSeries("Portfolio Value", generateLineItems(values)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(true)}))

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return line.Render(f)
}

func generateLineItems(values []float64) []opts.LineData {
	items := make([]opts.LineData, 0)
	for _, v := range values {
		items = append(items, opts.LineData{Value: v})
	}
	return items
}
