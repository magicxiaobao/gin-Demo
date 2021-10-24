package entity

import (
	"encoding/json"
	"fmt"
)

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	ID      string      `json:"id,omitempty"`
}

type Resp interface {
	i() // 防止方法被其他包实现
	WithData(data interface{}) Resp
	WithId(id string) Resp
	ToString() string
}

func NewError(code int, msg string) Result {
	return Result{
		code, msg, nil, "",
	}
}

func (e Result) i() {}

func (e Result) WithData(data interface{}) Result {
	e.Data = data
	return e
}

func (e Result) WithId(id string) Result {
	e.ID = id
	return e
}

func (e Result) ToString() string {
	raw, error := json.Marshal(e)
	if error != nil {
		fmt.Println(error)
	}
	return string(raw)
}
