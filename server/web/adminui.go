// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

var indexTpl = `
{{define "content"}}
<h1>HBCACSMP服务器运行状况仪表板</h1>
<p>
详细信息请参考我们的文档:
</p>
<p>
<a target="_blank" href="https://www.hbca.org.cn/">湖北CA</a>
</p>
{{.Content}}
{{end}}`

var profillingTpl = `
{{define "content"}}
<h1>{{.Title}}</h1>
<pre id="content">
<div>{{.Content}}</div>
</pre>
{{end}}`

var defaultScriptsTpl = ``

var gcAjaxTpl = `
{{define "scripts"}}
<script type="text/javascript">
	var app = app || {};
(function() {
	app.$el = $('#content');
	app.getGc = function() {
		var that = this;
		$.ajax("/admin/prof?command=gc%20summary&format=json").done(function(data) {
			that.$el.append($('<p>' + data.Content + '</p>'));
		});
	};
	$(document).ready(function() {
		setInterval(function() {
			app.getGc();
		}, 3000);
	});
})();
</script>
{{end}}
`

var qpsTpl = `{{define "content"}}
<h1>请求统计</h1>
<table class="table table-striped table-hover ">
	<thead>
	<tr>
	{{range .Content.Fields}}
		<th>
		{{.}}
		</th>
	{{end}}
	</tr>
	</thead>

	<tbody>
	{{range $i, $elem := .Content.Data}}

	<tr>
	    <td>{{index $elem 0}}</td>
	    <td>{{index $elem 1}}</td>
	    <td>{{index $elem 2}}</td>
	    <td data-order="{{index $elem 3}}">{{index $elem 4}}</td>
	    <td data-order="{{index $elem 5}}">{{index $elem 6}}</td>
	    <td data-order="{{index $elem 7}}">{{index $elem 8}}</td>
	    <td data-order="{{index $elem 9}}">{{index $elem 10}}</td>
	</tr>
	{{end}}
	</tbody>

</table>
{{end}}`

var configTpl = `
{{define "content"}}
<h1>配置</h1>
<pre>
{{range $index, $elem := .Content}}
{{$index}}={{$elem}}
{{end}}
</pre>
{{end}}
`

var routerAndFilterTpl = `{{define "content"}}


<h1>{{.Title}}</h1>

{{range .Content.Methods}}

<div class="panel panel-default">
<div class="panel-heading lead success"><strong>{{.}}</strong></div>
<div class="panel-body">
<table class="table table-striped table-hover ">
	<thead>
	<tr>
	{{range $.Content.Fields}}
		<th>
		{{.}}
		</th>
	{{end}}
	</tr>
	</thead>

	<tbody>
	{{$slice := index $.Content.Data .}}
	{{range $i, $elem := $slice}}

	<tr>
		{{range $elem}}
			<td>
			{{.}}
			</td>
		{{end}}
	</tr>

	{{end}}
	</tbody>

</table>
</div>
</div>
{{end}}


{{end}}`

var tasksTpl = `{{define "content"}}

<h1>{{.Title}}</h1>

{{if .Message }}
{{ $messageType := index .Message 0}}
<p class="message
{{if eq "error" $messageType}}
bg-danger
{{else if eq "success" $messageType}}
bg-success
{{else}}
bg-warning
{{end}}
">
{{index .Message 1}}
</p>
{{end}}


<table class="table table-striped table-hover ">
<thead>
<tr>
{{range .Content.Fields}}
<th>
{{.}}
</th>
{{end}}
</tr>
</thead>

<tbody>
{{range $i, $slice := .Content.Data}}
<tr>
	{{range $slice}}
	<td>
	{{.}}
	</td>
	{{end}}
	<td>
	<a class="btn btn-primary btn-sm" href="/admin/task?taskname={{index $slice 0}}">立即执行</a>
	</td>
</tr>
{{end}}
</tbody>
</table>

{{end}}`

