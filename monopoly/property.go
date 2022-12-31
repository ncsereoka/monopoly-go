package monopoly

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Property struct {
	owner        int
	Name         string
	Price        int
	houses       uint8
	hotel        bool
	RentSite     int
	Rent1house   int
	Rent2house   int
	Rent3house   int
	Rent4house   int
	RentHotel    int
	BuildingCost int
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
