package app

import (
	"github.com/calebgregory/full-stack-demo-shopping-cart/order"
	"github.com/calebgregory/full-stack-demo-shopping-cart/product"
	"github.com/calebgregory/full-stack-demo-shopping-cart/profile"
	"github.com/calebgregory/full-stack-demo-shopping-cart/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
)

type App struct {
	db             *gorm.DB
	OrderHandler   order.HttpHandler
	ProductHandler product.HttpHandler
	ProfileHandler profile.HttpHandler
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
		ProfileHandler: profile.New(db),
	}

	return a, nil
}

func (app *App) ListenAndServe(addr string, handler http.Handler) error {
	http.HandleFunc("/orders/add-product", util.AllowCORS(app.OrderHandler.HandleAddProduct))

	http.HandleFunc("/products/get-all", util.AllowCORS(app.ProductHandler.HandleGetAll))
	http.HandleFunc("/products/get-one", util.AllowCORS(app.ProductHandler.HandleGetOne))
	http.HandleFunc("/products/create", util.AllowCORS(app.ProductHandler.HandleCreate))
	http.HandleFunc("/products/update", util.AllowCORS(app.ProductHandler.HandleUpdate))
	http.HandleFunc("/products/delete", util.AllowCORS(app.ProductHandler.HandleDelete))

	http.HandleFunc("/profiles/get-all", util.AllowCORS(app.ProfileHandler.HandleGetAll))
	http.HandleFunc("/profiles/get-one", util.AllowCORS(app.ProfileHandler.HandleGetOne))
	http.HandleFunc("/profiles/create", util.AllowCORS(app.ProfileHandler.HandleCreate))
	http.HandleFunc("/profiles/update", util.AllowCORS(app.ProfileHandler.HandleUpdate))
	http.HandleFunc("/profiles/delete", util.AllowCORS(app.ProfileHandler.HandleDelete))

	return http.ListenAndServe(addr, handler)
}

func (a *App) Close() {
	a.db.Close()
}
