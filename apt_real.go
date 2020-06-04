package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)
  
  type AptRealEstate struct {
	ResultCode string
	ResultMsg string
	DealAmount string
	BuildYear string 
	RoadName string
	RoadNameBonbun string
	RoadNameBunbun string
	RoadNameSigunguCode string
	RoadNameSeq string
	RoadNameBasementCode string
	RoadNameCode string
	Dong string
	Bonbun string
	Bubun string
	SigunguCode string
	EubmyundongCode string
	LandCode string
	ApartmentName string
	DealMonth string
	DealDay string
	AreaForExclusiveUse float64
	Jibun string
	RegionalCode string
	Floor int
  }
  
  func getXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
	  return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()
  
	if resp.StatusCode != http.StatusOK {
	  return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
	}
  
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	  return []byte{}, fmt.Errorf("Read body: %v", err)
	}
  
	return data, nil
  }

  func main(){

	var url = "http://openapi.molit.go.kr/OpenAPI_ToolInstallPackage/service/rest/RTMSOBJSvc/getRTMSDataSvcAptTradeDev?"
	var serviceKey = ""	
	var LAWD_CD = "11110"
  	var DEAL_YMD = "201512"
	url = url+"LAWD_CD="+LAWD_CD+"&DEAL_YMD="+DEAL_YMD +"&serviceKey="+serviceKey
	log.Printf(url)
	if xmlBytes, err := getXML(url); err != nil {
		log.Printf("Failed to get XML: %v", err)
	  } else {
		  log.Printf(string(xmlBytes))
		var result AptRealEstate
		xml.Unmarshal(xmlBytes, &result)
	  }
	}