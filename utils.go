package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

func CheckNormal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckFatal(err error) {
	if err != nil {
		panic(err)
	}
}

func MarshalJson(p []Post) []byte {
	data, err := json.MarshalIndent(p, "", "     ")
	CheckNormal(err)
	return data
}

func UnMarshalJson(d []byte) []Post {
	var ps []Post
	err := json.Unmarshal(d, &ps)
	CheckNormal(err)
	return ps
}

func ReadFromFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	CheckFatal(err)
	return data
}

func WriteToFile(v []Post, path string) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0775)
	CheckFatal(err)
	defer f.Close()
	data := MarshalJson(v)
	//fmt.Println(string(data))
	_, err = f.Write(data)
	CheckFatal(err)
}

func StartTimer(work func()) {
	go func() {
		for {
			work()
			now := time.Now()
			next := now.Add(time.Minute*1)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t := time.NewTicker(next.Sub(now))
			<- t.C
		}
	}()
}

func GenerateFileName(filepath string, tag string) string {
	return filepath + tag + strconv.FormatInt(time.Now().Unix(), 10) + ".json"
}