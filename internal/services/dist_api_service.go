package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"github.com/joho/godotenv"
)

type GeopifyResponse struct{
	Features []struct{
		Properties struct{
			Distance float64 `json:"distance"`
			Time float64  `json:"time"`
		} `json:"properties"`
	} `json:"features"`
}

type GeocodeResponse struct{
	Features []struct{
		Properties struct{
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"Properties"`
	}`json:"features"`
}

func GeoCode(city string) (float64,float64 ,error){
	
	err:=godotenv.Load("D:/Files/go_basic/hla_finder/.env")
	if err!=nil{
		return 0,0,err
	}
	apiKey:=os.Getenv("API_KEY")
	encoded:=url.QueryEscape(city)
	url:=fmt.Sprintf("https://api.geoapify.com/v1/geocode/search?text=%s&apiKey=%s", encoded, apiKey)

	resp,err:=http.Get(url)
	if err!=nil{
		return 0,0,err
	}
	defer resp.Body.Close()
	var data GeocodeResponse
	err=json.NewDecoder(resp.Body).Decode(&data)
	if err!=nil{
		return 0,0,err
	}
	if len(data.Features)==0{
		return 0,0,errors.New("no geocoding result found")
	}
	p:=data.Features[0].Properties
	return p.Lat,p.Lon,nil
}

func City_Distance(origin, destination string ) (int,error){
	//check inside cache
	dist,err:=GetDistance(origin,destination)
	//cache hit
	if dist!=-1||err==nil{
		return dist,err
	} 


	//if not found in cache
	err=godotenv.Load("D:/Files/go_basic/hla_finder/.env")
	if err!=nil{
		return -1,err
	}
	apiKey:=os.Getenv("API_KEY")
	o_lat,o_lon,err:=GeoCode(origin)
	if err!=nil{
		return 10000,err
	}
	d_lat,d_lon,err:=GeoCode(destination)
	if err!=nil{
		return 10000,err
	}


	url:=fmt.Sprintf(
		"https://api.geoapify.com/v1/routing?waypoints=%f,%f|%f,%f&mode=drive&apiKey=%s",
		o_lat,o_lon,d_lat,d_lon, apiKey,
	)

	resp,err:=http.Get(url)
	if err!=nil{
		return 100000,err
	}
	defer resp.Body.Close()
	var data GeopifyResponse
	err=json.NewDecoder(resp.Body).Decode(&data)
	if err!=nil{
		return 10000,err
	}
	if len(data.Features)==0{
		return 10000 ,errors.New("no path found")
	}

	distKm:=data.Features[0].Properties.Distance/1000
	//save inside cache
	SetDistance(origin,destination,int(distKm))

	return int(distKm),nil
}

