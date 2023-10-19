package main

import (
	adr2cor "server.go/utils"
	start "server.go/utils"
)

func main() {
	//This is address to coordinates microservice's server. This microservice helps to give you coordinates from an address by using csv database file.

	/*
		If you have an address (Sample: Turkey,Ankara,Cayyolu) you can use it like this;
		First, you can print all country names by using /country.
		Second, you can print all country's state by using /state/countryname (Sample: /state/Turkey)
		Third, you can print all state's cities by using /city/statename (Sample: /city/Ankara)
		Lastly if you want to get city's coordinates you should use /coordinates/cityname (Sample: /coordinates/Cayyolu)
	*/

	/*
		! Important !
		Before starting server.go file make sure server.json file has correct info.
	*/

	//You can change database file path in here.
	adr2cor.LoadData("doc/ip2location-lite-db5.csv")
	//Starting server
	start.Start()
	// Wait forever.
	select {}
}
