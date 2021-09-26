package main

import (
	"html/template"
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
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// ---.htmlを読んで*template.Templateを返す
	t, _ := template.ParseFiles(tmpl + ".html")
	// HTMLを生成してhttp.ResponseWriter(w)に書き込む
	// HTML内の.---はpが補完
	t.Execute(w, p)
}
