package main

import (
	errors "github.com/mengxiaozhu/linkerror"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

var JSONError = errors.Type("JSON Error")
var HTTPError = errors.Type("HTTP Error")

func main() {
	resp, err := webTranslate("love me love my dog")
	if err != nil {
		if err.Catch(JSONError) {
			fmt.Println("捕获到JSON异常")
			fmt.Println(err)
		} else if err.Catch(HTTPError) {
			fmt.Println("捕获到HTTP异常")
			fmt.Println(err)
			return
		} else {
			fmt.Println("程序挂了")
			return
		}
	}

	fmt.Printf("%+v",resp)
}

//http://fanyi.youdao.com/openapi.do?keyfrom=<keyfrom>&key=<key>&type=data&doctype=<doctype>&version=1.1&q=要翻译的文本

func webTranslate(word string) (*YoudaoAPIResp, *errors.Error) {

	resp, err := http.Get("http://fanyi.youdao.com/openapi.do?" + url.Values{
		"keyfrom": {"mamashipu"},
		"key":     {"1350041455"},
		"type":    {"data"},
		"doctype": {"json"},
		"version": {"1.1"},
		"q":       {word},
	}.Encode())
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, errors.NewWith(HTTPError, "发送请求时出现异常", err)
	}
	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.NewWith(HTTPError, "读取数据时发生异常", err)
	}
	youdaoAPIResp := &YoudaoAPIResp{}
	err = json.Unmarshal(jsonBytes, youdaoAPIResp)
	if err != nil {
		return nil, errors.NewWith(JSONError, "解析有道翻译返回json时发生异常", err)
	}
	return youdaoAPIResp, nil
}

type YoudaoAPIResp struct {
	Translation []string `db:"translation" json:"translation"`
	Basic struct {
		UsPhonetic string `db:"us-phonetic" json:"us-phonetic"`
		Phonetic   string `db:"phonetic" json:"phonetic"`
		UkPhonetic string `db:"uk-phonetic" json:"uk-phonetic"`
		Explains   []string `db:"explains" json:"explains"`
	} `db:"basic" json:"basic"`
	Query     string `db:"query" json:"query"`
	ErrorCode int `db:"error_code" json:"errorCode"`
	Web []struct {
		Value []string `db:"value" json:"value"`
		Key   string `db:"key" json:"key"`
	} `db:"web" json:"web"`
}

