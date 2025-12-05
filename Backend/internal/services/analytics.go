package services

import "backend/internal/stripe"

type Analytics struct {
	Revenue  int64 `json:"revenue"`
	Refunded int64 `json:"refunded"`
	Disputes int64 `json:"disputes"`
	Profit   int64 `json:"profit"`
}

type AnalyticsService struct {
	client *stripeclient.Client
}

func NewAnalyticsService(c *stripeclient.Client) *AnalyticsService {
	return &AnalyticsService{client: c}
}

func (s *AnalyticsService) GetAnalytics() (*Analytics, error) {
	revenue, _ := s.client.GetRevenue()
	refunded, _ := s.client.GetRefunded()
	disputes, _ := s.client.GetDisputesLost()
	profit := revenue - refunded - disputes

	return &Analytics{
		Revenue:  revenue,
		Refunded: refunded,
		Disputes: disputes,
		Profit:   profit,
	}, nil
}
