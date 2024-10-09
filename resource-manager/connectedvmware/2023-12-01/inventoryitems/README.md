
## `github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-12-01/inventoryitems` Documentation

The `inventoryitems` SDK allows for interaction with Azure Resource Manager `connectedvmware` (API Version `2023-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-12-01/inventoryitems"
```


### Client Initialization

```go
client := inventoryitems.NewInventoryItemsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InventoryItemsClient.Create`

```go
ctx := context.TODO()
id := inventoryitems.NewInventoryItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vCenterName", "inventoryItemName")

payload := inventoryitems.InventoryItem{
	// ...
}


read, err := client.Create(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InventoryItemsClient.Delete`

```go
ctx := context.TODO()
id := inventoryitems.NewInventoryItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vCenterName", "inventoryItemName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InventoryItemsClient.Get`

```go
ctx := context.TODO()
id := inventoryitems.NewInventoryItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vCenterName", "inventoryItemName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InventoryItemsClient.ListByVCenter`

```go
ctx := context.TODO()
id := inventoryitems.NewVCenterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vCenterName")

// alternatively `client.ListByVCenter(ctx, id)` can be used to do batched pagination
items, err := client.ListByVCenterComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
