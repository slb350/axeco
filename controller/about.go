package controller

import (
	"net/http"

	"github.com/slb350/axeco/shared/view"
)

// Displays the About page
func AboutGET(w http.ResponseWriter, r *http.Request) {
	// Display the view
	v := view.New(r)
	v.Name = "about/about"
	v.Render(w)
}
