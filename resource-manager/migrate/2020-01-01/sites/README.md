
## `github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-01-01/sites` Documentation

The `sites` SDK allows for interaction with Azure Resource Manager `migrate` (API Version `2020-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-01-01/sites"
```


### Client Initialization

```go
client := sites.NewSitesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SitesClient.DeleteSite`

```go
ctx := context.TODO()
id := sites.NewVMwareSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vmwareSiteValue")

read, err := client.DeleteSite(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.GetSite`

```go
ctx := context.TODO()
id := sites.NewVMwareSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vmwareSiteValue")

read, err := client.GetSite(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.GetSiteHealthSummary`

```go
ctx := context.TODO()
id := sites.NewVMwareSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vmwareSiteValue")

// alternatively `client.GetSiteHealthSummary(ctx, id)` can be used to do batched pagination
items, err := client.GetSiteHealthSummaryComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SitesClient.GetSiteUsage`

```go
ctx := context.TODO()
id := sites.NewVMwareSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vmwareSiteValue")

read, err := client.GetSiteUsage(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.PatchSite`

```go
ctx := context.TODO()
id := sites.NewVMwareSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vmwareSiteValue")

payload := sites.VMwareSite{
	// ...
}


read, err := client.PatchSite(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.PutSite`

```go
ctx := context.TODO()
id := sites.NewVMwareSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vmwareSiteValue")

payload := sites.VMwareSite{
	// ...
}


read, err := client.PutSite(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SitesClient.RefreshSite`

```go
ctx := context.TODO()
id := sites.NewVMwareSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vmwareSiteValue")

read, err := client.RefreshSite(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
