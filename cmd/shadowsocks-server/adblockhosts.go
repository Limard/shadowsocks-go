package main

import (
	"path/filepath"
	"os"
	"bufio"
	"bytes"
	"strings"
	"log"
)

var adBlockList []string

func init() {
	adBlockList = loadHosts()
}

func isBlockHost(host string) (bool, string) {
	for _, h := range adBlockList {
		if strings.Contains(host, h) {
			return true, h
		}
	}
	return false, ""
}

func loadHosts() (list []string) {
	filepath.Walk(`hosts`, func(fPath string, fInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fInfo.IsDir() {
			return nil
		}

		list = append(list, loadOneFile(fPath)...)
		return nil
	})
	log.Println("load adblock rules:", len(list))
	return
}

func loadOneFile(fPath string) (list []string) {
	f, e := os.Open(fPath)
	if e != nil {
		return
	}
	defer f.Close()

	comment := []byte("#")
	split := []byte(" ")
	fr := bufio.NewReader(f)
	for {
		line, _, e := fr.ReadLine()
		if e != nil {
			break
		}
		if bytes.HasPrefix(line, comment) {
			continue
		}
		kv := bytes.Split(line, split)
		if len(kv) >= 2 {
			list = append(list, string(kv[1]))
		}
	}
	return
}

var adHostCounter = make(map[string]int, 0)
var hostCounter = make(map[string]int, 0)

//func statisticsServer() {
//	http.HandleFunc(`/`, func(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte(fmt.Sprintf("<p>current ip address: %v</p>", r.RemoteAddr)))
//		w.Write([]byte("<p></p>"))
//		w.Write([]byte("<p>AD block:</p>"))
//		for key, value := range adHostCounter {
//			w.Write([]byte(fmt.Sprintf("<p>%v [%v]</p>", key, value)))
//		}
//		w.Write([]byte("<p></p>"))
//		w.Write([]byte("<p>URL:</p>"))
//		for key, value := range hostCounter {
//			w.Write([]byte(fmt.Sprintf("<p>%v [%v]</p>", key, value)))
//		}
//	})
//}