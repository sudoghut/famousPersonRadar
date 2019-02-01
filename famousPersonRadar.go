package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// read the person list to query
func readPersonList(fileName string) []string {
	var output []string
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		output = append(output, line)
	}
	return output
}

func getWikiStatus(personList []string) []string {
	urlPrefix := "https://zh.wikipedia.org/wiki/"
	var output []string
	for _, personName := range personList {
		url := urlPrefix + personName
		resp, err := http.Head(url)
		time.Sleep(100)
		check(err)
		// fmt.Println(resp.Status)
		if resp.Status == "200 OK" {
			output = append(output, personName+","+"1")
		} else {
			output = append(output, personName+","+"0")
		}
	}
	return output
}

func outputResult(personListResult []string, outputFileName string) {
	output := strings.Join(personListResult, "\n")
	f, err := os.Create(outputFileName)
	check(err)
	defer f.Close()
	f.WriteString(output)
	f.Sync()
}

func main() {
	personList := readPersonList("personList.txt")
	personListResult := getWikiStatus(personList)
	outputResult(personListResult, "output.txt")
	fmt.Println("done")
}
