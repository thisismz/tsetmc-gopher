package tools

import (
	"tsetmc-gopher/global"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

func responserHandlepanic() {
	if a := recover(); a != nil {
		//fmt.Println("RECOVERED", a)
		global.BRC_LOG.Error("responserHandlepanic", zap.Any("err", a))
		//debug.PrintStack()
	}
}
func Responser(url string) string {
	defer responserHandlepanic()
	c := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := c.Get(url)
	defer resp.Body.Close()
	if err != nil {
		print(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	return string(body)
}
