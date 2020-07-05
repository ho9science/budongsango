package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"time"
	"strings"
	"strconv"
)
type Response struct {
	XMLName xml.Name `xml:"response"`
	Header Header `xml:"header"`
	Body Body `xml:"body"`
}
type Header struct{
	ResultCode string `xml:"resultCode"`
	ResultMsg string `xml:"resultMsg"`
}
type Body struct{
	Items Items `xml:"items"`
}
type Items struct{
	Item []AptRealEstate `xml:"item"`
}

type Item struct {
	Item AptRealEstate
}
  
  type AptRealEstate struct {
	ResultCode string `xml:"일련번호"`
	RealYear string `xml:"년"`
	DealAmount string `xml:"거래금액"`
	BuildYear string `xml:"건축년도"`
	RoadName string `xml:"도로명"`
	RoadNameBonbun string `xml:"도로명본번호코드"`
	RoadNameBunbun string `xml:"도로명건물부번호코드"`
	RoadNameSigunguCode string `xml:"도로명시군구코드"`
	RoadNameSeq string `xml:"도로명일련번호코드"`
	RoadNameBasementCode string `xml:"도로명지상지하코드"`
	RoadNameCode string `xml:"도로명코드"`
	Dong string `xml:"법정동"`
	Bonbun string `xml:"법정동본번코드"`
	Bubun string `xml:"법정동부번코드"`
	SigunguCode string `xml:"법정동시군구코드"`
	EubmyundongCode string `xml:"법정동읍면동코드"`
	LandCode string `xml:"법정동지번코드"`
	ApartmentName string `xml:"아파트"`
	DealMonth int `xml:"월"`
	DealDay int `xml:"일"`
	AreaForExclusiveUse float64 `xml:"전용면적"`
	Jibun string `xml:"지번"`
	RegionalCode string `xml:"지역코드"`
	Floor int `xml:"층"`
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

  func startDate()(string){
	nowFile, err := ioutil.ReadFile("now")
	if err != nil {
		fmt.Println("couldn't open now")
	}
	return string(nowFile)
  }

  func saveLast(last string)(){
	tomorrow := []byte(last)
	err := ioutil.WriteFile("now", tomorrow, 0644)
	if err != nil {
	  fmt.Println("couldn't see tomorrow")
	}
  }

  func nextDate(targetDate string)(string){
	
	layout := "060102"
	
	t, err := time.Parse(layout, targetDate)

	if err != nil {
		fmt.Println(err)
	}

	nextDay := 1
	nextDate := t.AddDate(0, 0, +nextDay).Format(layout)

	return nextDate
  }
  
  func readCode()([][]string){
	csvfile, err := os.Open("code/refined_code.csv")
	if err != nil {
		fmt.Println("Couldn't open the csv file", err)
	}
	reader := csv.NewReader(csvfile)
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	return data
  }

  func getServiceKey()(string){
	  serviceKey, err := ioutil.ReadFile("servicekey")
	  if err != nil {
		fmt.Println("not found service key")
	  }
	  return string(serviceKey)
  }
  func padNumberWithZero(value int) string {
	return fmt.Sprintf("%02d", value)
  }

  func main(){
	  //connect Database
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
	if err != nil {
        panic(err.Error())
	}
	defer db.Close()
	sqlStr := "INSERT INTO BUDONGSAN.APT_REAL(DEALDATE, DEALAMOUNT, BUILDYEAR, ROADNAME, ROADNAMEBONBUN, ROADNAMEBUNBUN, ROADNAMECODE, DONG, JIBUN, APARTMENTNAME, AREAFOREXCLUSIVEUSE, REGIONCODE, REGIONNAME, FlOOR) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	count := 0
	var url = "http://openapi.molit.go.kr/OpenAPI_ToolInstallPackage/service/rest/RTMSOBJSvc/getRTMSDataSvcAptTradeDev?"
	var serviceKey = getServiceKey()
	codeList := readCode()
	var LAWD_CD = "11000"
	var DEAL_YMD = startDate() //200601
	
	for {
		for _, codeValue := range codeList{
			LAWD_CD = codeValue[0]
			url = url+"LAWD_CD="+LAWD_CD+"&DEAL_YMD="+DEAL_YMD +"&serviceKey="+serviceKey
			url += "&pageNo=1&numOfRows=1000"
			if xmlBytes, err := getXML(url); err != nil {
				log.Printf("Failed to get XML: %v", err)
			} else {
				var result Response
				err := xml.Unmarshal(xmlBytes, &result)
				if err != nil {
					log.Printf("error: %v", err)
				}
				var Items = result.Body.Items.Item
				for i := 0; i < len(Items); i++ {
					stmt, _ := db.Prepare(sqlStr)
					layout1 := "2006-01-02"
					layout2 := "2006"
					str := Items[i].RealYear+"-"+padNumberWithZero(Items[i].DealMonth)+"-"+padNumberWithZero(Items[i].DealDay)
					dealDate, err := time.Parse(layout1, str)
					if err != nil {
						log.Printf("error: %v", err)
					}
					buildYear, err := time.Parse(layout2, Items[i].BuildYear)
					if err != nil {
						log.Printf("error: %v", err)
					}		
					dealAmountStr := strings.TrimSpace(strings.Replace(Items[i].DealAmount, ",", "", -1))
					dealAmount, err := strconv.Atoi(dealAmountStr)
					
					_, err = stmt.Exec(dealDate, dealAmount, buildYear, Items[i].RoadName, Items[i].RoadNameBonbun, Items[i].RoadNameBunbun, Items[i].RoadNameCode, 
					Items[i].Dong, Items[i].Jibun, Items[i].ApartmentName, Items[i].AreaForExclusiveUse, Items[i].RegionalCode, codeValue[1], Items[i].Floor)
					if err != nil {
						log.Printf("error: %v", err)
					}
					fmt.Println("num:", i)
					defer stmt.Close()
				}
			}
		}
		DEAL_YMD = nextDate(DEAL_YMD)
		if  count >= 365 {
			saveLast(DEAL_YMD)
			break
		}
		count++
	}
}