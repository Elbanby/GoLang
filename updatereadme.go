package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var mux sync.Mutex

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
	logError(err)
	return files
}

func updateReadme(path string, files []os.FileInfo) {
	mux.Lock()
	defer mux.Unlock()
	lines := scan(path, "README.md")
	beginIndex := indexOf(lines, "## Begin Directories")
	endIndex := indexOf(lines, "## End Directories")
	if beginIndex == -1 || endIndex == -1 {
		log.Fatal("Can't find the section of the file to append to!")
	}
	var directories []string
	for _, file := range files {
		if []byte(file.Name())[0] != []byte(".")[0] && file.IsDir() {
			directories = append(directories, " * "+file.Name())
		}
	}
	fileData := append(lines[0:beginIndex+1], directories...)
	fileData = append(fileData, lines[endIndex:]...)
	var fileContent string
	for _, line := range fileData {
		fileContent += line
		fileContent += "\n"
	}
	ioutil.WriteFile(path+"/README.md", []byte(fileContent), 0644)
}

func scan(path string, filename string) []string {
	file, err := os.Open(path + "/" + filename)
	logError(err)
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	logError(scanner.Err())
	return lines
}

func indexOf(sli []string, str string) int {
	for i, s := range sli {
		if s == str {
			return i
		}
	}
	return -1
}

func logError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
