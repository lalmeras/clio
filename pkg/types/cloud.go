package types

import "time"

type Project struct {
	Access       string    `json:"access"`
	CreationDate time.Time `json:"creationDate"`
	Description  string    `json:"description"`
	Expiration   time.Time `json:"expiration"`
	ManualQuota  bool      `json:"manualQuota"`
	OrderId      int64     `json:"orderId"`
	PlanCode     string    `json:"planCode"`
	ProjectName  string    `json:"projectName"`
	ProjectID    string    `json:"project_id"`
	Status       string    `json:"status"`
	Unleash      bool      `json:"unleash"`
}
