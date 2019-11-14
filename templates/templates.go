package templates

import (
	// "io/ioutil"
	// "log"
	"net/http"
	// "github.com/gobuffalo/packr"
)

var BaseTpl = `{{- define "base" }}
<div class="container">
    <div class="item" id="{{ .ChartID }}"
         style="width:{{ .InitOpts.Width }};height:{{ .InitOpts.Height }};"></div>
</div>
<script type="text/javascript">
    "use strict";
    var myChart___x__{{ .ChartID }}__x__ = echarts.init(document.getElementById('{{ .ChartID }}'), "{{ .Theme }}");
    var option___x__{{ .ChartID }}__x__ = {
        title: {{ .TitleOpts  }},
        tooltip: {{ .TooltipOpts }},
        legend: {{ .LegendOpts }},
    {{- if .HasGeo }}
        geo: {{ .GeoComponentOpts }},
    {{- end }}
    {{- if .HasRadar }}
        radar: {{ .RadarComponentOpts }},
    {{- end }}
    {{- if .HasParallel }}
        parallel: {{ .ParallelComponentOpts }},
        parallelAxis: {{ .ParallelAxisOpts }},
    {{- end }}
    {{- if .HasSingleAxis }}
        singleAxis: {{ .SingleAxisOpts }},
    {{- end }}
    {{- if .ToolboxOpts.Show }}
        toolbox: {{ .ToolboxOpts }},
    {{- end }}
    {{- if gt .DataZoomOptsList.Len 0 }}
        dataZoom:{{ .DataZoomOptsList }},
    {{- end }}
    {{- if gt .VisualMapOptsList.Len 0 }}
        visualMap:{{ .VisualMapOptsList }},
    {{- end }}
    {{- if .HasXYAxis }}
        xAxis: {{ .XAxisOptsList }},
        yAxis: {{ .YAxisOptsList }},
    {{- end }}
    {{- if .Has3DAxis }}
        xAxis3D: {{ .XAxis3D }},
        yAxis3D: {{ .YAxis3D }},
        zAxis3D: {{ .ZAxis3D }},
        grid3D: {{ .Grid3D }},
    {{- end }}
        series: [
        {{ range .Series }}
        {{- . }},
        {{ end -}}
        ],
    {{- if eq .Theme "white" }}
        color: {{ .Colors }},
    {{- end }}
    {{- if ne .BackgroundColor "" }}
        backgroundColor: {{ .BackgroundColor }}
    {{- end }}
    };
    myChart___x__{{ .ChartID }}__x__.setOption(option___x__{{ .ChartID }}__x__);

    {{- range .JSFunctions.Fns }}
		{{ . }}
	{{- end }}
</script>
{{ end }}

`
var ChartTpl = `{{- define "chart" }}
<!DOCTYPE html>
<html>
{{- template "header" . }}
<body>
{{- template "routers" . }}
{{- template "base" . }}
<style>
    .container {display: flex;justify-content: center;align-items: center;background-color:white;}
    .item {margin: auto;}
</style>
</body>
</html>
{{ end }}
`
var HeaderTpl = `{{ define "header" }}
<head>
    <meta charset="utf-8">
    <title>{{ .PageTitle }}</title>
{{- range .JSAssets.Values }}
    <script src="{{ . }}"></script>
{{- end }}
{{- range .CSSAssets.Values }}
    <link href="{{ . }}" rel="stylesheet">
{{- end }}
</head>
{{ end }}
`
var PageTpl = `{{- define "page" }}
<!DOCTYPE html>
<html>
{{- template "header" . }}
<body>
{{- template "routers" . }}
{{- range .Charts }}
    {{ template "base" . }}
    <br/>
{{- end }}
<style>
    .container {display: flex;justify-content: center;align-items: center;}
    .item {margin: auto;}
    body {margin: 0 !important;}
</style>
</body>
</html>
{{ end }}
`
var RoutersTpl = `{{- define "routers" }}
<div class="select" style="margin-right:10px; margin-top:10px; position:fixed; right:10px;">
{{- if gt .Routers.Len 0}}
    <select onchange="window.location.href=this.options[this.selectedIndex].value">
    {{- range .Routers }}
        <option value="{{ .URL }}">{{ .Text }}</option>
    {{- end }}
    </select>
{{- end -}}
</div>
{{ end }}
`

func init() {
	/*
		box := packr.NewBox(".")

		var err error
		BaseTpl, err = box.FindString("base.html")
		ChartTpl, err = box.FindString("chart.html")
		HeaderTpl, err = box.FindString("header.html")
		PageTpl, err = box.FindString("page.html")
		RoutersTpl, err = box.FindString("routers.html")

		if err != nil {
			log.Printf("packr load templates error: %v", err)
		}
	*/
}

func LoadTemplates(loader http.FileSystem) {
	/*
		var err error
		BaseTpl, err = _bytesToString(loader.Open("base.html"))
		ChartTpl, err = _bytesToString(loader.Open("chart.html"))
		HeaderTpl, err = _bytesToString(loader.Open("header.html"))
		PageTpl, err = _bytesToString(loader.Open("page.html"))
		RoutersTpl, err = _bytesToString(loader.Open("routers.html"))

		if err != nil {
			log.Fatal(err)
		}
	*/
}

// func _bytesToString(file http.File, err error) (string, error) {
// 	content, err := ioutil.ReadAll(file)
// 	return string(content), err
// }
