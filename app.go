package main

import (
  "fmt"
  "io/ioutil"
)

// a page data structure
type Page struct {
  Title string
  Body []byte
}

// create a save method on Page
func (p *Page) save() error {
  filename := p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
  filename := title + ".txt"
  body, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  return &Page{Title: title, Body: body}, nil
}

func main() {
  p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page")}
  // fmt.Println(p1)
  p1.save()
  p2, _ := loadPage("TestPage")
  fmt.Println(string(p2.Body))
}