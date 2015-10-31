package controllers

import (
	"fmt"
	"net/http"
)

type (
	HomeController struct {
	}
)

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (hc HomeController) HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "0.1")
}
