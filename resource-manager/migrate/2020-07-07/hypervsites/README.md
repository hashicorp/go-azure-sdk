
## `github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/hypervsites` Documentation

The `hypervsites` SDK allows for interaction with Azure Resource Manager `migrate` (API Version `2020-07-07`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/hypervsites"
```


### Client Initialization

```go
client := hypervsites.NewHyperVSitesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `HyperVSitesClient.DeleteSite`

```go
ctx := context.TODO()
id := hypervsites.NewHyperVSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hyperVSiteName")

read, err := client.DeleteSite(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HyperVSitesClient.GetSite`

```go
ctx := context.TODO()
id := hypervsites.NewHyperVSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hyperVSiteName")

read, err := client.GetSite(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HyperVSitesClient.GetSiteHealthSummary`

```go
ctx := context.TODO()
id := hypervsites.NewHyperVSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hyperVSiteName")

// alternatively `client.GetSiteHealthSummary(ctx, id)` can be used to do batched pagination
items, err := client.GetSiteHealthSummaryComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `HyperVSitesClient.GetSiteUsage`

```go
ctx := context.TODO()
id := hypervsites.NewHyperVSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hyperVSiteName")

read, err := client.GetSiteUsage(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HyperVSitesClient.PatchSite`

```go
ctx := context.TODO()
id := hypervsites.NewHyperVSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hyperVSiteName")

payload := hypervsites.HyperVSite{
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


### Example Usage: `HyperVSitesClient.PutSite`

```go
ctx := context.TODO()
id := hypervsites.NewHyperVSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hyperVSiteName")

payload := hypervsites.HyperVSite{
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


### Example Usage: `HyperVSitesClient.RefreshSite`

```go
ctx := context.TODO()
id := hypervsites.NewHyperVSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hyperVSiteName")

read, err := client.RefreshSite(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
