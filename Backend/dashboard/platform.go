package dashboard

import (
	"backend/utils"
	"github.com/gin-gonic/gin"
	"sync"
)

type Platform struct {
	Name string
	URL  string
}

// MAIN PLATFORMS
var Platforms = []Platform{
	{Name: "closerx", URL: "https://app.closerx.ai/api/health/"},
	{Name: "snowie", URL: "https://app.snowie.ai/api/health/"},
	{Name: "maya", URL: ""}, // special case → check microservices
}

type Microservice struct {
	Name string
	URL  string
}

// Maya microservices (update URLs to real ones)
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
// CHECK ALL MAYA MICROSERVICES
// ────────────────────────────────────────────
func CheckMayaServices() map[string]bool {
	results := make(map[string]bool)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, svc := range MayaServices {
		wg.Add(1)

		go func(service Microservice) {
			defer wg.Done()

			status, err := CheckExternalHealth(service.URL)
			if err != nil {
				status = false
			}

			mu.Lock()
			results[service.Name] = status
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

			// Special case for Maya → check microservices
			if platform.Name == "maya" {
				mayaStatus := CheckMayaServices()

				mu.Lock()
				results["maya"] = mayaStatus
				mu.Unlock()
				return
			}

			// Normal platforms
			status := false

			if platform.URL != "" {
				s, err := CheckExternalHealth(platform.URL)
				if err == nil {
					status = s
				}
			}

			mu.Lock()
			results[platform.Name] = status
			mu.Unlock()
		}(p)
	}

	wg.Wait()

	utils.CustomResponse(c, 200, true, "Platform health status", results)
}
