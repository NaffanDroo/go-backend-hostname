package hostname

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/render"
)

func getHostNameMap() (json map[string]string, err error) {
	name, err := os.Hostname()
	if err != nil {
		name = err.Error()
	}
	var to_return map[string]string = map[string]string{"hostname": name}
	return to_return, err
}

func GetHostName(w http.ResponseWriter, r *http.Request) {
	hostname, err := getHostNameMap()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to determine the host name: %s", err.Error()), http.StatusInternalServerError)
	}
	render.JSON(w, r, hostname)
}
