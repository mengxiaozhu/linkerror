package main

import (
	errors "github.com/mengxiaozhu/linkerror"
	"net/http"
	"net/url"
)

var JSONError = errors.Type("JSON Error")
var HTTPError = errors.Type("HTTP Error")

func main() {

}

//http://fanyi.youdao.com/openapi.do?keyfrom=<keyfrom>&key=<key>&type=data&doctype=<doctype>&version=1.1&q=要翻译的文本

func fetch(word string) {

	resp,err:=http.Get("fanyi.youdao.com/openapi.do?" + url.Values{
		"keyfrom": {"mamashipu"},
		"key":     {"1350041455"},
		"type":    {"data"},
		"doctype": {"json"},
		"version": {"1.1"},
		"q":       {word},
	}.Encode())
	if resp!=nil && resp.Body!=nil{
		defer resp.Body.Close()
	}
	if err!=nil{
		errors.NewWith(HTTPError,"发送请求时出现异常",)
	}
}
