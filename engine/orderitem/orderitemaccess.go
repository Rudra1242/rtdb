package orderitem

import (
	"rtdb/adapter/OrderItem"
)

type OrderItemAccess struct {
	OrderItemService OrderItem.Service
}

func (oi *OrderItemAccess) GetAllOrderCount() int {
	return oi.OrderItemService.GetAllOrderCount()
}
