package controllers

import (
	"fmt"
	"gotoko-postgres/app/models"
	"gotoko-postgres/database/seeders"
	"log"
	"math"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/cli"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
	AppConfig *AppConfig
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
	AppURL string
}

type DBconfig struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBPort     string
	DBHost     string
}

type PageLink struct {
	Page int32
	Url string
	IsCurrentPage bool
}

type PaginationLinks struct {
	CurrentPage string
	NextPage string
	PrevPage string
	TotalRows int32
	TotalPages int32
	Links []PageLink
}

type PaginationParams struct {
	Path string
	TotalRows int32
	PerPage int32
	CurrentPage int32
}

func (server *Server) Initialize(appConfig AppConfig, dbConfig DBconfig) {
	fmt.Println("Wellcome to " + appConfig.AppName)

	server.InitializeDB(dbConfig)
	server.InitializeAppConfig(appConfig)
	server.InitializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func (server *Server) InitializeDB(dbConfig DBconfig) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort)

	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed on connecting to Database")
	}

}

func (server *Server) InitializeAppConfig(appconfig AppConfig) {
	server.AppConfig = &appconfig
}

func (server *Server) dbMigrate() {
	for _, model := range models.RegisterModels() {
		err := server.DB.Debug().AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Database migration successfuly")
}

func (server *Server) InitCommands(config AppConfig, dbConfig DBconfig) {
	server.InitializeDB(dbConfig)

	cmdApp := cli.NewApp()
	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context) error {
				server.dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(c *cli.Context) error {
				err := seeders.DBSeed(server.DB)
				if err != nil {
					log.Fatal(err)
				}
				return nil
			},
		},
	}
	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPaginationLinks(config *AppConfig, params PaginationParams) (PaginationLinks, error) {
	var links []PageLink

	totalPage := int32(math.Ceil(float64(params.TotalRows)/ float64(params.PerPage)))

	for i := 1; int32(i) <= totalPage; i++ {
		links = append(links, PageLink{
			Page: int32(i),
			Url: fmt.Sprintf("%s/%s?page=%s", config.AppURL, params.Path, fmt.Sprint(i)),
			IsCurrentPage: int32(i) == params.CurrentPage,
		})
	}

	var nextPage int32
	var prevPage int32

	prevPage = 1
	nextPage = totalPage

	if params.CurrentPage > 2 {
		prevPage = params.CurrentPage - 1
	}

	if params.CurrentPage < totalPage {
		nextPage = params.CurrentPage + 1
	}

	return PaginationLinks{
		CurrentPage: fmt.Sprintf("%s/%s?page=%s", config.AppURL, params.Path, fmt.Sprint(params.CurrentPage)),
		NextPage: fmt.Sprintf("%s/%s?page=%s", config.AppURL, params.Path, fmt.Sprint(nextPage)),
		PrevPage: fmt.Sprintf("%s/%s?page=%s", config.AppURL, params.Path, fmt.Sprint(prevPage)),
		TotalRows: params.TotalRows,
		TotalPages: totalPage,
		Links: links,
	}, nil
}