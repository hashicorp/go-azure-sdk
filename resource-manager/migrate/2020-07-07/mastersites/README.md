
## `github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/mastersites` Documentation

The `mastersites` SDK allows for interaction with Azure Resource Manager `migrate` (API Version `2020-07-07`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/mastersites"
```


### Client Initialization

```go
client := mastersites.NewMasterSitesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MasterSitesClient.DeleteSite`

```go
ctx := context.TODO()
id := mastersites.NewMasterSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "masterSiteValue")

read, err := client.DeleteSite(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MasterSitesClient.GetSite`

```go
ctx := context.TODO()
id := mastersites.NewMasterSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "masterSiteValue")

read, err := client.GetSite(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MasterSitesClient.List`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MasterSitesClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MasterSitesClient.PatchSite`

```go
ctx := context.TODO()
id := mastersites.NewMasterSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "masterSiteValue")

payload := mastersites.MasterSite{
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


### Example Usage: `MasterSitesClient.PutSite`

```go
ctx := context.TODO()
id := mastersites.NewMasterSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "masterSiteValue")

payload := mastersites.MasterSite{
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
