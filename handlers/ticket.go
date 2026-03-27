package handlers

import (
	"net/http"

	"github.com/devjoemedia/chitodopostgress/database"
	"github.com/devjoemedia/chitodopostgress/utils"
)

type Ticket struct {
	ID int `json:"id"`
}

func GetTickets(w http.ResponseWriter, r *http.Request) {
	var tickets []Ticket

	database.DB.Find(&tickets)

	utils.JSON(w, http.StatusOK, tickets)
}
