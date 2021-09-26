package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title	string
	Body	[]byte
}

func main() {
	// ルーティング的な
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	// http.HandleFunc("/save/", saveHandler)
	// サーバー起動？
	// 予期せぬエラー時にlog吐く
	log.Fatal(http.ListenAndServe(":8080", nil))
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

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1>" +
			"<form action\"/save/%s\" method=\"POST\">" + 
			"<textarea name=\"body\">%s</textarea><br>" +
			"<input type=\"submit\" value=\"Save\">" +
			"</form>",
		p.Title, p.Title, p.Body)
}
