package main

import (
  "fmt"
  "os"
  "net/http"
  "github.com/PuerkitoBio/goquery"
)

func main(){
  url := os.Args[1]

  res, err := http.Get(url)
  defer res.Body.Close()

  doc, err := goquery.NewDocumentFromReader(res.Body)

  if err != nil {
    fmt.Println("error")
    return
  }

  link := doc.Find("link[rel='canonical']")

  if link == nil{
    fmt.Println(url)
  }

  can_url, found := link.Attr("href")
  if found {
    fmt.Println(can_url)
  } else{
    fmt.Println(url)
  }
}
