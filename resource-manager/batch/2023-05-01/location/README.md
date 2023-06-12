
## `github.com/hashicorp/go-azure-sdk/resource-manager/batch/2023-05-01/location` Documentation

The `location` SDK allows for interaction with the Azure Resource Manager Service `batch` (API Version `2023-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/batch/2023-05-01/location"
```


### Client Initialization

```go
client := location.NewLocationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LocationClient.GetQuotas`

```go
ctx := context.TODO()
id := location.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

read, err := client.GetQuotas(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LocationClient.ListSupportedCloudServiceSkus`

```go
ctx := context.TODO()
id := location.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

// alternatively `client.ListSupportedCloudServiceSkus(ctx, id, location.DefaultListSupportedCloudServiceSkusOperationOptions())` can be used to do batched pagination
items, err := client.ListSupportedCloudServiceSkusComplete(ctx, id, location.DefaultListSupportedCloudServiceSkusOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `LocationClient.ListSupportedVirtualMachineSkus`

```go
ctx := context.TODO()
id := location.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

// alternatively `client.ListSupportedVirtualMachineSkus(ctx, id, location.DefaultListSupportedVirtualMachineSkusOperationOptions())` can be used to do batched pagination
items, err := client.ListSupportedVirtualMachineSkusComplete(ctx, id, location.DefaultListSupportedVirtualMachineSkusOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
