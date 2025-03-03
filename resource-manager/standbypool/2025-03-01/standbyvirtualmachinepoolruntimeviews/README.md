
## `github.com/hashicorp/go-azure-sdk/resource-manager/standbypool/2025-03-01/standbyvirtualmachinepoolruntimeviews` Documentation

The `standbyvirtualmachinepoolruntimeviews` SDK allows for interaction with Azure Resource Manager `standbypool` (API Version `2025-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/standbypool/2025-03-01/standbyvirtualmachinepoolruntimeviews"
```


### Client Initialization

```go
client := standbyvirtualmachinepoolruntimeviews.NewStandbyVirtualMachinePoolRuntimeViewsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `StandbyVirtualMachinePoolRuntimeViewsClient.Get`

```go
ctx := context.TODO()
id := standbyvirtualmachinepoolruntimeviews.NewStandbyVirtualMachinePoolRuntimeViewID("12345678-1234-9876-4563-123456789012", "example-resource-group", "standbyVirtualMachinePoolName", "runtimeViewName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandbyVirtualMachinePoolRuntimeViewsClient.ListByStandbyPool`

```go
ctx := context.TODO()
id := standbyvirtualmachinepoolruntimeviews.NewStandbyVirtualMachinePoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "standbyVirtualMachinePoolName")

// alternatively `client.ListByStandbyPool(ctx, id)` can be used to do batched pagination
items, err := client.ListByStandbyPoolComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
