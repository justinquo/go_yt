package models

type Address struct {
	State   string `json: "state" bason: "state"`
	City    string `json: "city" bason: "city"`
	Pincode int    `json: "pincode" bason: "pincode"`
}

type User struct {
	Name    string  `json: "name" bason: "user_name"`
	Age     int     `json: "age" bason: "user_age"`
	Address Address `json: "address" bason: "user_address"`
}
