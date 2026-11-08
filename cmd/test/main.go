package main

import (
	"fmt"

	"github.com/maiconDeSouza/lotofacil_lab/pkg/random"
)

func main() {
	v := random.GenerateLotofacilNumbers15()
	fmt.Println(v, len(v))
}
