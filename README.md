# Address2Coordinates
This microservice sample reads ip2location database matches address to coordinates. Simply takes address and replies as that addresses coordinates.
This database includes Country-Providence-District
Before start the server file:
1-) Edit server and client JSON file with filling your own ip address and ports.
2-) If you want to use TLS connection update JSON file.

GET functions:
-> API Endpoint:
1-) /country : Lists all countries.
2-) /state/countryname : Lists all states in that country.
3-) /city/statename : Lists all cities in that state.
4-) /coordinates/cityname : Shows city coordinates.

-> Client App:
1-) country : Lists all countries
2-) state :  First, the province command is called and after entering, the country name is entered.
3-) city: First the district command is called and then the city name is entered.
4-) coordiantes: First, the coordinate command is called, entered and then the district name is entered.
