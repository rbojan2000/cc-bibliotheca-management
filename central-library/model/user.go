package model

type User struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Surname          string `json:"surname"`
	MembershipCard   string `json:"membershipCard"`
	NumOfRentedBooks int    `json:"numOfRentedBooks"`
}
