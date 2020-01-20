package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "njnvxqxp"
	DB_PASSWORD = "rNuckvNyXEZimXg7u3egwhEXg_eWhRAz"
	DB_NAME     = "njnvxqxp"
	DB_HOST     = "rajje.db.elephantsql.com"
)

var router *gin.Engine
var db *sqlx.DB

type Patient struct {
	id         int    `json:"id" db:"id"`
	name       string `json:"name" db:"name"`
	illness    string `json:"illness" db:"illness"`
	birthdate  string `json:"birth_date" db:"birth_date"`
	locationid int    `json:"location_id" db:"location_id"`
}

type Location struct {
	id         int    `json:"id" db:"id"`
	name       string `json:"name" db:"name"`
	hospitalid int    `json:"hospital_id" db:"hospital_id"`
}

type Hospital struct {
	id               int    `json:"id" db:"id"`
	name             string `json:"name" db:"name"`
	maxpatientamount int   `json:"max_patient_amount" db:"max_patient_amount"`
}

func main() {
	initDatabase()
	initAPI()
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func initDatabase() {

	psqlInfo := fmt.Sprintf("user=%v password=%v dbname=%v host=%v sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME, DB_HOST)
	var err error

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("ALIVE")
}

func initAPI() {
	router = gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/greet/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello " + name + "!",
		})
	})

	router.GET("/patient/all", func(c *gin.Context) {
		var patients []Patient
		err := db.Select(&patients, "SELECT * FROM patient")
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, patients)
	})

	router.GET("/location/all", func(c *gin.Context) {
		var locations []Location
		err := db.Select(&locations, "SELECT * FROM location")
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, locations)
	})

	router.GET("/hospital/all", func(c *gin.Context) {
		var hospitals []Hospital
		err := db.Select(&hospitals, "SELECT * FROM hospital")
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, hospitals)
	})
}
