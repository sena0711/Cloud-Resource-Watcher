package main

import (
	"fmt"
	"log"
	"net/http"
	"aws_asset_tracker_api/internal/routes"
	"aws_asset_tracker_api/internal/database"
)

func main() {
	// DB 연결
	database.ConnectDB()

	// 라우터 설정
	router := routes.SetupRouter()

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
