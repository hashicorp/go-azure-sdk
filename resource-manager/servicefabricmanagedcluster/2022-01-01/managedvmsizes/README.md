
## `github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2022-01-01/managedvmsizes` Documentation

The `managedvmsizes` SDK allows for interaction with the Azure Resource Manager Service `servicefabricmanagedcluster` (API Version `2022-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/servicefabricmanagedcluster/2022-01-01/managedvmsizes"
```


### Client Initialization

```go
client := managedvmsizes.NewManagedVMSizesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedVMSizesClient.ManagedUnsupportedVMSizesGet`

```go
ctx := context.TODO()
id := managedvmsizes.NewManagedUnsupportedVMSizeID("12345678-1234-9876-4563-123456789012", "locationValue", "managedUnsupportedVMSizeValue")

read, err := client.ManagedUnsupportedVMSizesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedVMSizesClient.ManagedUnsupportedVMSizesList`

```go
ctx := context.TODO()
id := managedvmsizes.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

// alternatively `client.ManagedUnsupportedVMSizesList(ctx, id)` can be used to do batched pagination
items, err := client.ManagedUnsupportedVMSizesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
