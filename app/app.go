package app

import (
	"github.com/calebgregory/full-stack-demo-shopping-cart/address"
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
	AddressHandler address.HttpHandler
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
		AddressHandler: address.New(db),
		OrderHandler:   order.New(db),
		ProductHandler: product.New(db),
		ProfileHandler: profile.New(db),
	}

	return a, nil
}

func (app *App) ListenAndServe(addr string, handler http.Handler) error {
	http.HandleFunc("/address/get-all", util.AllowCORS(app.AddressHandler.HandleGetAll))
	http.HandleFunc("/address/get-one", util.AllowCORS(app.AddressHandler.HandleGetOne))
	http.HandleFunc("/address/create", util.AllowCORS(app.AddressHandler.HandleCreate))
	http.HandleFunc("/address/update", util.AllowCORS(app.AddressHandler.HandleUpdate))
	http.HandleFunc("/address/delete", util.AllowCORS(app.AddressHandler.HandleDelete))

	http.HandleFunc("/order/add-product", util.AllowCORS(app.OrderHandler.HandleAddProduct))

	http.HandleFunc("/product/get-all", util.AllowCORS(app.ProductHandler.HandleGetAll))
	http.HandleFunc("/product/get-one", util.AllowCORS(app.ProductHandler.HandleGetOne))
	http.HandleFunc("/product/create", util.AllowCORS(app.ProductHandler.HandleCreate))
	http.HandleFunc("/product/update", util.AllowCORS(app.ProductHandler.HandleUpdate))
	http.HandleFunc("/product/delete", util.AllowCORS(app.ProductHandler.HandleDelete))

	http.HandleFunc("/profile/get-all", util.AllowCORS(app.ProfileHandler.HandleGetAll))
	http.HandleFunc("/profile/get-one", util.AllowCORS(app.ProfileHandler.HandleGetOne))
	http.HandleFunc("/profile/create", util.AllowCORS(app.ProfileHandler.HandleCreate))
	http.HandleFunc("/profile/update", util.AllowCORS(app.ProfileHandler.HandleUpdate))
	http.HandleFunc("/profile/delete", util.AllowCORS(app.ProfileHandler.HandleDelete))

	return http.ListenAndServe(addr, handler)
}

func (a *App) Close() {
	a.db.Close()
}
