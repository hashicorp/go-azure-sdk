
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


### Example Usage: `VirtualEndpointProvisioningPolicyClient.AssignVirtualEndpointProvisioningPolicy`

```go
ctx := context.TODO()
id := virtualendpointprovisioningpolicy.NewDeviceManagementVirtualEndpointProvisioningPolicyID("cloudPCProvisioningPolicyIdValue")

payload := virtualendpointprovisioningpolicy.AssignVirtualEndpointProvisioningPolicyRequest{
	// ...
}


read, err := client.AssignVirtualEndpointProvisioningPolicy(ctx, id, payload)
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

read, err := client.DeleteVirtualEndpointProvisioningPolicy(ctx, id, virtualendpointprovisioningpolicy.DefaultDeleteVirtualEndpointProvisioningPolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointProvisioningPolicyClient.GetVirtualEndpointProvisioningPoliciesCount`

```go
ctx := context.TODO()


read, err := client.GetVirtualEndpointProvisioningPoliciesCount(ctx, virtualendpointprovisioningpolicy.DefaultGetVirtualEndpointProvisioningPoliciesCountOperationOptions())
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

read, err := client.GetVirtualEndpointProvisioningPolicy(ctx, id, virtualendpointprovisioningpolicy.DefaultGetVirtualEndpointProvisioningPolicyOperationOptions())
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


// alternatively `client.ListVirtualEndpointProvisioningPolicies(ctx, virtualendpointprovisioningpolicy.DefaultListVirtualEndpointProvisioningPoliciesOperationOptions())` can be used to do batched pagination
items, err := client.ListVirtualEndpointProvisioningPoliciesComplete(ctx, virtualendpointprovisioningpolicy.DefaultListVirtualEndpointProvisioningPoliciesOperationOptions())
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
