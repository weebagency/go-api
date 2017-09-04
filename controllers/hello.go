package controllers

import (
	"log"
	"net/http"

	"github.com/weebagency/go-api/models"
)

// HelloController extends Controller
type HelloController struct {
	Controller
}

// Index page
func (c *HelloController) Index(w http.ResponseWriter, r *http.Request) {
	payload := models.GetUsers()
	log.Println(payload)
	c.JSONResponse(w, r, payload)
}
