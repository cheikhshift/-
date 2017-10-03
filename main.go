package main 

import (
	"encoding/json"
	"net/http"
	"log"
	"fmt"
)

func main() {
	
	var dat map[string]string

	r, err := http.Get("https://talaikis.com/api/quotes/random/")
    if err != nil {
        log.Println(err)
        return
    }
    defer r.Body.Close()
    if err := json.NewDecoder(r.Body).Decode(&dat); err != nil {
        log.Println(err)
        return
    }
    fmt.Println("🌈 ", dat["quote"], "\nBy ",dat["author"],",",dat["cat"] )
}