package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/swill/kad"
)

func main() {
	if len(os.Args[1:]) < 1 {
		log.Fatalf("usage: %s layout.json\n", filepath.Base(os.Args[0]))
	}
	// inFile := "layout-01.json"
	inFile := os.Args[1]
	drawingName := strings.TrimSuffix(inFile, ".json")

	// load settings and the layout from JSON file
	data, err := ioutil.ReadFile(inFile)
	if err != nil {
		log.Fatalf("Can't read %s\n", inFile)
	}

	// create a new KAD instance
	cad := kad.New()

	// populate the 'cad' instance with the JSON contents
	err = json.Unmarshal(data, cad)
	if err != nil {
		log.Fatalf("Failed to parse json data into the KAD file\nError: %s", err.Error())
	}

	// and you can define settings via the KAD instance
	cad.Hash = drawingName          // the name of the design
	cad.FileStore = kad.STORE_LOCAL // store the files locally
	cad.FileDirectory = "./"        // the path location where the files will be saved
	cad.FileServePath = "/"         // the url path for the 'results' (don't worry about this)

	cad.Result.Formats = append(cad.Result.Formats, []string{"eps"}...)
	fmt.Println("Formats to produce:")
	for _, ft := range cad.Result.Formats {
		fmt.Printf("  %s\n", ft)
	}

	// lets draw the SVG and EPS files now
	err = cad.Draw()
	if err != nil {
		log.Fatalf("Failed to Draw the KAD file\nError: %s", err.Error())
	}
}
