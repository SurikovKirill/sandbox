package admin

import "context"

func (h *AdminHandler) Create() {
	h.service.GetItems(context.TODO())
}
