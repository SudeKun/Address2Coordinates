package utils

var Config struct {
	IP          string `json:"IP"`
	Port        string `json:"Port"`
	Certificate string `json:"certFile"`
	Key         string `json:"keyFile"`
}

type Response struct {
	Body   interface{} `json:"body"`
	Error  string      `json:"error"`
	Status int         `json:"status"`
}
