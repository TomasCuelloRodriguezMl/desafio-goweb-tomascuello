package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"desafio-goweb-tomascuello/cmd/server/handler"
	"desafio-goweb-tomascuello/internal/domain"
	"desafio-goweb-tomascuello/internal/tickets"

	"github.com/gin-gonic/gin"
)

func main() {

	// Cargo csv.
	//list, err := Load... cambiada para evitar error de no uso
	list, err := LoadTicketsFromFile("../../tickets.csv")
	if err != nil {
		panic("Couldn't load tickets")
	}
	db := list
	repo := tickets.NewRepository(db)
	service := tickets.NewService(repo)
	tkts := handler.NewService(service)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	r.GET("/ticket/getByCountry/:dest", tkts.GetTicketsByCountry())
	r.GET("/ticket/getAverage/:dest", tkts.AverageDestination())
	// Rutas a desarollar:

	// GET - “/ticket/getByCountry/:dest”
	// GET - “/ticket/getAverage/:dest”
	if err := r.Run(); err != nil {
		panic(err)
	}

}

func LoadTicketsFromFile(path string) ([]domain.Ticket, error) {

	var ticketList []domain.Ticket

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	for _, row := range data {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return []domain.Ticket{}, err
		}
		ticketList = append(ticketList, domain.Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
	}

	return ticketList, nil
}
