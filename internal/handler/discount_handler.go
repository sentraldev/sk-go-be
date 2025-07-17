package handler

import "sk-go-be/internal/service"

type DiscountHandler struct {
	DiscountService service.DiscountService
}

func NewDiscountHandler(discountService service.DiscountService) *DiscountHandler {
	return &DiscountHandler{DiscountService: discountService}
}
