package server

import (
	"fmt"
	"github.com/bmizerany/pat"
	"github.com/cmdrkeene/placebunny.com/bunny"
	"log"
	"net/http"
)

const EMOTICON = `(\__/)` + "\n" + `(='.'=)` + "\n" + `(")_(")` + "\n"

func Start(port string) {
	log.Print("starting server")
	http.Handle("/", Handler())
	if e := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); e != nil {
		log.Fatal("unable to start server: " + e.Error())
	}
}

func Handler() http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(home))
	mux.Get("/:x/:y", http.HandlerFunc(scaled))
	return mux
}

func home(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, EMOTICON)
}

func scaled(w http.ResponseWriter, req *http.Request) {
	x := req.URL.Query().Get(":x")
	y := req.URL.Query().Get(":y")

	w.Header().Set("Content-Type", "image/jpeg")
	bunny.Write(w, x, y)
}

func errorHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err, ok := recover().(error); ok {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()
		fn(w, r)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
