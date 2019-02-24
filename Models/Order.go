package Models

type Order struct{
	Id int
	Customer Account
	Products []Item
	TotalPrice float32
}

func (order *Order) AddItem(orderId int, item Item) error {
	//TODO: Get order from database

	order.Products = append(order.Products, item)
	order.CountTotalPrice()
	return nil
}

func (order *Order) RemoveItem(orderId int, item Item) error {
	var index int

	for i := 0; i < len(order.Products); i++{
		if order.Products[i].Id == item.Id {
			index = i
			break
		}
	}
	order.Products = append(order.Products[:index], order.Products[index+1:]...)
	order.CountTotalPrice()

	return nil
}

func (order *Order) CountTotalPrice() {
	order.TotalPrice = 0.0

	if len(order.Products) > 0 {
		for _, value := range order.Products {
			order.TotalPrice += value.Price
		}
	}
}
