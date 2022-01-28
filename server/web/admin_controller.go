// Copyright 2020
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/beego/beego/v2/core/admin"
)

type adminController struct {
	Controller
	servers   []*HttpServer
	curUserId string
}

func (c *adminController) Prepare() {
	if c.Ctx.Request.URL.Path != "/admin/login" {
		c.checkLogin()
	}
}

func (c *adminController) checkLogin() {
	if c.curUserId != "admin" {
		//登录页面地址
		urlstr := "/admin/login"
		//登录成功后返回的址为当前
		c.Redirect(urlstr, 302)
		c.StopRun()
	}
}

func (a *adminController) registerHttpServer(svr *HttpServer) {
	a.servers = append(a.servers, svr)
}

// ProfIndex is a http.Handler for showing profile command.
// it's in url pattern "/prof" in admin module.
func (a *adminController) ProfIndex() {
	rw, r := a.Ctx.ResponseWriter, a.Ctx.Request
	r.ParseForm()
	command := r.Form.Get("command")
	if command == "" {
		return
	}

	var (
		format = r.Form.Get("format")
		data   = make(map[interface{}]interface{})
		result bytes.Buffer
	)
	admin.ProcessInput(command, &result)
	data["Content"] = template.HTMLEscapeString(result.String())

	if format == "json" && command == "gc summary" {
		dataJSON, err := json.Marshal(data)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		writeJSON(rw, dataJSON)
		return
	}

	data["Title"] = template.HTMLEscapeString(command)
	defaultTpl := defaultScriptsTpl
	if command == "gc summary" {
		defaultTpl = gcAjaxTpl
	}
	writeTemplate(rw, data, profillingTpl, defaultTpl)
}

func (a *adminController) PrometheusMetrics() {
	promhttp.Handler().ServeHTTP(a.Ctx.ResponseWriter, a.Ctx.Request)
}

// TaskStatus is a http.Handler with running task status (task name, status and the last execution).
// it's in "/task" pattern in admin module.
func (a *adminController) TaskStatus() {

	rw, req := a.Ctx.ResponseWriter, a.Ctx.Request

	data := make(map[interface{}]interface{})

	// Run Task
	req.ParseForm()
	taskname := req.Form.Get("taskname")
	if taskname != "" {
		cmd := admin.GetCommand("task", "run")
		res := cmd.Execute(taskname)
		if res.IsSuccess() {

			data["Message"] = []string{"success",
				template.HTMLEscapeString(fmt.Sprintf("%s 执行成功，目前状态为 %s",
					taskname, res.Content.(string)))}

		} else {
			data["Message"] = []string{"error", template.HTMLEscapeString(fmt.Sprintf("%s", res.Error))}
		}
	}

	// List Tasks
	content := make(M)
	resultList := admin.GetCommand("task", "list").Execute().Content.([][]string)
	var fields = []string{
		"任务名称",
		"任务周期",
		"任务状态",
		"上次执行时间",
		"操作",
	}

	content["Fields"] = fields
	content["Data"] = resultList
	data["Content"] = content
	data["Title"] = "计划任务"
	writeTemplate(rw, data, tasksTpl, defaultScriptsTpl)
}

func (a *adminController) AdminIndex() {
	// AdminIndex is the default http.Handler for admin module.
	// it matches url pattern "/".
	writeTemplate(a.Ctx.ResponseWriter, map[interface{}]interface{}{}, indexTpl, defaultScriptsTpl)
}

func (a *adminController) AdminLogin() {
	// AdminIndex is the default http.Handler for admin module.
	// it matches url pattern "/".
	writeTemplate(a.Ctx.ResponseWriter, map[interface{}]interface{}{}, loginTpl, defaultScriptsTpl)
}

func (a *adminController) DoLogin() {
	r := a.Ctx.Request
	r.ParseForm()
	password := r.Form.Get("password")
	if len(password) == 0 || password != "hbcaadminyw2022" {
		a.StopRun()
		return
	}
	a.curUserId = "admin"
	a.Redirect("/admin/", 302)
}

// Healthcheck is a http.Handler calling health checking and showing the result.
// it's in "/healthcheck" pattern in admin module.
func (a *adminController) Healthcheck() {
	heathCheck(a.Ctx.ResponseWriter, a.Ctx.Request)
}

