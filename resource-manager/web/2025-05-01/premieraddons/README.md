
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/premieraddons` Documentation

The `premieraddons` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/premieraddons"
```


### Client Initialization

```go
client := premieraddons.NewPremierAddonsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PremierAddonsClient.WebAppsAddPremierAddOn`

```go
ctx := context.TODO()
id := premieraddons.NewPremierAddonID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "premierAddonName")

payload := premieraddons.PremierAddOn{
	// ...
}


read, err := client.WebAppsAddPremierAddOn(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PremierAddonsClient.WebAppsDeletePremierAddOn`

```go
ctx := context.TODO()
id := premieraddons.NewPremierAddonID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "premierAddonName")

read, err := client.WebAppsDeletePremierAddOn(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PremierAddonsClient.WebAppsGetPremierAddOn`

```go
ctx := context.TODO()
id := premieraddons.NewPremierAddonID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "premierAddonName")

read, err := client.WebAppsGetPremierAddOn(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PremierAddonsClient.WebAppsUpdatePremierAddOn`

```go
ctx := context.TODO()
id := premieraddons.NewPremierAddonID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "premierAddonName")

payload := premieraddons.PremierAddOnPatchResource{
	// ...
}


read, err := client.WebAppsUpdatePremierAddOn(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
