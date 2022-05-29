package models

import (
	"aulia-majoo-test/pkg/c"
)

func SuccessResponses(data interface{}) Responses {
	return Responses{
		Status: true,
		Data:   data,
		Code:   c.SucceedCode,
	}
}

func NotFoundResponses(status interface{}) Responses {
	return Responses{
		Status: false,
		Data:   nil,
		Code:   c.NotFoundCode,
	}
}

func BadRequestResponses(data string) Responses {
	return Responses{
		Status: false,
		Data:   nil,
		Code:   c.BadRequestCode,
	}
}

func InternalServerResponses(data string) Responses {
	return Responses{
		Status: false,
		Data:   nil,
		Code:   c.InternalServerCode,
	}
}

func SuccessPaginationResponses(data interface{}, meta interface{}) (result JSONResponse) {
	return JSONResponse{
		Status: "Succeed",
		Code:   c.SucceedCode,
		Data:   data,
		Meta:   meta,
	}
}

func InternalServerPaginationResponses(data string) (result JSONResponse) {
	return JSONResponse{
		Status: data,
		Code:   c.InternalServerCode,
		Data:   nil,
		Meta:   Meta{},
	}
}

func NotFoundPaginationResponses(status string, data interface{}) (result JSONResponse) {
	return JSONResponse{
		Status: data,
		Code:   c.NotFoundCode,
		Data:   data,
		Meta:   Meta{},
	}
}

func BadRequestPaginationResponses(data string) (result JSONResponse) {
	return JSONResponse{
		Status: data,
		Code:   c.BadRequestCode,
		Data:   nil,
		Meta:   Meta{},
	}
}

func Success(data interface{}) (result Responses) {
	return Responses{
		Status: "Success",
		Data:   data,
		Code:   c.SucceedCode,
	}
}

func NotFound(status interface{}) (result Responses) {
	return Responses{
		Status: status,
		Data:   nil,
		Code:   c.NotFoundCode,
	}
}

func BadRequest(data string) (result Responses) {
	return Responses{
		Status: data,
		Data:   nil,
		Code:   c.BadRequestCode,
	}
}

func InternalServer(data string) (result Responses) {
	return Responses{
		Status: data,
		Data:   nil,
		Code:   c.InternalServerCode,
	}
}

func SuccessWPagination(data interface{}, meta interface{}) (result JSONResponse) {
	return JSONResponse{
		Status: "Succeed",
		Code:   c.SucceedCode,
		Data:   data,
		Meta:   meta,
	}
}

func InternalServerWPagination(data string) (result JSONResponse) {
	return JSONResponse{
		Status: data,
		Code:   c.InternalServerCode,
		Data:   nil,
		Meta:   Meta{},
	}
}

// NotFoundExceptions : Http Status Handle for 404
func NotFoundExceptions() Responses {
	return Responses{
		Code:   c.NotFoundCode,
		Status: "Data not found!",
	}
}

// BadRequestException : Http Status Handle for 400
func BadRequestException(message interface{}) Responses {
	return Responses{
		Code:   c.BadRequestCode,
		Status: message,
	}
}

func AuthException(message interface{}, code int) Responses {
	return Responses{
		Code:   code,
		Status: message,
	}
}

func ForbiddenException() Responses {
	return Responses{
		Code:   c.ForbiddenCode,
		Status: "Access Forbidden",
	}
}

func ForbiddenExceptionPagination() JSONResponse {
	return JSONResponse{
		Code:   c.ForbiddenCode,
		Status: "Access Forbidden",
		Meta:   nil,
	}
}
