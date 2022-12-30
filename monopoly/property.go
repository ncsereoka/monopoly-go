package monopoly

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Property struct {
	owner        int
	name         string
	price        int
	houses       uint8
	hotel        bool
	rentSite     int
	rent1house   int
	rent2house   int
	rent3house   int
	rent4house   int
	rentHotel    int
	buildingCost int
}

var PropertyMap = map[int]*Property{}

func InitProperties() {
	content, err := ioutil.ReadFile("./monopoly/property.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	err = json.Unmarshal(content, &PropertyMap)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
}
