package task

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"

	simplejson "github.com/bitly/go-simplejson"
)

type TaskInterface interface {
	Run()
}

type Task struct {
	id int
}

func (task *Task) Run() {

}

type RequestTask Task

func NewRequestTask() *RequestTask {
	reqTask := new(RequestTask)
	reqTask.id = rand.Int()
	return reqTask
}

func (task *RequestTask) Run() {
	resp, err := http.Get("http://api.map.baidu.com/telematics/v3/weather?location=%E5%98%89%E5%85%B4&output=json&ak=5slgyqGDENN7Sy7pw29IUvrZ")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}
	js, err := simplejson.NewJson(body)
	fmt.Println(js)
}
