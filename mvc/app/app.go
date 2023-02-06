// + build ignore
package app

import (
	"net/http"

	"github.com/prosline/micro-services/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("Error: Author name cannot be nil")
	}

}
