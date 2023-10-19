package middleware

import (
	"github.com/valyala/fasthttp"
	adr2cor "server.go/utils"
	"strings"
)

func RequestHandler(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())

	switch {
	// ipaddress:port/country ile bütün ülkeler çağrılır.
	case path == "/country":
		adr2cor.GetCountryNames(ctx)
	// ipaddress:port/state/countryname ile seçilen ülkedeki bütün iller çağrılır.
	case strings.HasPrefix(path, "/state/"):
		country := strings.TrimPrefix(path, "/state/")
		adr2cor.GetStatesByCountry(ctx, country)
	// ipaddress:port/city/statename ile seçilen ildeki bütün ilçeler çağrılır.
	case strings.HasPrefix(path, "/city/"):
		state := strings.TrimPrefix(path, "/city/")
		adr2cor.GetCitiesByState(ctx, state)
	// ipaddress:port/coordinates/cityname ile seçilen ilçedeki koordinatlar verilir.
	case strings.HasPrefix(path, "/coordinates/"):
		city := strings.TrimPrefix(path, "/coordinates/")
		adr2cor.GetCoordinatesByCity(ctx, city)
	default:
		ctx.Error("Not Found", fasthttp.StatusNotFound)
	}
}
