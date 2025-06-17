
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointprovisioningpolicy` Documentation

The `virtualendpointprovisioningpolicy` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointprovisioningpolicy"
```


### Client Initialization

```go
client := virtualendpointprovisioningpolicy.NewVirtualEndpointProvisioningPolicyClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VirtualEndpointProvisioningPolicyClient.AssignVirtualEndpointProvisioningPolicy`

```go
ctx := context.TODO()
id := virtualendpointprovisioningpolicy.NewDeviceManagementVirtualEndpointProvisioningPolicyID("cloudPCProvisioningPolicyId")

payload := virtualendpointprovisioningpolicy.AssignVirtualEndpointProvisioningPolicyRequest{
	// ...
}


read, err := client.AssignVirtualEndpointProvisioningPolicy(ctx, id, payload, virtualendpointprovisioningpolicy.DefaultAssignVirtualEndpointProvisioningPolicyOperationOptions())
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


read, err := client.CreateVirtualEndpointProvisioningPolicy(ctx, payload, virtualendpointprovisioningpolicy.DefaultCreateVirtualEndpointProvisioningPolicyOperationOptions())
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
id := virtualendpointprovisioningpolicy.NewDeviceManagementVirtualEndpointProvisioningPolicyID("cloudPCProvisioningPolicyId")

payload := virtualendpointprovisioningpolicy.CreateVirtualEndpointProvisioningPolicyApplyRequest{
	// ...
}


read, err := client.CreateVirtualEndpointProvisioningPolicyApply(ctx, id, payload, virtualendpointprovisioningpolicy.DefaultCreateVirtualEndpointProvisioningPolicyApplyOperationOptions())
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


read, err := client.CreateVirtualEndpointProvisioningPolicyApplyConfig(ctx, payload, virtualendpointprovisioningpolicy.DefaultCreateVirtualEndpointProvisioningPolicyApplyConfigOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointProvisioningPolicyClient.CreateVirtualEndpointProvisioningPolicySchedulePolicyApplyTask`

```go
ctx := context.TODO()
id := virtualendpointprovisioningpolicy.NewDeviceManagementVirtualEndpointProvisioningPolicyID("cloudPCProvisioningPolicyId")

payload := virtualendpointprovisioningpolicy.CreateVirtualEndpointProvisioningPolicySchedulePolicyApplyTaskRequest{
	// ...
}


read, err := client.CreateVirtualEndpointProvisioningPolicySchedulePolicyApplyTask(ctx, id, payload, virtualendpointprovisioningpolicy.DefaultCreateVirtualEndpointProvisioningPolicySchedulePolicyApplyTaskOperationOptions())
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
id := virtualendpointprovisioningpolicy.NewDeviceManagementVirtualEndpointProvisioningPolicyID("cloudPCProvisioningPolicyId")

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
id := virtualendpointprovisioningpolicy.NewDeviceManagementVirtualEndpointProvisioningPolicyID("cloudPCProvisioningPolicyId")

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
id := virtualendpointprovisioningpolicy.NewDeviceManagementVirtualEndpointProvisioningPolicyID("cloudPCProvisioningPolicyId")

payload := virtualendpointprovisioningpolicy.CloudPCProvisioningPolicy{
	// ...
}


read, err := client.UpdateVirtualEndpointProvisioningPolicy(ctx, id, payload, virtualendpointprovisioningpolicy.DefaultUpdateVirtualEndpointProvisioningPolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
