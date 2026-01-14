
## `github.com/hashicorp/go-azure-sdk/resource-manager/batch/2024-07-01/openapis` Documentation

The `openapis` SDK allows for interaction with Azure Resource Manager `batch` (API Version `2024-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/batch/2024-07-01/openapis"
```


### Client Initialization

```go
client := openapis.NewOpenapisClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OpenapisClient.LocationCheckNameAvailability`

```go
ctx := context.TODO()
id := openapis.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := openapis.CheckNameAvailabilityParameters{
	// ...
}


read, err := client.LocationCheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.LocationGetQuotas`

```go
ctx := context.TODO()
id := openapis.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

read, err := client.LocationGetQuotas(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.LocationListSupportedVirtualMachineSkus`

```go
ctx := context.TODO()
id := openapis.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.LocationListSupportedVirtualMachineSkus(ctx, id, openapis.DefaultLocationListSupportedVirtualMachineSkusOperationOptions())` can be used to do batched pagination
items, err := client.LocationListSupportedVirtualMachineSkusComplete(ctx, id, openapis.DefaultLocationListSupportedVirtualMachineSkusOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
