package utils

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	handler "server.go/middleware"
)

func Start() {
	// Read the configuration from the JSON file
	configFile, err := os.Open("permanent/server.json")
	if err != nil {
		log.Fatalf("Error opening configuration file %v\n", err)

	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&Config); err != nil {
		log.Fatalf("Error reading configuration from file: %v", err)
	}
	var addr = Config.IP + ":" + Config.Port

	server := &fasthttp.Server{
		Handler: handler.RequestHandler,
	}

	//Running Server. if there is no info about certificate than start it in http server.
	fmt.Printf("Server is running on port %s\n", addr)
	if Config.Certificate == "" && Config.Key == "" {
		err = server.ListenAndServe(addr)
		if err != nil {
			log.Println("Can not open HTTP server.")
		}
	} else {
		err = server.ListenAndServeTLS(addr, Config.Certificate, Config.Key)
		if err != nil {
			log.Println("Can not open HTTPS server.")
		}
	}

}
