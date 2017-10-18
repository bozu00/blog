package responses

import (
)

type Response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

func SafeResponse(err error, res interface{}) Response {
	return Response{
		200,
		"OK",
		res,
	}
}

func AuthFailResponse(err error, res interface{}) Response {
	return Response{
		401,
		"OK",
		res,
	}
}
