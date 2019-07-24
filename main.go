package main

import (
	"fmt"
	"math"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	calc(0.00001)
	fmt.Fprint(w, "Code.education Rocks!")
}

func calc(x float64) float64{
	for i := 0; i < 1000000; i++ {
		x += math.Sqrt(x)
	}
	return x
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server is up and listening on port 8080.")
	http.ListenAndServe(":8080", nil)
}