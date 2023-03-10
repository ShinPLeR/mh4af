package open

import (
	"fmt"
	"html/template"
	"mh4af/internal/model"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("web/templates/index.html.tmpl"))
}

func RunServer(port int, model model.RearrangedDefinition) {
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web/"))))
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		if err := tpl.ExecuteTemplate(w, "index.html.tmpl", model); err != nil {
			fmt.Println(err)
		}
	})
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
