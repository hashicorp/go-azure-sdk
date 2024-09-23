
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/compliancepolicy` Documentation

The `compliancepolicy` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/compliancepolicy"
```


### Client Initialization

```go
client := compliancepolicy.NewCompliancePolicyClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CompliancePolicyClient.AssignCompliancePolicies`

```go
ctx := context.TODO()
id := compliancepolicy.NewDeviceManagementCompliancePolicyID("deviceManagementCompliancePolicyId")

payload := compliancepolicy.AssignCompliancePoliciesRequest{
	// ...
}


// alternatively `client.AssignCompliancePolicies(ctx, id, payload, compliancepolicy.DefaultAssignCompliancePoliciesOperationOptions())` can be used to do batched pagination
items, err := client.AssignCompliancePoliciesComplete(ctx, id, payload, compliancepolicy.DefaultAssignCompliancePoliciesOperationOptions())
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


read, err := client.CreateCompliancePolicy(ctx, payload, compliancepolicy.DefaultCreateCompliancePolicyOperationOptions())
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
id := compliancepolicy.NewDeviceManagementCompliancePolicyID("deviceManagementCompliancePolicyId")

read, err := client.DeleteCompliancePolicy(ctx, id, compliancepolicy.DefaultDeleteCompliancePolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CompliancePolicyClient.GetCompliancePoliciesCount`

```go
ctx := context.TODO()


read, err := client.GetCompliancePoliciesCount(ctx, compliancepolicy.DefaultGetCompliancePoliciesCountOperationOptions())
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
id := compliancepolicy.NewDeviceManagementCompliancePolicyID("deviceManagementCompliancePolicyId")

read, err := client.GetCompliancePolicy(ctx, id, compliancepolicy.DefaultGetCompliancePolicyOperationOptions())
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


// alternatively `client.ListCompliancePolicies(ctx, compliancepolicy.DefaultListCompliancePoliciesOperationOptions())` can be used to do batched pagination
items, err := client.ListCompliancePoliciesComplete(ctx, compliancepolicy.DefaultListCompliancePoliciesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CompliancePolicyClient.SetCompliancePolicyScheduledActions`

```go
ctx := context.TODO()
id := compliancepolicy.NewDeviceManagementCompliancePolicyID("deviceManagementCompliancePolicyId")

payload := compliancepolicy.SetCompliancePolicyScheduledActionsRequest{
	// ...
}


// alternatively `client.SetCompliancePolicyScheduledActions(ctx, id, payload, compliancepolicy.DefaultSetCompliancePolicyScheduledActionsOperationOptions())` can be used to do batched pagination
items, err := client.SetCompliancePolicyScheduledActionsComplete(ctx, id, payload, compliancepolicy.DefaultSetCompliancePolicyScheduledActionsOperationOptions())
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
id := compliancepolicy.NewDeviceManagementCompliancePolicyID("deviceManagementCompliancePolicyId")

payload := compliancepolicy.DeviceManagementCompliancePolicy{
	// ...
}


read, err := client.UpdateCompliancePolicy(ctx, id, payload, compliancepolicy.DefaultUpdateCompliancePolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
