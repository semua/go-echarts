package main

import (
	"log"
	"net/http"
	"os"

	"github.com/semua/go-echarts/charts"
)

func customChartBase() *charts.CustomChart {
	customChart := charts.NewCustomChart()
	// customChart.SetGlobalOptions(charts.TitleOpts{Title: "custom-示例图"}, charts.ToolboxOpts{Show: true})
	customChart.SetCustomOption(true)
	customChart.SetCustomVar(`var preclose = 28.45;`)
	customChart.SetCustomJson(`
	
	{
        title: {"text":"bar示例图",},
        tooltip: {},
        legend: {},
        xAxis: [{"data":["衬衫","牛仔裤","运动裤","袜子","冲锋衣","羊毛衫"],"splitArea":{"show":false,},"splitLine":{"show":false,}}],
        yAxis: [{"axisLabel":{"show":true},"splitArea":{"show":false,},"splitLine":{"show":false,}}],
        series: [
        {"name":"商家A","type":"bar","data":[20,30,40,10,24,35],"label":{"show":false},"emphasis":{"label":{"show":false},},"markLine":{"label":{"show":false}},"markPoint":{"label":{"show":false}},},
        {"name":"商家B","type":"bar","data":[35,14,25,60,44,23],"label":{"show":false},"emphasis":{"label":{"show":false},},"markLine":{"label":{"show":false}},"markPoint":{"label":{"show":false}},},
        ],
        color: ["#c23531","#2f4554","#61a0a8","#d48265","#91c7ae","#749f83","#ca8622","#bda29a","#6e7074","#546570","#c4ccd3"]
    }
	
	`)
	return customChart
}

func customChartHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("customChart")...)
	page.Add(
		customChartBase(),
	)
	f, err := os.Create(getRenderPath("customChart.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(w, f)
}
