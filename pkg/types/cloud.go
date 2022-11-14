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

type Image struct {
	CreationDate time.Time `json:"creationDate"`
	FlavorType   string    `json:"flavorType"`
	ID           string    `json:"id"`
	MinDisk      int64     `json:"minDisk"`
	MinRam       int64     `json:"minRam"`
	Name         string    `json:"name"`
	PlanCode     string    `json:"planCode"`
	Region       string    `json:"region"`
	Size         float32   `json:"size"`
	Status       string    `json:"status"`
	Tags         []string  `json:"tags"`
	Type         string    `json:"type"`
	User         string    `json:"user"`
	Visibility   string    `json:"visibility"`
}
