package model

type User struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Surname          string `json:"surname"`
	NumOfRentedBooks int    `json:"numOfRentedBooks"`
}
