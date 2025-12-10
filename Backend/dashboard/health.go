package dashboard

import (
	"backend/utils"
	"github.com/gin-gonic/gin"
	"sync"
)

// ────────────────────────────────────────────
// MODELS
// ────────────────────────────────────────────

type Platform struct {
	Name string
	URL  string
}

type Microservice struct {
	Name string
	URL  string
}

type HealthResponse struct {
	Status string `json:"status"`
}

// ────────────────────────────────────────────
// PLATFORM LIST
// ────────────────────────────────────────────

var Platforms = []Platform{
	{Name: "closerx", URL: "https://app.closerx.ai/api/health/"},
	{Name: "snowie", URL: "https://app.snowie.ai/api/health/"},
	{Name: "maya", URL: ""}, // handled separately
}

var MayaServices = []Microservice{
	{Name: "agentservice", URL: "https://maya.ravan.ai/agentservice/health"},
	{Name: "authservice", URL: "https://maya.ravan.ai/authservice/health"},
	{Name: "communicationsservice", URL: "https://maya.ravan.ai/communicationsservice/health"},
	{Name: "crmservice", URL: "https://maya.ravan.ai/crmservice/health"},
	{Name: "deviceservice", URL: "https://maya.ravan.ai/deviceservice/health"},
	{Name: "productivityservice", URL: "https://maya.ravan.ai/productivityservice/health"},
	{Name: "scraperservice", URL: "https://maya.ravan.ai/scraperservice/health"},
	{Name: "thunderservice", URL: "https://maya.ravan.ai/thunderservice/health"},
}

// ────────────────────────────────────────────
// RESTY HEALTH CHECKER
// ────────────────────────────────────────────

func CheckExternalHealth(url string) bool {
	resp := &HealthResponse{}

	_, err := Resty.R().
		SetResult(resp).
		Get(url)

	if err != nil {
		return false
	}

	return resp.Status == "ok"
}

// ────────────────────────────────────────────
// CHECK ALL MAYA MICROSERVICES
// ────────────────────────────────────────────

func CheckMayaServices() map[string]bool {
	results := make(map[string]bool)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, svc := range MayaServices {
		wg.Add(1)

		go func(s Microservice) {
			defer wg.Done()

			status := CheckExternalHealth(s.URL)

			mu.Lock()
			results[s.Name] = status
			mu.Unlock()
		}(svc)
	}

	wg.Wait()
	return results
}

// ────────────────────────────────────────────
// MAIN HEALTH CHECK HANDLER
// ────────────────────────────────────────────

func HealthCheckAll(c *gin.Context) {
	results := make(map[string]any)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, p := range Platforms {
		wg.Add(1)

		go func(platform Platform) {
			defer wg.Done()

			// maya → microservices only
			if platform.Name == "maya" {
				maya := CheckMayaServices()

				mu.Lock()
				results["maya"] = maya
				mu.Unlock()
				return
			}

			// normal platforms
			status := false
			if platform.URL != "" {
				status = CheckExternalHealth(platform.URL)
			}

			mu.Lock()
			results[platform.Name] = status
			mu.Unlock()
		}(p)
	}

	wg.Wait()

	utils.CustomResponse(c, 200, true, "Platform health status", results)
}