func heathCheck(rw http.ResponseWriter, r *http.Request) {
	var (
		result     []string
		data       = make(map[interface{}]interface{})
		resultList = new([][]string)
		content    = M{
			"Fields": []string{"名称", "消息", "状态"},
		}
	)

	for name, h := range admin.AdminCheckList {
		if err := h.Check(); err != nil {
			result = []string{
				"error",
				template.HTMLEscapeString(name),
				template.HTMLEscapeString(err.Error()),
			}
		} else {
			result = []string{
				"success",
				template.HTMLEscapeString(name),
				"OK",
			}
		}
		*resultList = append(*resultList, result)
	}

	queryParams := r.URL.Query()
	jsonFlag := queryParams.Get("json")
	shouldReturnJSON, _ := strconv.ParseBool(jsonFlag)

	if shouldReturnJSON {
		response := buildHealthCheckResponseList(resultList)
		jsonResponse, err := json.Marshal(response)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		} else {
			writeJSON(rw, jsonResponse)
		}
		return
	}

	content["Data"] = resultList
	data["Content"] = content
	data["Title"] = "健康检查"

	writeTemplate(rw, data, healthCheckTpl, defaultScriptsTpl)
}

// QpsIndex is the http.Handler for writing qps statistics map result info in http.ResponseWriter.
// it's registered with url pattern "/qps" in admin module.
func (a *adminController) QpsIndex() {
	data := make(map[interface{}]interface{})
	data["Content"] = StatisticsMap.GetMap()

	// do html escape before display path, avoid xss
	if content, ok := (data["Content"]).(M); ok {
		if resultLists, ok := (content["Data"]).([][]string); ok {
			for i := range resultLists {
				if len(resultLists[i]) > 0 {
					resultLists[i][0] = template.HTMLEscapeString(resultLists[i][0])
				}
			}
		}
	}
	writeTemplate(a.Ctx.ResponseWriter, data, qpsTpl, defaultScriptsTpl)
}

// ListConf is the http.Handler of displaying all beego configuration values as key/value pair.
// it's registered with url pattern "/listconf" in admin module.
func (a *adminController) ListConf() {
	rw := a.Ctx.ResponseWriter
	r := a.Ctx.Request
	r.ParseForm()
	command := r.Form.Get("command")
	if command == "" {
		rw.Write([]byte("不支持的命令"))
		return
	}

	data := make(map[interface{}]interface{})
	switch command {
	case "conf":
		m := make(M)
		list("BConfig", BConfig, m)
		m["appConfigPath"] = template.HTMLEscapeString(appConfigPath)
		m["appConfigProvider"] = template.HTMLEscapeString(appConfigProvider)
		tmpl := template.Must(template.New("dashboard").Parse(dashboardTpl))
		tmpl = template.Must(tmpl.Parse(configTpl))
		tmpl = template.Must(tmpl.Parse(defaultScriptsTpl))

		data["Content"] = m

		tmpl.Execute(rw, data)

	case "router":
		content := BeeApp.PrintTree()
		content["Fields"] = []string{
			"路由路径",
			"方法",
			"控制器",
		}
		data["Content"] = content
		data["Title"] = "路由"
		writeTemplate(rw, data, routerAndFilterTpl, defaultScriptsTpl)
	case "filter":
		var (
			content = M{
				"Fields": []string{
					"路由路径",
					"过滤器方法",
				},
			}
		)

		filterTypeData := BeeApp.reportFilter()

		filterTypes := make([]string, 0, len(filterTypeData))
		for k := range filterTypeData {
			filterTypes = append(filterTypes, k)
		}

		content["Data"] = filterTypeData
		content["Methods"] = filterTypes

		data["Content"] = content
		data["Title"] = "过滤器"
		writeTemplate(rw, data, routerAndFilterTpl, defaultScriptsTpl)
	default:
		rw.Write([]byte("不支持的命令"))
	}
}

func writeTemplate(rw http.ResponseWriter, data map[interface{}]interface{}, tpls ...string) {
	tmpl := template.Must(template.New("dashboard").Parse(dashboardTpl))
	for _, tpl := range tpls {
		tmpl = template.Must(tmpl.Parse(tpl))
	}
	tmpl.Execute(rw, data)
}

func buildHealthCheckResponseList(healthCheckResults *[][]string) []map[string]interface{} {
	response := make([]map[string]interface{}, len(*healthCheckResults))

	for i, healthCheckResult := range *healthCheckResults {
		currentResultMap := make(map[string]interface{})

		currentResultMap["name"] = healthCheckResult[0]
		currentResultMap["message"] = healthCheckResult[1]
		currentResultMap["status"] = healthCheckResult[2]

		response[i] = currentResultMap
	}

	return response

}

// PrintTree print all routers
// Deprecated using BeeApp directly
func PrintTree() M {
	return BeeApp.PrintTree()
}
