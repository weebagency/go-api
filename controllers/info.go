package controllers

import (
	"net/http"

	"apps/api/models"
)

// InfoController extends Controller
type InfoController struct {
	Controller
}

// Info about info
func (c *InfoController) Info(w http.ResponseWriter, r *http.Request) {
	si := models.GetServiceInfo()
	c.JSONResponse(w, r, si)
}
