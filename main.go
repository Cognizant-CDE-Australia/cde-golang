
package main

import (
	"cde-restservices/com/cde/route"
	_ "encoding/json"
	_ "github.com/gorilla/mux"
	_ "strconv"
)



func main() {
	route.RegisterRoutes()
}