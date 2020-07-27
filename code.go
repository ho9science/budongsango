package main

import (
	"os"
	"encoding/csv"
	"log"
	"net/http"
	"strings"
	"sort"
	"github.com/suapapa/go_hangul/encoding/cp949"
)

func readCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	r, _ := cp949.NewReader(resp.Body)
	reader := csv.NewReader(r)
	reader.Comma = '\t'
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func writeCSVFile(data map[string]string, keys []string) {
	file, err := os.Create("code/refined_code.csv")
	if err != nil {
        log.Fatal("can't create file", err)
    }
	defer file.Close()
	
	writer := csv.NewWriter(file)
	defer writer.Flush()
	objects := []string{}
	for _, k := range keys {
		err := writer.Write(append(objects, k, data[k]))
		if err != nil {
			log.Fatal("can't write to file")
		}
	}
	
}

func main() {
	var m = make(map[string]string)
	// Open the file
	url := "https://raw.githubusercontent.com/ho9science/budongsango/master/code/%EB%B2%95%EC%A0%95%EB%8F%99%EC%BD%94%EB%93%9C%20%EC%A0%84%EC%B2%B4%EC%9E%90%EB%A3%8C.txt"
	data, err := readCSVFromUrl(url)
	if err != nil {
		panic(err)
	}

	for idx, row := range data {
		// skip header
		if idx == 0 {
			continue
		}

		lawd_cd := row[0][0:5]
		location := strings.Split(row[1], " ")
		disable := row[2]
		if disable == "폐지" {
			continue
		}
		if(len(location)>1){
			m[lawd_cd] = location[0]+" "+location[1]
			if location[0] == "세종특별자치시"{
				m[lawd_cd] = location[0]
			}
			if(len(location)>2){
				if (len(location[2])>=9 && strings.LastIndex(location[2], "구") > 3) || (len(location[2])==6 && strings.LastIndex(location[2], "구")== 3){
					m[lawd_cd] = location[0]+" "+location[1] + " " + location[2]
				}
			}			
		}
		
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	writeCSVFile(m, keys)

}