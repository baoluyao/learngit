package app

import (
	"fmt"
	"net/http"

	"noticeservice/global"
	"noticeservice/pkg/errcode"

	// "ffplat/internal/model"

	"github.com/carr123/fmx"
)

type Response struct {
	Ctx *fmx.Context
}
type Pager struct {
	Page      int `json:"page"`       //页码
	PageSize  int `json:"page_size"`  //每页容量
	TotalRows int `json:"total_rows"` //总行数
}
type ReData struct {
	Code int                    `json:"code"`
	Data map[string]interface{} `json:"data"`
	Msg  string                 `json:"msg"`
}
type ReDatas struct {
	Code int           `json:"code"`
	Data []interface{} `json:"data"`
	Msg  string        `json:"msg"`
}

func NewResponse(ctx *fmx.Context) *Response {
	header := ctx.Writer.Header()
	var nocache []string
	var expires []string
	nocache = append(nocache, "no-cache")
	header["Cache-Control"] = nocache
	expires = append(expires, "-1")
	header["Expires"] = expires
	header["Pragma"] = nocache

	return &Response{
		Ctx: ctx,
	}
}

// func (r *Response) ToResponse(data ...interface{}) {
// 	if len(data) == 0 {
// 		r.Ctx.JSON(http.StatusOK, fmx.H{"code": errcode.Success.Code(), "msg": errcode.Success.Msg()})
// 	} else if len(data) == 1 {
// 		r.Ctx.JSON(http.StatusOK, fmx.H{"code": errcode.Success.Code(), "msg": errcode.Success.Msg(), "data": data[0]})
// 	} else {
// 		r.Ctx.JSON(http.StatusOK, fmx.H{"code": errcode.Success.Code(), "msg": errcode.Success.Msg(), "data": data})
// 	}
// }

func (r *Response) ToResponse(data ...interface{}) {
	if len(data) == 0 {
		r.Ctx.JSON(http.StatusOK, fmx.H{"code": errcode.Success.Code(), "msg": errcode.Success.Msg()})
	} else if len(data) == 1 {
		r.Ctx.JSON(http.StatusOK, fmx.H{"code": errcode.Success.Code(), "msg": errcode.Success.Msg(), "data": data[0]})
	} else {
		r.Ctx.JSON(http.StatusOK, fmx.H{"code": errcode.Success.Code(), "msg": errcode.Success.Msg(), "data": data})
	}
}

func (r *Response) ErrorResponse(err error) {
	fmt.Println(err)
	r.Ctx.JSON(1, fmx.H{"data": err})

}
func (r *Response) ToSuccessResponse(data interface{}) {
	fmt.Println(data)
	r.Ctx.JSON(http.StatusOK, fmx.H{"data": data})

}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := fmx.H{"errorCode": err.Code(), "errorMsg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}

// func (r *Response) ToErrorResponse(err *errcode.Error) {
// 	response := fmx.H{"code": err.Code(), "msg": err.Msg()}
// 	details := err.Details()
// 	if len(details) > 0 {
// 		response["details"] = details
// 	}
// 	r.Ctx.JSON(err.StatusCode(), response)
// }
func (r *Response) ToErrorResponseExtra(err *errcode.Error, data interface{}) {
	response := fmx.H{"code": err.Code(), "msg": err.Msg(), "data": data}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}

func PublicNats(subj string, data []byte) *errcode.Error {
	err := global.Nats_client.Publish(subj, data)
	if err != nil {
		WriteErrorlog("subj:%s,param:%s,err:%v", subj, string(data), err)
		return errcode.RPCNetError
	}
	return nil
}

// func RequestNats(subj string, data []byte, timeout time.Duration) (*nats.Msg, *errcode.Error) {
// 	resp, err := global.Nats_client.Request(subj, data, time.Second*10)
// 	if err != nil {
// 		return nil, errcode.RPCNetError
// 	}
// 	return resp, nil
// }
