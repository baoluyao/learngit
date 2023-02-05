package model

type Msgdata struct {
	Code    int         `json:"code,omitempty"`
	ID      string      `json:"id"`
	SN      string      `json:"sn"`
	Message string      `json:"message"`
	Module  string      `json:"module"`
	Data    interface{} `json:"data,omitempty"`
	TS      int64       `json:"ts"`
}
