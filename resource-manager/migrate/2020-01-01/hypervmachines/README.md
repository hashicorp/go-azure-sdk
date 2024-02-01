
## `github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-01-01/hypervmachines` Documentation

The `hypervmachines` SDK allows for interaction with the Azure Resource Manager Service `migrate` (API Version `2020-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-01-01/hypervmachines"
```


### Client Initialization

```go
client := hypervmachines.NewHyperVMachinesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `HyperVMachinesClient.GetAllMachinesInSite`

```go
ctx := context.TODO()
id := hypervmachines.NewHyperVSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hyperVSiteValue")

// alternatively `client.GetAllMachinesInSite(ctx, id, hypervmachines.DefaultGetAllMachinesInSiteOperationOptions())` can be used to do batched pagination
items, err := client.GetAllMachinesInSiteComplete(ctx, id, hypervmachines.DefaultGetAllMachinesInSiteOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `HyperVMachinesClient.GetMachine`

```go
ctx := context.TODO()
id := commonids.NewHyperVSiteMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hyperVSiteValue", "machineValue")

read, err := client.GetMachine(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
