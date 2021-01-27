package goadmin

import (
	"context"
	"github.com/DiaElectronics/lea-central-wash/cmd/storage/internal/goadmin/pages"
	"github.com/DiaElectronics/lea-central-wash/cmd/storage/internal/goadmin/washing"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin" // Import the adapter, it must be imported. If it is not imported, you need to define it yourself.
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres" // Import the sql driver
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	_ "github.com/GoAdminGroup/themes/adminlte" // Import the theme

	"github.com/gin-gonic/gin"

	"github.com/powerman/pqx"

	app "github.com/DiaElectronics/lea-central-wash/cmd/storage/internal/app"
)

var Repo app.Repo

func StartAdmin(conf pqx.Config) {
	pages.PagesRepo = Repo

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r := gin.Default()

	eng := engine.Default()
	template.AddComp(chartjs.NewChart())

	cfg := config.Config{
		Databases: config.DatabaseList{
			"default": {
				Host:       conf.Host,
				Port:       strconv.Itoa(conf.Port),
				User:       conf.User,
				Pwd:        conf.Pass,
				Name:       "postgres",
				MaxIdleCon: 50,
				MaxOpenCon: 150,
				Driver:     db.DriverPostgresql,
			},
		},
		UrlPrefix: "admin",
		IndexUrl:  "/",
		Debug:     true,
		Language:  language.EN,
	}

	_ = eng.AddConfig(cfg).
		AddGenerators(washingadmin.Generators).
		Use(r)
	r.Static("/uploads", "./uploads")

	eng.HTML("GET", "/admin", pages.GetDashboardPage)

	srv := &http.Server{
		Addr:    ":9033",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
