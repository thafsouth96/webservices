package main

import (
	"log"
	"net/rpc"
)

type HouseDAO struct {
	Id         int
	Name       string
	Region     string
	CoatOfArms string
	Words      string
}

var houses = []HouseDAO{
	HouseDAO{
		Id:         1,
		Name:       "House Algood",
		Region:     "The Westerlands",
		CoatOfArms: "A golden wreath, on a blue field with a gold border(Azure, a garland of laurel within a bordure or)",
		Words:      "",
	},
	HouseDAO{
		Id:         2,
		Name:       "House Allyrion of Godsgrace",
		Region:     "Dorne",
		CoatOfArms: "Gyronny Gules and Sable, a hand couped Or",
		Words:      "No Foe May Pass",
	},
	HouseDAO{
		Id:         3,
		Name:       "House Amber",
		Region:     "The North",
		CoatOfArms: "",
		Words:      "",
	},
}

type Args struct {
	Id int
}

type House int

func main() {
	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error in dialing. %d", err)
	}
	args := &Args{
		Id: 2,
	}
	var result HouseDAO
	log.Printf("Call of function House.GetHouse with the id : %d ", args.Id)
	err = client.Call("House.GetHouse", args, &result)
	if err != nil {
		log.Fatalf("error in House : %d", err)
	}
	log.Printf("Id : %d", result.Id)
	log.Print("Name : " + result.Name)
	log.Print("Region : " + result.Region)
	log.Print("CoatOfArms : " + result.CoatOfArms)
	log.Print("Words : " + result.Words)

	log.Println("Call of function House.GetHouses")
	var resultHouses []HouseDAO
	args = &Args{}
	err = client.Call("House.GetHouses", args, &resultHouses)
	if err != nil {
		log.Fatalf("error in House : %d", err)
	}
	for i := range resultHouses {
		log.Printf("House n°%d", i+1)
		log.Print("Name : " + resultHouses[i].Name)
		log.Print("Region : " + resultHouses[i].Region)
		log.Print("CoatOfArms : " + resultHouses[i].CoatOfArms)
		log.Print("Words : " + resultHouses[i].Words)
	}
}
