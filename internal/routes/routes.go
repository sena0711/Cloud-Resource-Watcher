package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// 라우터 설정
func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods("GET")
	return router
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "AWS Asset Tracker API is running!") // 변경된 코드
}
