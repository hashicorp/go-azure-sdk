
## `github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2022-09-04/machinepools` Documentation

The `machinepools` SDK allows for interaction with the Azure Resource Manager Service `redhatopenshift` (API Version `2022-09-04`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2022-09-04/machinepools"
```


### Client Initialization

```go
client := machinepools.NewMachinePoolsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MachinePoolsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := machinepools.NewMachinePoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "openShiftClusterValue", "machinePoolValue")

payload := machinepools.MachinePool{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MachinePoolsClient.Delete`

```go
ctx := context.TODO()
id := machinepools.NewMachinePoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "openShiftClusterValue", "machinePoolValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MachinePoolsClient.Get`

```go
ctx := context.TODO()
id := machinepools.NewMachinePoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "openShiftClusterValue", "machinePoolValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MachinePoolsClient.List`

```go
ctx := context.TODO()
id := machinepools.NewOpenShiftClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "openShiftClusterValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MachinePoolsClient.Update`

```go
ctx := context.TODO()
id := machinepools.NewMachinePoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "openShiftClusterValue", "machinePoolValue")

payload := machinepools.MachinePoolUpdate{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
