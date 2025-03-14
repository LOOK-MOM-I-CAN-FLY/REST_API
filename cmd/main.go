package main

import (
	"compoze_rest_api/internal/build"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const (
	apiPrefix string = "/api/v1"
)

var (
	port                    string
	bookResourcePrefix      string = apiPrefix + "/book"  //api/v1/book/
	manyBooksResourcePrefix string = apiPrefix + "/books" //api/v1/book/
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Не найден .env файл")
	}
	port = os.Getenv("PORT")
}
func main() {

	log.Println("Создание простого REST API сервера произошло успешно на порту " + port)

	router := gin.Default()
	log.Println("Роутер инициализирован")

	build.BuildBookResource(router, bookResourcePrefix)
	build.BuildManyBooksResourcePrefix(router, manyBooksResourcePrefix)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
