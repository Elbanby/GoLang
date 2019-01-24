package main

import (
  "errors"
  "fmt"
  "log"
)

type Date struct {
  day   int
  month int
  year  int
}

func (d *Date) SetYear(year int) error {
  if year < 1 {
    return errors.New("Invalid year")
  }
  d.year = year
  return nil
}

func (d *Date) SetMonth(month int) {
  d.month = month
}

func (d *Date) SetDay(day int) {
  d.day = day
}

func (d *Date) Year() int {
  return d.year
}

func (d *Date) Month() int {
  return d.month
}

func (d *Date) Day() int {
  return d.day
}

func main() {
  date := Date{}
  err := date.SetYear(2010)
  if err != nil {
    log.Fatal(err)
  }

  date.SetDay(32)
  date.SetMonth(4)
  fmt.Printf("%#v\n",date)
  fmt.Printf("Year is %d\n", date.Year() )
  fmt.Printf("Month is %d\n", date.Month() )
  fmt.Printf("Day is %d\n", date.Day() )
}
