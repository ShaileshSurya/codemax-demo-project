package main

import (
	"context"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

var log = logrus.New()

func main() {
	ctx := context.TODO()
	routeHandler := getRouterWithMiddlewareNegroni(ctx)

	s := &http.Server{
		Addr:    ":" + port,
		Handler: routeHandler,
	}

	log.Info("Email sending service RESTfully... on port : " + port)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start email sending Service : %s", err.Error())
	}
}

func getRouterWithMiddlewareNegroni(ctx context.Context) (n *negroni.Negroni) {
	n = negroni.Classic()

	n.UseHandler(initHandlers(ctx))
	return n
}

func initHandlers(ctx context.Context) (serveMux *http.ServeMux) {
	r := mux.NewRouter().StrictSlash(true)
	serveMux = http.NewServeMux()

	serveMux.Handle("/", r)
	serveMux.Handle("/mail", registerMiddlewares(ctx, r))

	r.HandleFunc("/mail", func(w http.ResponseWriter, r *http.Request) {
		sendMailHandler(ctx, w, r)
	}).Methods("POST")

	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		log.Info(t)
		return nil
	})
	return serveMux
}

func registerMiddlewares(ctx context.Context, router *mux.Router) *negroni.Negroni {
	return negroni.New(
		negroni.HandlerFunc(xAPIAuthenticationMiddleware),
		negroni.HandlerFunc(loggerMiddleware),
		negroni.Wrap(router))
}
