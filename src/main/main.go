package main

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"net"
	"path/filepath"
)

func main() {
	path, err := filepath.Abs(filepath.Dir("."))
	fmt.Println(path)
	if err == nil {
		dbFilePath := path + string(filepath.Separator) + "src" + string(filepath.Separator) + "GeoIP2-City.mmdb"
		fmt.Println(dbFilePath)
		db, err := geoip2.Open(dbFilePath)
		defer db.Close()

		ip := net.ParseIP("124.127.205.82")
		record, err := db.City(ip)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(record.City.Names["cn"])
		fmt.Println(record.City.Names)
		fmt.Println(record.City.Names["zh-CN"])
		//fmt.Printf("Portuguese (BR) city name: %v\n", record.City.Names["pt-BR"])
		//fmt.Printf("English subdivision name: %v\n", record.Subdivisions[0].Names["en"])
		//fmt.Printf("Russian country name: %v\n", record.Country.Names["ru"])
		//fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
		//fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
		//fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude
	}

}
