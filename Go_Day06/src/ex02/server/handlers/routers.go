package handlers

import (
	"ex02/server/converter"
	"ex02/server/credentials"
	"ex02/server/database"
	"ex02/server/limiter"

	"github.com/gorilla/mux"

	"html/template"
	"log"
	"net/http"
	"time"
)

type App struct {
	Router      *mux.Router
	DB          *database.Database
	Logger      *log.Logger
	Limiter     *limmiter.RateLimiter
	Converter   converter.Converter
	Templates   map[string]*template.Template
	adminLogin  credentials.AdminContent
	dbLogin     credentials.DBContent
	auth        map[string]bool
	articlePath string
	savePath    string
}

func (a *App) Initialize(logger *log.Logger) {
	var err error
	a.auth = make(map[string]bool)
	a.adminLogin, err = credentials.GetAdminCredentials()
	if err != nil {
		log.Fatal(err)
	}
	a.dbLogin, err = credentials.GetDBCredentials()
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
	a.DB = database.NewDB(a.dbLogin.Login, a.dbLogin.Password, "postgres", "localhost", 5332)
	a.Logger = logger
	a.readTemplates()
	a.Limiter = limmiter.NewRateLimiter()
	a.initializeRoutes()
	a.articlePath = "./articles/"
	a.savePath = "./html/"
	a.Converter = converter.Converter{SavePath: a.savePath}
}

func (a *App) initializeRoutes() {
	routes := []struct {
		Name        string
		Method      string
		Pattern     string
		HandlerFunc http.HandlerFunc
	}{
		{Name: "Logo", Method: "GET", Pattern: "/img", HandlerFunc: func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "./img/amazing_logo.png") }},
		{Name: "Article", Method: "GET", Pattern: "/article", HandlerFunc: a.Article},
		{Name: "IndexRedir", Method: "GET", Pattern: "/", HandlerFunc: func(w http.ResponseWriter, r *http.Request) { http.Redirect(w, r, "/main", http.StatusSeeOther) }},
		{Name: "Index", Method: "GET", Pattern: "/main", HandlerFunc: a.Index},
		{Name: "AdminLogin", Method: "POST", Pattern: "/login", HandlerFunc: a.AdminLogin},
		{Name: "LoginPage", Method: "GET", Pattern: "/login", HandlerFunc: a.AdminLoginPage},
		{Name: "Admin", Method: "POST", Pattern: "/admin", HandlerFunc: a.AdminPost},
		{Name: "Admin", Method: "GET", Pattern: "/admin", HandlerFunc: a.AdminPage},
	}

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name, a.Logger)
		if route.Pattern == "/admin" {
			handler = a.Authorized(handler)
		}
		handler = a.Limiter.LimiterMidlewire(handler)
		a.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
}

func Logger(inner http.Handler, name string, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		logger.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func (a *App) readTemplates() {
	a.Templates = make(map[string]*template.Template)
	a.Templates["index"] = template.Must(template.ParseFiles("templates/index.html"))
	a.Templates["article"] = template.Must(template.ParseFiles("templates/article.html"))
	a.Templates["adminLogin"] = template.Must(template.ParseFiles("templates/adminLogin.html"))
	a.Templates["adminPost"] = template.Must(template.ParseFiles("templates/adminPost.html"))
}
