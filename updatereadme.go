package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for _, path := range os.Args[1:] {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			updateReadme(path, readDir(path))
		}(path)
	}
	wg.Wait()
	fmt.Println("All Readme.md files updated")
}

func readDir(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func updateReadme(path string, files []os.FileInfo) {
	var beginSection int
	var endSection int
	readmeFile, err := os.Open(path + "/README.md")
	var lines []string
	defer readmeFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(readmeFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatal(err)
	}

	for i, line := range lines {
		if line == "## Begin Directories" {
			beginSection = i
		}
		if line == "## End Directories" {
			endSection = i
		}
	}

	var fileTop = lines[0 : beginSection+1]
	var fileBottom = lines[endSection:]
	var filesSection []string

	for _, file := range files {
		if []byte(file.Name())[0] == []byte(".")[0] || !file.IsDir() {
			continue
		}
		filesSection = append(filesSection, "* "+file.Name())
	}

	completedata := append(fileTop, filesSection...)
	completedata = append(completedata, fileBottom...)
	fmt.Println(completedata)

	var fileContent string
	for _, line := range completedata {
		fileContent += line
		fileContent += "\n"
	}
	ioutil.WriteFile(path+"/README.md", []byte(fileContent), 0644)
}
