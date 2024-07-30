
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/compliancepolicy` Documentation

The `compliancepolicy` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/compliancepolicy"
```


### Client Initialization

```go
client := compliancepolicy.NewCompliancePolicyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CompliancePolicyClient.AssignDeviceManagementCompliancePolicies`

```go
ctx := context.TODO()
id := compliancepolicy.NewDeviceManagementCompliancePolicyID("deviceManagementCompliancePolicyIdValue")

payload := compliancepolicy.AssignDeviceManagementCompliancePoliciesRequest{
	// ...
}


// alternatively `client.AssignDeviceManagementCompliancePolicies(ctx, id, payload)` can be used to do batched pagination
items, err := client.AssignDeviceManagementCompliancePoliciesComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CompliancePolicyClient.CreateCompliancePolicy`

```go
ctx := context.TODO()

payload := compliancepolicy.DeviceManagementCompliancePolicy{
	// ...
}


read, err := client.CreateCompliancePolicy(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CompliancePolicyClient.DeleteCompliancePolicy`

```go
ctx := context.TODO()
id := compliancepolicy.NewDeviceManagementCompliancePolicyID("deviceManagementCompliancePolicyIdValue")

read, err := client.DeleteCompliancePolicy(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CompliancePolicyClient.GetCompliancePolicy`

```go
ctx := context.TODO()
id := compliancepolicy.NewDeviceManagementCompliancePolicyID("deviceManagementCompliancePolicyIdValue")

read, err := client.GetCompliancePolicy(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CompliancePolicyClient.GetCompliancePolicyCount`

```go
ctx := context.TODO()


read, err := client.GetCompliancePolicyCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CompliancePolicyClient.ListCompliancePolicies`

```go
ctx := context.TODO()


// alternatively `client.ListCompliancePolicies(ctx)` can be used to do batched pagination
items, err := client.ListCompliancePoliciesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CompliancePolicyClient.SetDeviceManagementCompliancePolicyScheduledActions`

```go
ctx := context.TODO()
id := compliancepolicy.NewDeviceManagementCompliancePolicyID("deviceManagementCompliancePolicyIdValue")

payload := compliancepolicy.SetDeviceManagementCompliancePolicyScheduledActionsRequest{
	// ...
}


// alternatively `client.SetDeviceManagementCompliancePolicyScheduledActions(ctx, id, payload)` can be used to do batched pagination
items, err := client.SetDeviceManagementCompliancePolicyScheduledActionsComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CompliancePolicyClient.UpdateCompliancePolicy`

```go
ctx := context.TODO()
id := compliancepolicy.NewDeviceManagementCompliancePolicyID("deviceManagementCompliancePolicyIdValue")

payload := compliancepolicy.DeviceManagementCompliancePolicy{
	// ...
}


read, err := client.UpdateCompliancePolicy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
