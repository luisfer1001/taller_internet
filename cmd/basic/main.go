package main

import (
	"log"

	"github.com/test/pkg/tools"
)

func main() {
	rs, rr := tools.SumaResta(10, 5)
	rd, err := tools.Division(10, 0)

	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("This suma result is:", rs)
	log.Println("This resta result is:", rr)
	log.Println("This division result is:", rd)
}
