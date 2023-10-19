package utils

import (
	"encoding/csv"
	"github.com/valyala/fasthttp"
	"io"
	"log"
	"os"
	jsonresponce "server.go/middleware"
	"sync"
)

var (
	countries        map[string]struct{}
	states           map[string]map[string]struct{}
	cities           map[string]map[string]struct{}
	coordinates      map[string][]string
	countriesMutex   sync.RWMutex
	statesMutex      sync.RWMutex
	citiesMutex      sync.RWMutex
	coordinatesMutex sync.RWMutex
	logFile          = SetupLogging()
)

// This fuction reads csv file and writes informations in to 4 maps.
func LoadData(csvFile string) {
	log.SetOutput(logFile)
	file, err := os.Open(csvFile)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	reader := csv.NewReader(file)

	//Making mappings
	countries = make(map[string]struct{})
	states = make(map[string]map[string]struct{})
	cities = make(map[string]map[string]struct{})
	coordinates = make(map[string][]string)

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Error reading CSV record %v", err)
			break
		}

		if record[3] != "-" && record[3] != "" && record[4] != "-" && record[4] != "" && record[5] != "-" && record[5] != "" {
			countries[record[3]] = struct{}{}

			if _, exists := states[record[3]]; !exists {
				states[record[3]] = make(map[string]struct{})
			}
			states[record[3]][record[4]] = struct{}{}
			if _, exists := cities[record[4]]; !exists {
				cities[record[4]] = make(map[string]struct{})
			}

			cities[record[4]][record[5]] = struct{}{}
			if _, exists := coordinates[record[5]]; !exists {
				coordinates[record[5]] = []string{record[6], record[7]}
			}
		}
	}
}

// This function sends country names.
func GetCountryNames(ctx *fasthttp.RequestCtx) {
	countriesMutex.RLock()
	jsonresponce.SendJSONResponse(ctx, getKeysFromMap(countries), "No error", fasthttp.StatusOK)
	countriesMutex.RUnlock()
}

// This function sends chosen country's states names.
func GetStatesByCountry(ctx *fasthttp.RequestCtx, country string) {
	log.SetOutput(logFile)
	statesInCountry := getStatesWithinCountry(country)
	if len(statesInCountry) > 0 {
		statesMutex.RLock()
		jsonresponce.SendJSONResponse(ctx, statesInCountry, "No error", fasthttp.StatusOK)
		statesMutex.RUnlock()
	} else {
		log.Printf("Country not found or no states available  %d\n", fasthttp.StatusNotFound)
		jsonresponce.SendErrorResponse(ctx, "Country not found or no states available", fasthttp.StatusNotFound)
	}
}

// This function returns states in specific country in slice, I had to do this because I don't want to repeating states.
func getStatesWithinCountry(countryName string) []string {
	log.SetOutput(logFile)
	var statesInCountry []string
	if statesMap, found := states[countryName]; found {
		for city := range statesMap {
			statesInCountry = append(statesInCountry, city)
		}
	} else {
		log.Println("Error in getStatesWithinCountry")
	}

	return statesInCountry
}

// This function sends chosen state's cities names.
func GetCitiesByState(ctx *fasthttp.RequestCtx, state string) {
	log.SetOutput(logFile)
	citiesInState := getCitiesWithinState(state)
	if len(citiesInState) > 0 {
		citiesMutex.RLock()
		jsonresponce.SendJSONResponse(ctx, citiesInState, "No error", fasthttp.StatusOK)
		citiesMutex.RUnlock()
	} else {
		log.Printf("State not found or no cities available %d\n", fasthttp.StatusNotFound)
		jsonresponce.SendErrorResponse(ctx, "State not found or no cities available", fasthttp.StatusNotFound)
	}
}

// This function returns cities in specific state in slice, I had to do this because I don't want to repeating cities.
func getCitiesWithinState(stateName string) []string {
	log.SetOutput(logFile)
	var citiesInState []string
	if citiesMap, found := cities[stateName]; found {
		for city := range citiesMap {
			citiesInState = append(citiesInState, city)
		}
	} else {
		log.Println("Error in getCitiesWithinState")
	}

	return citiesInState
}

// This function sends chosen city's coordinates.
func GetCoordinatesByCity(ctx *fasthttp.RequestCtx, city string) {
	log.SetOutput(logFile)
	coordinates, found := coordinates[city]
	if found {
		coordinatesMutex.RLock()
		jsonresponce.SendJSONResponse(ctx, coordinates, "No error", fasthttp.StatusOK)
		coordinatesMutex.RUnlock()
	} else {
		log.Fatalf("City not found %d\n", fasthttp.StatusNotFound)
		jsonresponce.SendErrorResponse(ctx, "City not found", fasthttp.StatusNotFound)
	}
}

// This function returns specific map's key form as slice.
func getKeysFromMap(inputMap map[string]struct{}) []string {
	keys := make([]string, 0, len(inputMap))
	for key := range inputMap {
		keys = append(keys, key)
	}
	return keys
}
