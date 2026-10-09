package piscine

type Food struct {
	name     string
	preptime int
}

func FoodDeliveryTime(order string) int {
	food := []Food{
		{"burger", 15},
		{"chips", 10},
		{"nuggets", 12},
	}
	for _, v := range food {
		if v.name == order {
			return v.preptime
		}
	}
	return 404
}
