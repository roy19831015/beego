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
<h1>CTYCSMP服务器运行状况仪表板</h1>
<p>
详细信息请参考我们的文档:
</p>
<p>
<a target="_blank" href="http://www.chutianyun.gov.cn/">楚天云</a>
</p>
{{.Content}}
{{end}}`

var loginTpl = `
{{define "content"}}
<h1>CTYCSMP服务器运行状况仪表板-登陆</h1>
<p>
<form class="login-form" novalidate>
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
<a target="_blank" href="http://www.chutianyun.gov.cn/">楚天云</a>
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

欢迎使用CTYCSMP服务器运行状况仪表板

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
<img id="logo" src="data:image/png;base64,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"/>
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