var healthCheckTpl = `
{{define "content"}}

<h1>{{.Title}}</h1>
<table class="table table-striped table-hover ">
<thead>
<tr>
{{range .Content.Fields}}
	<th>
	{{.}}
	</th>
{{end}}
</tr>
</thead>
<tbody>
{{range $i, $slice := .Content.Data}}
	{{ $header := index $slice 0}}
	{{ if eq "success" $header}}
	<tr class="success">
	{{else if eq "error" $header}}
	<tr class="danger">
	{{else}}
	<tr>
	{{end}}
		{{range $j, $elem := $slice}}
		{{if ne $j 0}}
		<td>
		{{$elem}}
		</td>
		{{end}}
		{{end}}
		<td>
		{{$header}}
		</td>
	</tr>
{{end}}

</tbody>
</table>
{{end}}`

// The base dashboardTpl
var dashboardTpl = `
<!DOCTYPE html>
<html lang="en">
<head>
<!-- Meta, title, CSS, favicons, etc. -->
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1">

<title>

欢迎使用HBCACSMP服务器运行状况仪表板

</title>

<link href="//maxcdn.bootstrapcdn.com/bootstrap/3.2.0/css/bootstrap.min.css" rel="stylesheet">
<link href="//cdn.datatables.net/plug-ins/725b2a2115b/integration/bootstrap/3/dataTables.bootstrap.css" rel="stylesheet">

<style type="text/css">
ul.nav li.dropdown:hover > ul.dropdown-menu {
	display: block;    
}
#logo {
	width: 102px;
	height: 32px;
	margin-top: 5px;
}
.message {
	padding: 15px;
}
</style>

</head>
<body>

<header class="navbar navbar-default navbar-static-top bs-docs-nav" id="top" role="banner">
<div class="container">
<div class="navbar-header">
<button class="navbar-toggle" type="button" data-toggle="collapse" data-target=".bs-navbar-collapse">
<span class="sr-only">开关导航</span>
<span class="icon-bar"></span>
<span class="icon-bar"></span>
<span class="icon-bar"></span>
</button>

<a href="/admin/">
<img id="logo" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAN8AAAAyCAIAAABTQZknAAAIiklEQVR4Ae2bMYhWRxDHX+QsFPECFl6hJKJYaHEWWohRg0IuVYRcfxbpEmJAIZ0WSRcwECUWgRRekc5ArHKBHChJLDRwV9iIokGLsxByIlp4RX6fkxvnZt/be+9733f3ncxyfOzOm52dnf3vzM6+d0URJSwQFggLhAXCAmGBsEBYICwQFggLhAXCAmGBsEBYICwQFggLhAXCAmGBsEBYICwQFggLhAXCAmGBuhZ4qy7jm8236e1i/FQxerQzyz+vFlOXi2f/vtkzXhOzW0TnvqPFt78v0fj08WLm2hJKaaOq4/RCKXst4t3ZYmqyuPJdCfOHJ4sPJgoGbVTycwGaP90t+NWCAnQJgKpBVqmybpXGzQ67a7T47Hzxw99LEEOPL3/s/DWFZnaozsNPz/uBUABXGmW1LTCQ6BSjABGwqAW84jh7XkbeLRcLOq037fm4IbCGBQYYnWj/3okCjFLA0PgXNabTnGXibHkfoDnWh81QPlhQyy0w2OhE59H3O4qPTXR+e16AYMYfj3/e8wGdwA0bNhw8eJBfoY+MjIyNja1fv96xlTYd87Fjx3bs2FHKaYnwMKKltKwjcO/evakQRtm9e7ejb968ef/+/Y6YaQ5lnrV6dPmrbrpvGvY+Egrl0AkvjZzp2bwnVrXn/il/kj9cStD/9XJ5315QX7x48eTJkwMHDszMzDx9+tSKZCEtjO7cuXP//n3L0Kc6aEtR5cZ6+fLl9PS0EOfm5tDz+fPnddRjjuw90Hz79m121+joq8DopC82Z2dnBwydaFYawSW+L+pd/PFL8f0ZbXVZwXHm0Yncj08V/UQnIwA7fCdYdOiUSQECoGBh2miyDuLaFw+tdSogDCgIBZBV4QxpAAtt0Vm7s8Hoi0dEz0ePHimdChTblDr7UNHMuMogqsp8ldg3dOoI7Stpkv7X1fZSO4dal/fMPehQLJFdweh1btZaKMTq4q4UMcRohOFduhaJ/8MzgQORYJdcHt24cUOF5x2YsqEhfcGT7StPcf/gsuaBBDSD/prMawGdaiGtAKP2Jc2Hzo0Xu/YtuShglIlzxczx9qOVSsDlsLSsFt6IIsFOwIQvSbvgdUrpyrlt27bt27fjtAA3OKgJAu1eWhGXiSh2kfV2lrnRXqryzVag1KvRyaW3vDtJOwklf7J0d/tVQhzd+i19lKrBbXnLQjLEsdIWHCRi+QO19hG+k2ZP9oMd7lUdl4Pr2rhxY9Xqih+FV2Bh/VZ6bsO9gXVElR4SksHrEiTpuX79uusAXlW9FLgcAJiX7WL5mYgoiQTAWoXXanRmklkZM4/ONBxbTevXZ6+VbJL2b3HYe65MLqZxk18n7vNs8c0njr0nTcEcDo+VO3z4ML+IZcGgy4I1OnfifUuXWTGkOuspQihVHlH5SXq0rhU89NTUFM1UPhOhbNmyRZmpCL+cLy09U69GZ6bTij3Ck+HSUt/ZUgF2jts8MpCIJeVyb4/YqEC2P+4TZCg4BF64wNTxtJlxzXMneEpxJuMCKVy104GdgMKOqPxgcWhoiIlwwCjlKSWqAnhW/Otgo7N0Bu2JZOKu/HzhNQHHfOVCcfLcawo1LlzzsWIJ95psqC902kuKTex29EwTNEvgxn26RD7Ti0d2I9Fcl+de5aekzO4uqb1CHCLJ1m3BKbprI9DpDg99e7GJdyE3wm9ZjahD5xdfQhRmsd3TZZtHjhxJvd2yvVoykNRLNCcz4wRMyehAiEhn7RRYKd8JAqquxK1GZEUOjuk9vOXvop6m6kRtV4Am8d2evFGMF5uln025vg2bLCfgw2/RjwWjUCHWQycmSi7i7jsBLmyZhUeCRFU5L2q4VNWanju1Y6ZC8oTOXGOBUUYXl0ld/ajtC3xhfvjwoSWm9Wp05r86SyXlKXwRVzMypp/e3UtiStdJNCCzmENnAWKqPJB1nLzY7AM6h4eHNcXWtIaVw6GmKQ6IJIVivXn0+PHjqqtKfBJFD3w2XAIXhNjcX6bOBmiT7O/Zs4fhbt68ubCwAPJQT/YbGAW1djgUY0SmALNqmJpfKIMd2UVHF2QhjrxTNZ9l6OnLoTSIi4g03LMlHF6XGazWY4BIBIQVDAkccY2sKETNOWRF4QGa0JXTDsBKK9vWrVt5ND8/D+5JqwUoyqxsSgGvwEWbrpLyWwY6wsDot27d4pcNY99qMgWeSkCQXjTRSpitnNJ6te8sZV8VYnok2Lmvm/c3OE6HTnAPOqsKqZKDY69fbAJNsKg5uygCVlhg1o8mQZnlBF7CI8RSfUE2DgzXKE+BRalnQg7yXWSnC05OXThNtOLkKqKQI9tGmu6Xp2wYyZnkAO2UJNZDp5dsNhSgMCnVoXOaWfxexB5CGHRNoPNBJ/6CLS1k0F0E2fTVJYfL1DHrKHLNZO+eev1iUxyJgxGrQhEtcJOqjqvIMivRNZXuKoxl46x7qk3Y5C5TKVUVdo7AUTYGdeeqxVMCUCooKXLgqSN/pdAJnupcW1oIWnu4HAWU8OV8Bli2L/VLp/9/CeTokg8Rsu1nzpYnPUL09MVmTUhZjQa2bjeVUxJcZraZY7bNlUInCOCviyIpEZ99uCDrUvu8ZECP43QKcIvE4ZJCFm8dZF4UnMiRjnnOeNraAoOdFeEd5fsgfGdLQKQ38L9NdqwnwG1kx/ROqlH3YK5tgcFG56Uzr8N3mzfdHCqcdwT0gvsurtnx4s4N1zZ3MDaywACjEzjaVziACUr9s6Y1Q/qPH/LNB47TZfG2V6aeCswwx6NuLbB47tQYqoJq4qCqo3gmldaocm+muHKxJJQDVsTy9ujQR43kdZjttZTqzOGVxLyLsnO0i07RJSwQFggLhAXCAmGBsEBYICwQFggLhAXCAmGBsEBYICwQFggLhAXCAmGBsEBYICwQFggLhAXCAmGBsED/LPAfu4CkZI99uUgAAAAASUVORK5CYII="/>
</a>

</div>
<nav class="collapse navbar-collapse bs-navbar-collapse" role="navigation">
<ul class="nav navbar-nav">
<li>
<a href="/admin/qps">
请求统计
</a>
</li>
<li>

<li class="dropdown">
<a href="#" class="dropdown-toggle disabled" data-toggle="dropdown">性能状况<span class="caret"></span></a>
<ul class="dropdown-menu" role="menu">

<li><a href="/admin/prof?command=lookup goroutine">查看协程情况</a></li>
<li><a href="/admin/prof?command=lookup heap">查看堆栈情况</a></li>
<li><a href="/admin/prof?command=lookup threadcreate">查看线程创建</a></li>
<li><a href="/admin/prof?command=lookup block">查看块情况</a></li>
<li><a href="/admin/prof?command=get cpuprof">生成CPU分析文件</a></li>
<li><a href="/admin/prof?command=get memprof">生成内存分析文件</a></li>
<li><a href="/admin/prof?command=gc summary">内存回收统计</a></li>

</ul>
</li>

<li>
<a href="/admin/healthcheck">
健康检查
</a>
</li>

<li>
<a href="/admin/task" class="dropdown-toggle disabled" data-toggle="dropdown">计划任务情况</a>
</li>

<li class="dropdown">
<a href="#" class="dropdown-toggle disabled" data-toggle="dropdown">配置情况<span class="caret"></span></a>
<ul class="dropdown-menu" role="menu">
<li><a href="/admin/listconf?command=conf">所有配置</a></li>
<li><a href="/admin/listconf?command=router">所有路由</a></li>
<li><a href="/admin/listconf?command=filter">所有过滤器</a></li>
</ul>
</li>
</ul>
</nav>
</div>
</header>

<div class="container">
{{template "content" .}}
</div>

<script src="//code.jquery.com/jquery-1.11.1.min.js"></script>
<script src="//maxcdn.bootstrapcdn.com/bootstrap/3.2.0/js/bootstrap.min.js"></script>
<script src="//cdn.datatables.net/1.10.2/js/jquery.dataTables.min.js"></script>
<script src="//cdn.datatables.net/plug-ins/725b2a2115b/integration/bootstrap/3/dataTables.bootstrap.js
"></script>

<script type="text/javascript">
$(document).ready(function() {
    $('.table').dataTable();
});
</script>
{{template "scripts" .}}
</body>
</html>
`
