package main

import (
	"log"
	httpapi "loyalid-test/api/http"
	appProduct "loyalid-test/application/product"
	appUser "loyalid-test/application/user"
	"loyalid-test/repository/sql"
	"os"
)

var (
	userService    *appUser.Service
	productService *appProduct.Service
)

func main() {
	connStr := os.Getenv("DATABASE_URL")
	sqlRepo, err := sql.NewSQLRepository("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return
	}

	userService = appUser.New()
	productService = appProduct.New(sqlRepo)

	httpapi.InitRoutes(userService, productService).Run(":8080")
}
