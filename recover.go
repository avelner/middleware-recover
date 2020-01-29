package recover

import (
	"fmt"
	"html"
	"net/http"
	"runtime/debug"
)

func Do(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				// fmt.Fprint(w, "<pre>")
				fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v", r)))
				fmt.Fprint(w, "\n\n")
				fmt.Fprint(w, html.EscapeString(string(debug.Stack())))
				// fmt.Fprint(w, "</pre>")
			}
		}()
		next.ServeHTTP(w, req)
	})
}
