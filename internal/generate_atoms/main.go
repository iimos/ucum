package main

import (
	"github.com/iimos/ucum/internal/xmlparser"
	"log"
	"net/http"
	"os"
	"time"
)

const url = "https://raw.githubusercontent.com/ucum-org/ucum/main/ucum-essence.xml"

func main() {
	resp, err := (&http.Client{Timeout: 10 * time.Second}).Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data := xmlparser.Parse(resp.Body)
	gen := Generator{
		packageName: "data",
	}
	gen.Generate(data)
	gocode := gen.Format()

	if err = os.WriteFile("./internal/data/atoms.gen.go", gocode, 0644); err != nil {
		log.Fatalf("writing output: %s", err)
	}
}
