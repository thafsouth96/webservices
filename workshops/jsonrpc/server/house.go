package main

import (
	"log"
	"net"
	"net/http"
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

func (t *House) GetHouse(args Args, result *HouseDAO) error {
	log.Printf("ID : %d", args.Id)
	*result = houses[args.Id-1]
	log.Print(result)
	return nil
}

func (t *House) GetHouses(args Args, result *[]HouseDAO) error {
	*result = houses
	return nil
}

func main() {
	house := new(House)
	listener, e := net.Listen("tcp", "0.0.0.0:8080")
	if e != nil {
		log.Fatal("Listen error : ", e)
	}
	log.Printf("Starting server on port 8080")

	server := rpc.NewServer()
	server.Register(house)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		server.ServeHTTP(w, r)
	})
	err := http.Serve(listener, nil)

	if err != nil {
		log.Fatal("Error serving : ", err)

	}

}
