
## `github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-09-01/virtualmachines` Documentation

The `virtualmachines` SDK allows for interaction with Azure Resource Manager `vmware` (API Version `2023-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-09-01/virtualmachines"
```


### Client Initialization

```go
client := virtualmachines.NewVirtualMachinesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VirtualMachinesClient.Get`

```go
ctx := context.TODO()
id := virtualmachines.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "clusterName", "virtualMachineId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualMachinesClient.List`

```go
ctx := context.TODO()
id := virtualmachines.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "clusterName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VirtualMachinesClient.RestrictMovement`

```go
ctx := context.TODO()
id := virtualmachines.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "clusterName", "virtualMachineId")

payload := virtualmachines.VirtualMachineRestrictMovement{
	// ...
}


if err := client.RestrictMovementThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
