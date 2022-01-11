package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, html string) {
	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	t, ok := tc[html]
	if !ok {
		log.Fatal(err)
	}

	buf := new (bytes.Buffer)
	_ = t.Execute(buf, nil)
	
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template", err)
	}
}

// CreateTemplateCache() creates template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		// fmt.Println("rendering ", page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts

	}

	return myCache, nil
}
