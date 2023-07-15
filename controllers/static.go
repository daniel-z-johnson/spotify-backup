package controllers

import "net/http"

func StaticPage(template Template) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		template.Execute(resp, req, nil)
	}
}
