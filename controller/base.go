package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres

	"github.com/nilerajput91/userloginapp/middlewares"
	"github.com/nilerajput91/userloginapp/models"
	"github.com/nilerajput91/userloginapp/responses"
)

// App routing struct
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize connect to the database and wire up routes
func (a *App) Initialize(DbHost, DbPort, DbUser, DbName, DbPassword string) {
	var err error
	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	a.DB, err = gorm.Open("postgres", DBURI)
	if err != nil {
		fmt.Printf("\n Cannot connect to database %s", DbName)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the database %s", DbName)
	}

	a.DB.Debug().AutoMigrate(&models.User{}) //database migration

	a.Router = mux.NewRouter().StrictSlash(true)
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.Use(middlewares.SetContentTypeMiddleware) // setting content-type to json

	a.Router.HandleFunc("/", home).Methods("GET")
	a.Router.HandleFunc("/register", a.UserSignUp).Methods("POST")
	a.Router.HandleFunc("/login", a.Login).Methods("POST")
	a.Router.HandleFunc("/logout",a.LogOut).Methods("GET")
}

// RunServer to run the server
func (a *App) RunServer() {
	log.Printf("\nServer starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", a.Router))
}

func home(w http.ResponseWriter, r *http.Request) { // this is the home route
	responses.JSON(w, http.StatusOK, "Welcome To Login APP")
}
