package responses

import (
)

type Response struct {
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

func SafeResponse(err error, res interface{}) Response {
	return Response{
		"OK",
		res,
	}
}

func AuthFailResponse(err error, res interface{}) Response {
	return Response{
		"OK",
		res,
	}
}
