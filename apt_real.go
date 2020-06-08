package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)
type Response struct {
	XMLName xml.Name `xml:"response"`
	Header Header `xml: "header"`
	Body Body `xml: "body"`
}
type Header struct{
	ResultCode string `xml: "resultCode"`
	ResultMsg string `xml: "resultMsg"`
}
type Body struct{
	Items Items `xml: "items"`
}
type Items struct{
	Item []AptRealEstate `xml:"item"`
}

type Item struct {
	Item AptRealEstate
}
  
  type AptRealEstate struct {
	ResultCode string `xml: "일련번호"`
	RealYear string `xml: "년"`
	DealAmount string `xml: "거래금액"`
	BuildYear string `xml: "건축년도"`
	RoadName string `xml: "도로명"`
	RoadNameBonbun string `xml: "도로명본번호코드"`
	RoadNameBunbun string `xml: "도로명건물부번호코드"`
	RoadNameSigunguCode string `xml: "도로명시군구코드"`
	RoadNameSeq string `xml: "도로명일련번호코드"`
	RoadNameBasementCode string `xml: "도로명지상지하코드"`
	RoadNameCode string `xml: "도로명코드"`
	Dong string `xml: "법정동"`
	Bonbun string `xml: "법정동본번코드"`
	Bubun string `xml: "법정동부번코드"`
	SigunguCode string `xml: "법정동시군구코드"`
	EubmyundongCode string `xml: "법정동읍면동코드"`
	LandCode string `xml: "법정동지번코드"`
	ApartmentName string `xml: "아파트"`
	DealMonth string `xml: "월"`
	DealDay string `xml: "일"`
	AreaForExclusiveUse float64 `xml: "전용면적"`
	Jibun string `xml: "지번"`
	RegionalCode string `xml: "지역코드"`
	Floor int `xml: "층"`
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
	var serviceKey = "6n9eEV402Wxn5iBELz6R89uQRbB0tr5S2k23uKx%2F853VLXwdGY%2FlK%2B5Oo8W56xvG%2BtNe0NLj536ruGWjzkIAEQ%3D%3D"	
	var LAWD_CD = "11110"
  	var DEAL_YMD = "201512"
	url = url+"LAWD_CD="+LAWD_CD+"&DEAL_YMD="+DEAL_YMD +"&serviceKey="+serviceKey
	log.Printf(url)
	if xmlBytes, err := getXML(url); err != nil {
		log.Printf("Failed to get XML: %v", err)
	  } else {
		var result Response
		err := xml.Unmarshal(xmlBytes, &result)
		if err != nil {
			log.Printf("error: %v", err)
		}
		for i := 0; i < len(result.Body.Items.Item); i++ {
			log.Printf(result.Body.Items.Item[i].ResultCode)
		}
	  }
	}