package Models

type Order struct{
	Id int
	Customer Account
	Products []Product
	TotalPrice float32
}