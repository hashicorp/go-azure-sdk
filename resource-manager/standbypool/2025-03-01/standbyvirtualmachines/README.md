
## `github.com/hashicorp/go-azure-sdk/resource-manager/standbypool/2025-03-01/standbyvirtualmachines` Documentation

The `standbyvirtualmachines` SDK allows for interaction with Azure Resource Manager `standbypool` (API Version `2025-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/standbypool/2025-03-01/standbyvirtualmachines"
```


### Client Initialization

```go
client := standbyvirtualmachines.NewStandbyVirtualMachinesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `StandbyVirtualMachinesClient.Get`

```go
ctx := context.TODO()
id := standbyvirtualmachines.NewStandbyVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "standbyVirtualMachinePoolName", "standbyVirtualMachineName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandbyVirtualMachinesClient.ListByStandbyVirtualMachinePoolResource`

```go
ctx := context.TODO()
id := standbyvirtualmachines.NewStandbyVirtualMachinePoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "standbyVirtualMachinePoolName")

// alternatively `client.ListByStandbyVirtualMachinePoolResource(ctx, id)` can be used to do batched pagination
items, err := client.ListByStandbyVirtualMachinePoolResourceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
