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

var loginTpl = `
{{define "content"}}
<h1>HBCACSMP服务器运行状况仪表板-登陆</h1>
<p>
<form class="login-form" method="post" novalidate url="/admin/login">
        <div class="input-content">
            <div style="margin-top: 16px">
                <input type="password"
                       autocomplete="off" placeholder="运维密码" name="password" required maxlength="32"/>
				<button type="submit" class="enter-btn" >登录</button>
            </div>
        </div>
</form>

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
<html lang="zh">
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
	width: 32;
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
<img id="logo" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAJAAAACQCAMAAAGnnyIcAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAyJpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuMy1jMDExIDY2LjE0NTY2MSwgMjAxMi8wMi8wNi0xNDo1NjoyNyAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENTNiAoV2luZG93cykiIHhtcE1NOkluc3RhbmNlSUQ9InhtcC5paWQ6RDYwQzZEM0E3NDQzMTFFQzkwMjNGMjQwQTk1MkMyQTkiIHhtcE1NOkRvY3VtZW50SUQ9InhtcC5kaWQ6RDYwQzZEM0I3NDQzMTFFQzkwMjNGMjQwQTk1MkMyQTkiPiA8eG1wTU06RGVyaXZlZEZyb20gc3RSZWY6aW5zdGFuY2VJRD0ieG1wLmlpZDpENjBDNkQzODc0NDMxMUVDOTAyM0YyNDBBOTUyQzJBOSIgc3RSZWY6ZG9jdW1lbnRJRD0ieG1wLmRpZDpENjBDNkQzOTc0NDMxMUVDOTAyM0YyNDBBOTUyQzJBOSIvPiA8L3JkZjpEZXNjcmlwdGlvbj4gPC9yZGY6UkRGPiA8L3g6eG1wbWV0YT4gPD94cGFja2V0IGVuZD0iciI/PtuDBzEAAABgUExURTNVkg84g83V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"/>
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
    $('.table').dataTable({
  
   language: {
       "sProcessing": "处理中...",
       "sLengthMenu": "显示 _MENU_ 项结果",
       "sZeroRecords": "没有匹配结果",
       "sInfo": "显示第 _START_ 至 _END_ 项结果，共 _TOTAL_ 项",
       "sInfoEmpty": "显示第 0 至 0 项结果，共 0 项",
       "sInfoFiltered": "(由 _MAX_ 项结果过滤)",
       "sInfoPostFix": "",
       "sSearch": "搜索:",
       "sUrl": "",
       "sEmptyTable": "表中数据为空",
       "sLoadingRecords": "载入中...",
       "sInfoThousands": ",",
       "oPaginate": {
           "sFirst": "首页",
           "sPrevious": "上页",
           "sNext": "下页",
           "sLast": "末页"
       },
       "oAria": {
           "sSortAscending": ": 以升序排列此列",
           "sSortDescending": ": 以降序排列此列"
       }
   }
});
});
</script>
{{template "scripts" .}}
</body>
</html>
`
