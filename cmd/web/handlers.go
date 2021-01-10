package main

import (
	"awesomeProject3/pkg/models"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request){
	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := &templateData{Snippets: s}

	app.render(w, r, "home.page.tmpl", data)
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id,err := strconv.Atoi(r.URL.Query().Get(":id"))
	if  err!=nil || id<0 {
		app.notFound(w)
		return
	}

	s,err := app.snippets.Get(id)

	data := &templateData{Snippet: s}

	if err != nil {
		if errors.Is(err, models.ErrNoRecord){
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, r, "show.page.tmpl", data)


}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	title := "0 snail"
	content := "0 snail\nClimb Mount Fuji,\\nBut slowly, slowly!\\n\\n– Kobayashi Issa"
	expires := "7"

	id,err := app.snippets.Insert(title, content, expires)
	if err!= nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}