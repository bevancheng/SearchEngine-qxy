package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type PicurlText struct {
	ID       uint32   `json:"id"`
	Text     string   `json:"text"`
	Document Document `json:"document"`
}

type Document struct {
	Url     string `json:"url"`
	RawText string `json:"rawtext"`
}

func main() {
	filename := "./wukong50k_release.csv"
	ReadCsv(filename)

}

func ReadCsv(filepath string) {
	opencast, err := os.Open(filepath)
	if err != nil {
		log.Println("csv open failed!")
	}
	defer opencast.Close()

	rdcsv := csv.NewReader(opencast)
	csvhead, _ := rdcsv.Read()
	log.Println(csvhead)

	//send
	client := &http.Client{}
	var cnt uint32 = 1
	for line, err1 := rdcsv.Read(); err1 == nil; line, err1 = rdcsv.Read() {
		doc := Document{Url: line[0], RawText: line[1]}
		data := PicurlText{ID: cnt, Text: line[1], Document: doc}
		cnt++
		jsonStu, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		var postbytes = bytes.NewReader(jsonStu)
		req, err := http.NewRequest("POST", "http://127.0.0.1:5678/api/index", postbytes)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		bodyText, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		fmt.Printf("%s\n", bodyText)
	}

}
