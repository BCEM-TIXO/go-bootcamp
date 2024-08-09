package handlers

import (
	"ex02/server/model"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func (a App) Article(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	file, err := os.Open(fmt.Sprintf("./html/%s.html", title))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ar_, _ := io.ReadAll(file)
	ar := string(ar_)
	data := model.Article{Title: model.ArticleTitle{Title: title},
		Content: template.HTML(ar),
	}
	if err := a.Templates["article"].Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
