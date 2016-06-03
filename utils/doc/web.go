package doc

import (
	"net/http"
	"html/template"
)

func (d *Doc) Handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./html/doc.html")

	if err != nil {
		w.WriteHeader(403)
		w.Write([]byte(err.Error()))
	}

	doc := Doc{path:d.path,Host:d.Host,Port:d.Port,Comments:d.Comments}

	err = t.Execute(w, doc)

	if err != nil {
		w.WriteHeader(403)
		w.Write([]byte(err.Error()))
	}
}


