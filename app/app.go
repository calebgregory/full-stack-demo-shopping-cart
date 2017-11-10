package app

import (
	"github.com/calebgregory/full-stack-demo-shopping-cart/product"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
)

type App struct {
	db             *gorm.DB
	ProductHandler product.HttpHandler
}

func New(pathToDb string) (*App, error) {
	db, err := gorm.Open("sqlite3", pathToDb)
	if err != nil {
		return nil, err
	}

	a := &App{
		db:             db,
		ProductHandler: product.New(db),
	}

	return a, nil
}

func (app *App) ListenAndServe(addr string, handler http.Handler) error {
	http.HandleFunc("/products/get-all", app.ProductHandler.HandleGetAll)
	http.HandleFunc("/products/create", app.ProductHandler.HandleCreate)
	return http.ListenAndServe(addr, handler)
}

func (a *App) Close() {
	a.db.Close()
}
