package middleware

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
	structs "server.go/utils"
)

// This function sends JSON response with error message and status code.
func SendErrorResponse(ctx *fasthttp.RequestCtx, errorMsg string, statusCode int) {
	resp := structs.Response{
		Body:   nil,
		Error:  errorMsg,
		Status: statusCode,
	}

	response, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error encoding JSON error response %d\n", fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)

	if _, err := ctx.Write(response); err != nil {
		log.Fatalf("Error writing response %v\n", err)
	}
}

// This function sends JSON response with data, error message, and status code.
func SendJSONResponse(ctx *fasthttp.RequestCtx, data interface{}, errorMsg string, statusCode int) {
	resp := structs.Response{
		Body:   data,
		Error:  errorMsg,
		Status: statusCode,
	}

	response, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error encoding JSON response %d\n", fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)

	if _, err := ctx.Write(response); err != nil {
		log.Fatalf("Error writing response %v\n", err)
	}
}
