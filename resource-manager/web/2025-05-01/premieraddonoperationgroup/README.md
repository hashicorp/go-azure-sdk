
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/premieraddonoperationgroup` Documentation

The `premieraddonoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/premieraddonoperationgroup"
```


### Client Initialization

```go
client := premieraddonoperationgroup.NewPremierAddOnOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PremierAddOnOperationGroupClient.WebAppsAddPremierAddOnSlot`

```go
ctx := context.TODO()
id := premieraddonoperationgroup.NewSlotPremierAddonID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "premierAddonName")

payload := premieraddonoperationgroup.PremierAddOn{
	// ...
}


read, err := client.WebAppsAddPremierAddOnSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PremierAddOnOperationGroupClient.WebAppsDeletePremierAddOnSlot`

```go
ctx := context.TODO()
id := premieraddonoperationgroup.NewSlotPremierAddonID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "premierAddonName")

read, err := client.WebAppsDeletePremierAddOnSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PremierAddOnOperationGroupClient.WebAppsGetPremierAddOnSlot`

```go
ctx := context.TODO()
id := premieraddonoperationgroup.NewSlotPremierAddonID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "premierAddonName")

read, err := client.WebAppsGetPremierAddOnSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PremierAddOnOperationGroupClient.WebAppsUpdatePremierAddOnSlot`

```go
ctx := context.TODO()
id := premieraddonoperationgroup.NewSlotPremierAddonID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "premierAddonName")

payload := premieraddonoperationgroup.PremierAddOnPatchResource{
	// ...
}


read, err := client.WebAppsUpdatePremierAddOnSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
