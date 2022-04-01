package hostname

import (
	"net/http"
	"os"

	"github.com/go-chi/render"
)

func getHostNameMap() map[string]string {
	name, _ := os.Hostname()
	var to_return map[string]string = map[string]string{"hostname": name}
	return to_return
}

func GetHostName(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, getHostNameMap())
}
