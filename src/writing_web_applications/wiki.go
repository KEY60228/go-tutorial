package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title	string
	Body	[]byte
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save();
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	// ファイルの書き出し (永続化)
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	// ファイルの読み込み
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
