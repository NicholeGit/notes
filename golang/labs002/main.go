package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"rest-master/rest"
)

var (
	RuleName = &rest.StringRule{"name", 4, 20}
	RuleAge  = &rest.IntRule{"age", 10, 60}
)

func authorize(pattern string, request *http.Request, params rest.Value) bool {
	return params.String("name") == "user1"
}

func test(params rest.Value) rest.Value {
	fmt.Println("in:", params)

	// 使用规则验证输入参数。
	if !rest.RuleCheck(params, RuleName, RuleAge) {
		return rest.NewValue().Error("invalid parameters")
	}

	// 返回数据。
	params["__time__"] = time.Now()
	return params
}

func main() {

	s := rest.NewServer("0.0.0.0:8080")

	s.StaticDir = "."       // 静态文件目录。
	s.Pattern = "/v2/"      // REST 路径前缀匹配模式。
	s.Authorize = authorize // 请求验证函数。

	s.HandleFunc("!/test/{name}/{age}", rest.GET, test) // http://localhost:8080/v2/test/user1/23

	fmt.Printf("%#v", s)
	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
