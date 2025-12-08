package services

import (
	"fmt"
	stripeclient "backend/clients"
)

type Analytics struct {
	Revenue  int64 `json:"revenue"`
	Refunded int64 `json:"refunded"`
	Disputes int64 `json:"disputes"`
	Profit   int64 `json:"profit"`
}

type AnalyticsService struct {
	clients map[string]*stripeclient.Client
}

func NewAnalyticsService(c map[string]*stripeclient.Client) *AnalyticsService {
	return &AnalyticsService{clients: c}
}

func (s *AnalyticsService) GetAnalytics(product string) (*Analytics, error) {
	
	client := s.clients[product]
	if client == nil {
		return nil, fmt.Errorf("invalid product")
	}

	revenue, _ := client.GetRevenue()
	refunded, _ := client.GetRefunded()
	disputes, _ := client.GetDisputesLost()

	profit := revenue - refunded - disputes

	return &Analytics{
		Revenue:  revenue,
		Refunded: refunded,
		Disputes: disputes,
		Profit:   profit,
	}, nil
}
