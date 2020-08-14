package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)
var p *WorkerPool
type Score struct {
	Num int
}
type chanData struct {
	res http.ResponseWriter
	req *http.Request
	body string
}
func (s *chanData) Do() {
	//fmt.Println("res:", s.res)
	fmt.Println("req:", s.req)

	//defer r.Body.Close()
	fmt.Println(s.body)

	//err := r.ParseForm()
	//if err != nil {
	//	log.Fatal("parse form error ",err)
	//}
	//// 初始化请求变量结构
	//formData := make(map[string]interface{})
	//// 调用json包的解析，解析请求body
	//json.NewDecoder(r.Body).Decode(&formData)
	//for key,value := range formData {
	//	fmt.Println("key:", key, " => value :", value)
	//}
}
func (s *Score) Do() {
	fmt.Println("num:", s.Num)
	time.Sleep(1 * 1 * time.Second)
}

func main() {
	num := 100 * 100 * 20
	// debug.SetMaxThreads(num + 1000) //设置最大线程数
	// 注册工作池，传入任务
	// 参数1 worker并发个数
	p = NewWorkerPool(num)
	p.Run()
	http.HandleFunc("/hi", Router)
	http.HandleFunc("/api", Api)
	http.HandleFunc("/report", Report)
	http.HandleFunc("/test", Test)
	http.ListenAndServe("127.0.0.1:12306", nil)
	//datanum := 100 * 100 * 100 * 100
	//go func() {
	//	for i := 1; i <= datanum; i++ {
	//		sc := &Score{Num: i}
	//		p.JobQueue <- sc
	//	}
	//}()
	//
	//for {
	//	fmt.Println("runtime.NumGoroutine() :", runtime.NumGoroutine())
	//	time.Sleep(2 * time.Second)
	//}

}
func process(w  *chanData) {
	p.JobQueue <- w
}

func Router(resp http.ResponseWriter, request *http.Request) {
	con, err := ioutil.ReadAll(request.Body) //获取post的数据
	if(err!=nil){
		fmt.Println(err)
	}
	dataAddr := &chanData{resp,request,string(con)}
	process(dataAddr)
	resp.Write([]byte("hello world"))
}
func Api(resp http.ResponseWriter, request *http.Request) {
	resp.Write([]byte("hello world Api"))
}
func Report(resp http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id := query.Get("id")
	resp.Write([]byte("hello world Report"+id))
}
func Test(resp http.ResponseWriter, request *http.Request) {
	resp.Write([]byte("hello world Test"))
}
