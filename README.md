# lighthouse-go

Go client SDK for the [Lighthouse](https://github.com/atvirokodosprendimai/lighthouse) CDN control plane API.

## Install

```sh
go get github.com/atvirokodosprendimai/lighthouse-go
```

## Usage

```go
import lighthouse "github.com/atvirokodosprendimai/lighthouse-go"

// Create a client
client := lighthouse.NewClient("https://lighthouse.mymesh.wgmesh.dev", "cr_your_api_key")

// Register a site
site, err := client.CreateSite(lighthouse.CreateSiteRequest{
    Domain: "ollama.mymesh.wgmesh.dev",
    Origin: lighthouse.Origin{
        MeshIP:   "10.42.100.1",
        Port:     11434,
        Protocol: "http",
    },
})

// List sites
sites, err := client.ListSites()

// Get a site
site, err := client.GetSite("site_abc123")

// Update a site
site, err := client.UpdateSite("site_abc123", lighthouse.UpdateSiteRequest{
    Origin: &lighthouse.Origin{Port: 8080, Protocol: "http", MeshIP: "10.42.100.1"},
})

// Delete a site
err := client.DeleteSite("site_abc123")
```

## Discovery

Find the Lighthouse URL for a mesh automatically via DNS SRV:

```go
url, err := lighthouse.DiscoverLighthouse("mymeshid")
// Tries: _lighthouse._tcp.mymeshid.wgmesh.dev (SRV)
// Falls back to: https://lighthouse.mymeshid.wgmesh.dev
```

## License

Apache-2.0
