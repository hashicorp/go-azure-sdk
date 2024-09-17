
## `github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-09-01/workloadnetworkvirtualmachines` Documentation

The `workloadnetworkvirtualmachines` SDK allows for interaction with Azure Resource Manager `vmware` (API Version `2023-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-09-01/workloadnetworkvirtualmachines"
```


### Client Initialization

```go
client := workloadnetworkvirtualmachines.NewWorkloadNetworkVirtualMachinesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkloadNetworkVirtualMachinesClient.WorkloadNetworksGetVirtualMachine`

```go
ctx := context.TODO()
id := workloadnetworkvirtualmachines.NewDefaultVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "virtualMachineIdValue")

read, err := client.WorkloadNetworksGetVirtualMachine(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadNetworkVirtualMachinesClient.WorkloadNetworksListVirtualMachines`

```go
ctx := context.TODO()
id := workloadnetworkvirtualmachines.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue")

// alternatively `client.WorkloadNetworksListVirtualMachines(ctx, id)` can be used to do batched pagination
items, err := client.WorkloadNetworksListVirtualMachinesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
