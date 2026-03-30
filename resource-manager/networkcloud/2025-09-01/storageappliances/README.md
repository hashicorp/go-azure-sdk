
## `github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/storageappliances` Documentation

The `storageappliances` SDK allows for interaction with Azure Resource Manager `networkcloud` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/storageappliances"
```


### Client Initialization

```go
client := storageappliances.NewStorageAppliancesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `StorageAppliancesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := storageappliances.NewStorageApplianceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageApplianceName")

payload := storageappliances.StorageAppliance{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload, storageappliances.DefaultCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `StorageAppliancesClient.Delete`

```go
ctx := context.TODO()
id := storageappliances.NewStorageApplianceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageApplianceName")

if err := client.DeleteThenPoll(ctx, id, storageappliances.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `StorageAppliancesClient.DisableRemoteVendorManagement`

```go
ctx := context.TODO()
id := storageappliances.NewStorageApplianceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageApplianceName")

if err := client.DisableRemoteVendorManagementThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `StorageAppliancesClient.EnableRemoteVendorManagement`

```go
ctx := context.TODO()
id := storageappliances.NewStorageApplianceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageApplianceName")

payload := storageappliances.StorageApplianceEnableRemoteVendorManagementParameters{
	// ...
}


if err := client.EnableRemoteVendorManagementThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `StorageAppliancesClient.Get`

```go
ctx := context.TODO()
id := storageappliances.NewStorageApplianceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageApplianceName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StorageAppliancesClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id, storageappliances.DefaultListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id, storageappliances.DefaultListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StorageAppliancesClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id, storageappliances.DefaultListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id, storageappliances.DefaultListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StorageAppliancesClient.RunReadCommands`

```go
ctx := context.TODO()
id := storageappliances.NewStorageApplianceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageApplianceName")

payload := storageappliances.StorageApplianceRunReadCommandsParameters{
	// ...
}


if err := client.RunReadCommandsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `StorageAppliancesClient.Update`

```go
ctx := context.TODO()
id := storageappliances.NewStorageApplianceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageApplianceName")

payload := storageappliances.StorageAppliancePatchParameters{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload, storageappliances.DefaultUpdateOperationOptions()); err != nil {
	// handle the error
}
```
