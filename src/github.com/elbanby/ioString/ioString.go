package ioString

import (
  "os"
  "bufio"
)

func Read(fileName string) ([]string, error) {
 array := make([]string,0)
 file, err := os.Open(fileName)

 if err != nil {
  return array, err
 }

 scanner := bufio.NewScanner(file)

 for scanner.Scan() {
  array = append(array, scanner.Text())
 }

 err = file.Close()

 if scanner.Err()!= nil {
  return array, scanner.Err()
 }
 if err != nil {
  return array, err
 }
 return array, err
}
