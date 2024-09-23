
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/devicecompliancepolicy` Documentation

The `devicecompliancepolicy` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/devicecompliancepolicy"
```


### Client Initialization

```go
client := devicecompliancepolicy.NewDeviceCompliancePolicyClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeviceCompliancePolicyClient.AssignDeviceCompliancePolicies`

```go
ctx := context.TODO()
id := devicecompliancepolicy.NewDeviceManagementDeviceCompliancePolicyID("deviceCompliancePolicyId")

payload := devicecompliancepolicy.AssignDeviceCompliancePoliciesRequest{
	// ...
}


// alternatively `client.AssignDeviceCompliancePolicies(ctx, id, payload, devicecompliancepolicy.DefaultAssignDeviceCompliancePoliciesOperationOptions())` can be used to do batched pagination
items, err := client.AssignDeviceCompliancePoliciesComplete(ctx, id, payload, devicecompliancepolicy.DefaultAssignDeviceCompliancePoliciesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeviceCompliancePolicyClient.CreateDeviceCompliancePolicy`

```go
ctx := context.TODO()

payload := devicecompliancepolicy.DeviceCompliancePolicy{
	// ...
}


read, err := client.CreateDeviceCompliancePolicy(ctx, payload, devicecompliancepolicy.DefaultCreateDeviceCompliancePolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceCompliancePolicyClient.CreateDeviceCompliancePolicyScheduleActionsForRule`

```go
ctx := context.TODO()
id := devicecompliancepolicy.NewDeviceManagementDeviceCompliancePolicyID("deviceCompliancePolicyId")

payload := devicecompliancepolicy.CreateDeviceCompliancePolicyScheduleActionsForRuleRequest{
	// ...
}


read, err := client.CreateDeviceCompliancePolicyScheduleActionsForRule(ctx, id, payload, devicecompliancepolicy.DefaultCreateDeviceCompliancePolicyScheduleActionsForRuleOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceCompliancePolicyClient.DeleteDeviceCompliancePolicy`

```go
ctx := context.TODO()
id := devicecompliancepolicy.NewDeviceManagementDeviceCompliancePolicyID("deviceCompliancePolicyId")

read, err := client.DeleteDeviceCompliancePolicy(ctx, id, devicecompliancepolicy.DefaultDeleteDeviceCompliancePolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceCompliancePolicyClient.GetDeviceCompliancePoliciesCount`

```go
ctx := context.TODO()


read, err := client.GetDeviceCompliancePoliciesCount(ctx, devicecompliancepolicy.DefaultGetDeviceCompliancePoliciesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceCompliancePolicyClient.GetDeviceCompliancePolicy`

```go
ctx := context.TODO()
id := devicecompliancepolicy.NewDeviceManagementDeviceCompliancePolicyID("deviceCompliancePolicyId")

read, err := client.GetDeviceCompliancePolicy(ctx, id, devicecompliancepolicy.DefaultGetDeviceCompliancePolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceCompliancePolicyClient.ListDeviceCompliancePolicies`

```go
ctx := context.TODO()


// alternatively `client.ListDeviceCompliancePolicies(ctx, devicecompliancepolicy.DefaultListDeviceCompliancePoliciesOperationOptions())` can be used to do batched pagination
items, err := client.ListDeviceCompliancePoliciesComplete(ctx, devicecompliancepolicy.DefaultListDeviceCompliancePoliciesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeviceCompliancePolicyClient.UpdateDeviceCompliancePolicy`

```go
ctx := context.TODO()
id := devicecompliancepolicy.NewDeviceManagementDeviceCompliancePolicyID("deviceCompliancePolicyId")

payload := devicecompliancepolicy.DeviceCompliancePolicy{
	// ...
}


read, err := client.UpdateDeviceCompliancePolicy(ctx, id, payload, devicecompliancepolicy.DefaultUpdateDeviceCompliancePolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
