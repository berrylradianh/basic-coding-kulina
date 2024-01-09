package dashboard

import (
	dc "basic-coding-kulina/modules/usecase/admin/dashboard"
)

type DashboardHandler struct {
	dashboardUsecase dc.DashboardUsecase
}

func New(dashboardUsecase dc.DashboardUsecase) *DashboardHandler {
	return &DashboardHandler{
		dashboardUsecase,
	}
}
