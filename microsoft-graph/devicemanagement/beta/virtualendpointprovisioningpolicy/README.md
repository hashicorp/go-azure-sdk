
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointprovisioningpolicy` Documentation

The `virtualendpointprovisioningpolicy` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointprovisioningpolicy"
```


### Client Initialization

```go
client := virtualendpointprovisioningpolicy.NewVirtualEndpointProvisioningPolicyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VirtualEndpointProvisioningPolicyClient.AssignDeviceManagementVirtualEndpointProvisioningPolicy`

```go
ctx := context.TODO()
id := virtualendpointprovisioningpolicy.NewDeviceManagementVirtualEndpointProvisioningPolicyID("cloudPCProvisioningPolicyIdValue")

payload := virtualendpointprovisioningpolicy.AssignDeviceManagementVirtualEndpointProvisioningPolicyRequest{
	// ...
}


read, err := client.AssignDeviceManagementVirtualEndpointProvisioningPolicy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointProvisioningPolicyClient.CreateVirtualEndpointProvisioningPolicy`

```go
ctx := context.TODO()

payload := virtualendpointprovisioningpolicy.CloudPCProvisioningPolicy{
	// ...
}


read, err := client.CreateVirtualEndpointProvisioningPolicy(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointProvisioningPolicyClient.CreateVirtualEndpointProvisioningPolicyApply`

```go
ctx := context.TODO()
id := virtualendpointprovisioningpolicy.NewDeviceManagementVirtualEndpointProvisioningPolicyID("cloudPCProvisioningPolicyIdValue")

payload := virtualendpointprovisioningpolicy.CreateVirtualEndpointProvisioningPolicyApplyRequest{
	// ...
}


read, err := client.CreateVirtualEndpointProvisioningPolicyApply(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointProvisioningPolicyClient.CreateVirtualEndpointProvisioningPolicyApplyConfig`

```go
ctx := context.TODO()

payload := virtualendpointprovisioningpolicy.CreateVirtualEndpointProvisioningPolicyApplyConfigRequest{
	// ...
}


read, err := client.CreateVirtualEndpointProvisioningPolicyApplyConfig(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointProvisioningPolicyClient.DeleteVirtualEndpointProvisioningPolicy`

```go
ctx := context.TODO()
id := virtualendpointprovisioningpolicy.NewDeviceManagementVirtualEndpointProvisioningPolicyID("cloudPCProvisioningPolicyIdValue")

read, err := client.DeleteVirtualEndpointProvisioningPolicy(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointProvisioningPolicyClient.GetVirtualEndpointProvisioningPolicy`

```go
ctx := context.TODO()
id := virtualendpointprovisioningpolicy.NewDeviceManagementVirtualEndpointProvisioningPolicyID("cloudPCProvisioningPolicyIdValue")

read, err := client.GetVirtualEndpointProvisioningPolicy(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointProvisioningPolicyClient.GetVirtualEndpointProvisioningPolicyCount`

```go
ctx := context.TODO()


read, err := client.GetVirtualEndpointProvisioningPolicyCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointProvisioningPolicyClient.ListVirtualEndpointProvisioningPolicies`

```go
ctx := context.TODO()


// alternatively `client.ListVirtualEndpointProvisioningPolicies(ctx)` can be used to do batched pagination
items, err := client.ListVirtualEndpointProvisioningPoliciesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VirtualEndpointProvisioningPolicyClient.UpdateVirtualEndpointProvisioningPolicy`

```go
ctx := context.TODO()
id := virtualendpointprovisioningpolicy.NewDeviceManagementVirtualEndpointProvisioningPolicyID("cloudPCProvisioningPolicyIdValue")

payload := virtualendpointprovisioningpolicy.CloudPCProvisioningPolicy{
	// ...
}


read, err := client.UpdateVirtualEndpointProvisioningPolicy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
