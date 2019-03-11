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
	fileContnent := ""
	for _, line := range lines {
		if line == "## Directories" {
			fileContnent = fmt.Sprintf("%s%s\n ", fileContnent, line)
			for _, file := range files {
				// ignore hidden folders.
				if []byte(file.Name())[0] != []byte(".")[0] && file.IsDir() {
					fileContnent += fmt.Sprintf("* %v\n ", file.Name())
				}
			}
			continue
		}
		fileContnent += line
		fileContnent += "\n"
	}
	ioutil.WriteFile(path+"/README.md", []byte(fileContnent), 0644)
}
