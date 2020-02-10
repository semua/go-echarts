package charts

import (
	// "html/template"
	"io"
)

// CustomChart.
type CustomChart struct {
	BaseOpts
}

func (CustomChart) chartType() string { return ChartType.CustomChart }

func (c *CustomChart) SetCustomJson(customJson string) *CustomChart {
	c.BaseOpts.CustomOptionJson = customJson
	return c
}
func (c *CustomChart) SetCustomVar(customVar string) *CustomChart {
	c.BaseOpts.CustomVar = customVar
	return c
}

func (c *CustomChart) SetCustomOption(customOption bool) *CustomChart {
	c.BaseOpts.CustomOption = customOption
	return c
}

func (c *CustomChart) validateOpts() {
	c.validateAssets(c.BaseOpts.AssetsHost)
}

// CustomChart creates a new custom chart.
func NewCustomChart(routers ...RouterOpts) *CustomChart {
	chart := new(CustomChart)
	chart.initBaseOpts(routers...)
	chart.SetCustomOption(true)
	return chart
}

// Render renders the chart and writes the output to given writers.
func (c *CustomChart) Render(w ...io.Writer) error {
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
