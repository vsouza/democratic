package models

import "time"

type User struct {
	ID        string       `json:"id"`
	Email     string       `json:"email"`
	Username  string       `json:"username"`
	FullName  string       `json:"fullName"`
	Gender    int          `json:"gender"`
	Status    int          `json:"status"`
	BirthDate time.Time    `json:"birthDate"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"UpdatedAt"`
	Address   *UserAddress `json:"address"`
}

type UserAddress struct {
	Neighborhood   string `json:"nighborhood"`
	ZipCode        string `json:"zipCode"`
	AddressLineOne string `json:"addressLineOne"`
	AddressLineTwo string `json:"addressLineTwo"`
	AddressType    int    `json:"addressType"`
	Number         int32  `json:"number"`
	City           string `json:"city"`
	State          string `json:"state"`
	Country        string `json:"country"`
}

const (
	Male = iota
	Female
	Another
)

const (
	UserStatusActive = iota
	UserStatusBlocked
	UserStatusDeactive
)

func GetUserByID(id string) *User {
	time := time.Now()
	return &User{
		ID:        "30a0e28b-bff8-4393-8a8e-26a7333658ff",
		Email:     "example@example.com",
		Username:  "example",
		FullName:  "John Doe",
		Gender:    Male,
		Status:    UserStatusBlocked,
		BirthDate: time,
		CreatedAt: time,
		UpdatedAt: time,
		Address: &UserAddress{
			Neighborhood:   "Silicon",
			ZipCode:        "01010101",
			AddressLineOne: "Mapple Street",
			AddressLineTwo: "Head quarter",
			Number:         300,
			City:           "Vienza",
			State:          "California",
			Country:        "USA",
		},
	}

}
