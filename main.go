package main

import (
	"exhttp/handler"
	"exhttp/models"
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SecReg struct {
	Port       string
	DB         string
	DBUser     string
	DBPassword string
	DBServer   string
	DBPort     string
}

func CreateRegistry(Port string, DB string, DBUser string, DBPassword string, DBServer string, DBPort string) *SecReg {
	return &SecReg{Port: Port, DB: DB, DBUser: DBUser, DBPassword: DBPassword, DBServer: DBServer, DBPort: DBPort}
}
func (sr *SecReg) InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", sr.DBServer, sr.DBUser, sr.DBPassword, sr.DB, sr.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	// Auto-migrate your models
	err = db.AutoMigrate(&models.Service{}, &models.Component{}, &models.SecAudit{}, &models.Team{}, &models.SecurityControl{}, &models.ServiceSecurityControl{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
		return nil, err
	}

	log.Println("Database connection and migration successful!")
	return db, nil

}

func (sr *SecReg) SetupRoutes(app *handler.AppDB) {
	//http.HandleFunc("/welcome", app.Welcome)
	http.HandleFunc("/Dashboard", app.DashboardCount)
	http.HandleFunc("/Component/", app.FetchComponentById)
	http.HandleFunc("/Component/Service/", app.FetchComponentByService)
	http.HandleFunc("/Component/List", app.ListComponents)
	http.HandleFunc("/Component/Create", app.CreateComponent)
	http.HandleFunc("/Service/Create", app.CreateService)
	http.HandleFunc("/Service/List", app.ListServices)
	http.HandleFunc("/Service/", app.FetchService)
	http.HandleFunc("/Team/Create", app.CreateTeam)
	http.HandleFunc("/Team/List", app.ListTeam)
	http.HandleFunc("/SecurityControl/Create", app.CreateSecurityControl)
	http.HandleFunc("/SecurityControl/List", app.ListSecurityControl)
	http.HandleFunc("/SecurityControl/AddService", app.MapControlsHandler)
	http.HandleFunc("/Service/SecurityControl/", app.FetchControlsByServiceId)

	err := http.ListenAndServe(":9090", nil)
	log.Println("Server Started at Port :9090")
	if err != nil {
		log.Println("error while trying to start server", err)
	}
}
func main() {

	sr := CreateRegistry("9090", "secureg3", "secureg", "appsec", "postgres", "5432")
	db, err := sr.InitDB()
	if err != nil {
		log.Println("error while trying to start server", err)
	}
	app := &handler.AppDB{DB: db}
	sr.SetupRoutes(app)

}
