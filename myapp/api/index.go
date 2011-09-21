package index

import (
	"http"
	"fmt"
	"app"
)

func init() {
	http.HandleFunc("/", main)
}

func main(w http.ResponseWriter, r *http.Request) {
	
}
