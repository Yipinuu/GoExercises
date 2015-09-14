package main

import (
  "fmt"
  "time"
)

var week time.Duration
func main() {
  t := time.Now()
  fmt.Println(t) //e.g. 2015-09-14 18:03:37.807328833 +0800 CST
  fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year()) //e.g. 14.09.2015

  t = time.Now().UTC()
  fmt.Println(t) //e.g. 2015-09-14 10:03:37.807844123 +0000 UTC
  fmt.Println(time.Now()) //2015-09-14 18:03:37.807849935 +0800 CST

  week = 60*60*24*7*1e9
  week_from_now := t.Add(week)
  fmt.Println(week_from_now) //2015-09-21 10:03:37.807844123 +0000 UTC

  fmt.Println(t.Format(time.RFC822)) //e.g. 14 Sep 15 10:04 UTC
  fmt.Println(t.Format(time.ANSIC)) //e.g. Mon Sep 14 10:04:36 2015
  fmt.Println(t.Format("02 Jan 2006 15:04")) //e.g. 4 Sep 2015 10:08
  s := t.Format("20060102")
  fmt.Println(t, "=>", s) //e.g. 2015-09-14 10:09:20.304262072 +0000 UTC => 20150914
}
