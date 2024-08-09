package handlers

import (
	"ex01/server/model"
	"fmt"
	"net/http"
	"strconv"
)

func (a App) Index(w http.ResponseWriter, r *http.Request) {
	pageSize := 3
	db := a.DB
	pageCount := db.PageCount(pageSize)
	if pageCount == 0 {
		pageCount = 1
	}
	var page int
	var err error
	pageStr := r.URL.Query().Get("page")
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		page = 1
	}
	if page < 1 {
		http.Redirect(w, r, "/main", http.StatusSeeOther)
		return
	} else if page > pageCount {
		http.Redirect(w, r, fmt.Sprintf("/main?page=%d", pageCount), http.StatusSeeOther)
		return
	}
	pag := model.Paging{
		Page:     page,
		PrevPage: page - 1,
		NextPage: page + 1,
		HasPrev:  page != 1 || page <= pageCount,
		HasNext:  page < pageCount,
	}
	// log.Println(db.GetPage(page, pageSize))
	data := model.HomePage{
		Title:          "Home",
		Paging:         pag,
		ArticlesTitles: db.GetPage(page, pageSize),
	}

	// fmt.Println(r.Header)
	if err := a.Templates["index"].Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
