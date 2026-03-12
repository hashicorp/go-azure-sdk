
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/bysiteslotoperationgroup` Documentation

The `bysiteslotoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/bysiteslotoperationgroup"
```


### Client Initialization

```go
client := bysiteslotoperationgroup.NewBySiteSlotOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BySiteSlotOperationGroupClient.ResourceHealthMetadataGetBySiteSlot`

```go
ctx := context.TODO()
id := bysiteslotoperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

read, err := client.ResourceHealthMetadataGetBySiteSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BySiteSlotOperationGroupClient.ResourceHealthMetadataListBySiteSlot`

```go
ctx := context.TODO()
id := bysiteslotoperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.ResourceHealthMetadataListBySiteSlot(ctx, id)` can be used to do batched pagination
items, err := client.ResourceHealthMetadataListBySiteSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
