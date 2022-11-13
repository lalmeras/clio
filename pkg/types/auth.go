package types

import (
	"time"
)

type ApiCredential struct {
	AllowedIPs    []string     `json:"allowedIPs"`
	ApplicationId int64        `json:"applicationId"`
	Creation      time.Time    `json:"creation"`
	CredentialId  int64        `json:"credentialId"`
	LastUse       time.Time    `json:"lastUse"`
	OvhSupport    bool         `json:"ovhSupport"`
	Rules         []AccessRule `json:"rules"`
	Status        string       `json:"status"`
}

type Details struct {
	AllowedRoutes []AccessRule `json:"allowedRoutes"`
	Description   string       `json:"description"`
	Method        string       `json:"method"`
	Roles         []string     `json:"roles"`
	User          []string     `json:"user"`
}

type AccessRule struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}
