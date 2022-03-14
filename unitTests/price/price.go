package price

func TotalPrice(nights, rate, cityTax uint) uint {
	return nights * (rate + cityTax)
}