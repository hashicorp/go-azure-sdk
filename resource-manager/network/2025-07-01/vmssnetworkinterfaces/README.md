
## `github.com/hashicorp/go-azure-sdk/resource-manager/network/2025-07-01/vmssnetworkinterfaces` Documentation

The `vmssnetworkinterfaces` SDK allows for interaction with Azure Resource Manager `network` (API Version `2025-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/network/2025-07-01/vmssnetworkinterfaces"
```


### Client Initialization

```go
client := vmssnetworkinterfaces.NewVMSSNetworkInterfacesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VMSSNetworkInterfacesClient.NetworkInterfacesGetVirtualMachineScaleSetNetworkInterface`

```go
ctx := context.TODO()
id := commonids.NewVirtualMachineScaleSetNetworkInterfaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetName", "virtualMachineName", "networkInterfaceName")

read, err := client.NetworkInterfacesGetVirtualMachineScaleSetNetworkInterface(ctx, id, vmssnetworkinterfaces.DefaultNetworkInterfacesGetVirtualMachineScaleSetNetworkInterfaceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VMSSNetworkInterfacesClient.NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfaces`

```go
ctx := context.TODO()
id := vmssnetworkinterfaces.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineScaleSetName", "virtualMachineName")

// alternatively `client.NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfaces(ctx, id)` can be used to do batched pagination
items, err := client.NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
