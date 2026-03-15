// Package lighthouse provides a Go client for the Lighthouse CDN control plane API.
//
// Lighthouse is the managed ingress service for wgmesh networks. It routes
// traffic from public domains to origin servers on the WireGuard mesh.
//
// Usage:
//
//	client := lighthouse.NewClient("https://lighthouse.example.com", "cr_your_api_key")
//	site, err := client.CreateSite(lighthouse.CreateSiteRequest{
//	    Domain: "ollama.mymesh.wgmesh.dev",
//	    Origin: lighthouse.Origin{MeshIP: "10.42.100.1", Port: 11434, Protocol: "http"},
//	})
package lighthouse

import "time"

// TLSMode controls how TLS is handled for a site.
type TLSMode string

const (
	TLSModeAuto   TLSMode = "auto"   // Automatic Let's Encrypt via edge Caddy
	TLSModeCustom TLSMode = "custom" // Customer provides cert
	TLSModeOff    TLSMode = "off"    // HTTP only (not recommended)
)

// SiteStatus tracks the lifecycle of a site registration.
type SiteStatus string

const (
	SiteStatusPendingDNS    SiteStatus = "pending_dns"
	SiteStatusPendingVerify SiteStatus = "pending_verify"
	SiteStatusActive        SiteStatus = "active"
	SiteStatusSuspended     SiteStatus = "suspended"
	SiteStatusDeleted       SiteStatus = "deleted"
	SiteStatusDNSFailed     SiteStatus = "dns_failed"
)

// HealthCheck configures periodic HTTP probing for an origin endpoint.
type HealthCheck struct {
	Path      string        `json:"path"`                // e.g., "/healthz"
	Interval  time.Duration `json:"interval,omitempty"`  // default 10s
	Timeout   time.Duration `json:"timeout,omitempty"`   // default 5s
	Unhealthy int           `json:"unhealthy,omitempty"` // consecutive failures before marking down (default 2)
	Healthy   int           `json:"healthy,omitempty"`   // consecutive successes before marking up (default 2)
}

// Origin defines where traffic should be proxied to.
type Origin struct {
	MeshIP      string      `json:"mesh_ip"`                // WireGuard mesh IP of the origin node
	Port        int         `json:"port"`                   // Port on the origin
	Protocol    string      `json:"protocol"`               // "http" or "https" (to origin)
	HealthCheck HealthCheck `json:"health_check,omitempty"` // Optional HTTP health probe config
}

// Site represents a customer domain routed through the CDN.
type Site struct {
	ID        string     `json:"id"`
	OrgID     string     `json:"org_id"`
	Domain    string     `json:"domain"`
	Origin    Origin     `json:"origin"`
	TLS       TLSMode    `json:"tls"`
	Status    SiteStatus `json:"status"`
	DNSTarget string     `json:"dns_target"` // Where the customer should point DNS
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// CreateSiteRequest is the payload for POST /v1/sites.
type CreateSiteRequest struct {
	Domain string `json:"domain"`
	Origin Origin `json:"origin"`
	TLS    string `json:"tls,omitempty"`
}

// UpdateSiteRequest is the payload for PATCH /v1/sites/{id}.
type UpdateSiteRequest struct {
	Origin *Origin `json:"origin,omitempty"`
	TLS    *string `json:"tls,omitempty"`
}
