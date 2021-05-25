package main

import (
	"encoding/json"
	"log"

	"github.com/swill/kad"
)

func main() {
	// you can define settings and the layout in JSON
	json_bytes := []byte(`{
		"switch-type":3,
		"stab-type":1,
		"layout":[
			["Num Lock","/","*","-"],
			[{"f":3},"7\nHome","8\n↑","9\nPgUp",{"h":2}," "],
			["4\n←","5","6\n→"],["1\nEnd","2\n↓","3\nPgDn",{"h":2},"Enter"],
			[{"w":2},"0\nIns",".\nDel"]
		],
		"case": {
			"case-type":"sandwich",
			"mount-holes-num":4,
			"mount-holes-size":3,
			"mount-holes-edge":6
		},
		"top-padding":9,
		"left-padding":9,
		"right-padding":9,
		"bottom-padding":9
	}`)

	// create a new KAD instance
	cad := kad.New()

	// populate the 'cad' instance with the JSON contents
	err := json.Unmarshal(json_bytes, cad)
	if err != nil {
		log.Fatalf("Failed to parse json data into the KAD file\nError: %s", err.Error())
	}

	// and you can define settings via the KAD instance
	cad.Hash = "usage_example"      // the name of the design
	cad.FileStore = kad.STORE_LOCAL // store the files locally
	cad.FileDirectory = "./"        // the path location where the files will be saved
	cad.FileServePath = "/"         // the url path for the 'results' (don't worry about this)

	// here are some more settings defined for this case
	cad.Case.UsbWidth = 12 // all dimension are in 'mm'
	cad.Fillet = 3         // 3mm radius on the rounded corners of the case

	cad.Result.Formats = append(cad.Result.Formats, []string{"eps"}...)

	// lets draw the SVG files now
	err = cad.Draw()
	if err != nil {
		log.Fatal("Failed to Draw the KAD file\nError: %s", err.Error())
	}
}

