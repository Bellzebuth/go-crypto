package services

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func GeneratePieChart(assets []string, values []float64, outputPath string) error {
	pie := charts.NewPie()

	items := make([]opts.PieData, 0)
	for i, asset := range assets {
		items = append(items, opts.PieData{Name: asset, Value: values[i]})
	}

	pie.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "Portfolio Distribution"})).
		AddSeries("Portfolio", items)

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return pie.Render(f)
}
