package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"intro_web/src/model"
	"net/http"
)

func main() {
	// 处理静态资源
	// http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("views/static"))))

	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/success", successHandler)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/check", checkHandler)
	http.ListenAndServe(":8080", nil)
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("username")
	pwd := r.PostFormValue("password")
	if name == "admin" && pwd == "0000" {
		w.Header().Set("Location", "http://localhost:8080/success")
		w.WriteHeader(302)
	}
}

// 初始页面
func indexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("src/main/index.html"))
	t.Execute(w, "")
}

// 登录成功的页面
func successHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("src/main/success.html")
	t.Execute(w, "admin")
}

// 将员工信息一json返回
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	employee := model.Employee{Id: "23010119751201312X"}
	err := employee.Info()
	json, err := json.Marshal(employee)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	w.Write(json)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "请求地址", r.URL.Path)
	fmt.Fprintln(w, "查询字符串是", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头中Accept-Encoding信息", r.Header.Get("Accept-Encoding")) // gzip, deflate, br
	fmt.Fprintln(w, "请求头中Accept-Encoding信息", r.Header["Accept-Encoding"])     // [gzip, deflate, br]

	//body := make([]byte, r.ContentLength)
	//r.Body.Read(body)
	//fmt.Fprintln(w, "请求体中的内容有: ", string(body))

	//r.ParseForm()
	//fmt.Fprintln(w, "请求参数有:", r.Form)

	fmt.Fprintln(w, "请求参数有:", r.PostFormValue("username"))
}
