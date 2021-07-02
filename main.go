package user_rest_api

import (
	"github.com/satorunooshie/user-rest-api/infra/database"
	"log"
	"net/http"

	"github.com/satorunooshie/user-rest-api/infra"
	"github.com/satorunooshie/user-rest-api/interface/controller"
	"github.com/satorunooshie/user-rest-api/usecase/user"
)

func main() {
	db := infra.NewMySQLDB()
	userController := controller.NewUserController(
		user.NewUserInteractor(
			database.NewUserRepository(db),
		),
	)
	http.HandleFunc("/users/:id", get(userController.Get))
	http.HandleFunc("/users", get(userController.List))
	http.HandleFunc("/users", post(userController.Post))
	http.HandleFunc("/users", put(userController.Put))
	http.HandleFunc("/users/:id", delete(userController.Delete))
}

func get(api http.HandlerFunc) http.HandlerFunc {
	return httpMethod(api, http.MethodGet)
}

func post(api http.HandlerFunc) http.HandlerFunc {
	return httpMethod(api, http.MethodPost)
}

func put(api http.HandlerFunc) http.HandlerFunc {
	return httpMethod(api, http.MethodPut)
}

func delete(api http.HandlerFunc) http.HandlerFunc {
	return httpMethod(api, http.MethodDelete)
}

func httpMethod(api http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// CORS
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Accept, Origin, x-token")
		// preflight request
		if r.Method == http.MethodOptions {
			return
		}
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			if _, err := w.Write([]byte("Method not allowed")); err != nil {
				log.Println(err)
			}
			return
		}
		w.Header().Add("Content-Type", "application/json")
		api(w, r)
	}
}
