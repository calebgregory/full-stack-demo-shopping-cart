package app

import (
	"github.com/calebgregory/full-stack-demo-shopping-cart/order"
	"github.com/calebgregory/full-stack-demo-shopping-cart/product"
	"github.com/calebgregory/full-stack-demo-shopping-cart/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
)

type App struct {
	db             *gorm.DB
	OrderHandler   order.HttpHandler
	ProductHandler product.HttpHandler
}

func New(pathToDb string) (*App, error) {
	db, err := gorm.Open("sqlite3", pathToDb)
	if err != nil {
		return nil, err
	}

	a := &App{
		db:             db,
		OrderHandler:   order.New(db),
		ProductHandler: product.New(db),
	}

	return a, nil
}

func (app *App) ListenAndServe(addr string, handler http.Handler) error {
	http.HandleFunc("/orders/add-product", util.AllowCORS(app.OrderHandler.HandleAddProduct))
	http.HandleFunc("/products/get-all", util.AllowCORS(app.ProductHandler.HandleGetAll))
	http.HandleFunc("/products/create", util.AllowCORS(app.ProductHandler.HandleCreate))
	return http.ListenAndServe(addr, handler)
}

func (a *App) Close() {
	a.db.Close()
}
