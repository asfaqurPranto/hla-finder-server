
package services

import (
	"fmt"
	"testing"
)

func TestCity_Distance(t *testing.T) {
	origin:="Chandpur"
	destination:="Dhaka"
	result, err := City_Distance(origin,destination)

	fmt.Print(origin," to ",destination," total distance: ")
	fmt.Println(result)
	fmt.Println(err)
}

func TestGeoCode(t *testing.T){
	origin:="Pabna"
	lat,lon,err:=GeoCode(origin)
	fmt.Println("City : ",origin)
	fmt.Print("Lat: ",lat," Lon: ",lon," Error:",err)
}

func TestSetDistance(t *testing.T){
	source:="pabna"
	destination:="rajshahi"
	err:=SetDistance(source,destination,109)
	fmt.Println(err)
}

func TestGetDistance(t *testing.T){
	source:="rajshahi"
	destination:="dhaka"
	dist,err:=GetDistance(source,destination)
	fmt.Println(source,destination,dist,err)
}