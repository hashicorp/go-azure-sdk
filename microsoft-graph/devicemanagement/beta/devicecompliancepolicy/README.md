
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecompliancepolicy` Documentation

The `devicecompliancepolicy` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecompliancepolicy"
```


### Client Initialization

```go
client := devicecompliancepolicy.NewDeviceCompliancePolicyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeviceCompliancePolicyClient.AssignDeviceCompliancePolicies`

```go
ctx := context.TODO()
id := devicecompliancepolicy.NewDeviceManagementDeviceCompliancePolicyID("deviceCompliancePolicyIdValue")

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


read, err := client.CreateDeviceCompliancePolicy(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceCompliancePolicyClient.CreateDeviceCompliancePolicyRefreshReportSummarization`

```go
ctx := context.TODO()


read, err := client.CreateDeviceCompliancePolicyRefreshReportSummarization(ctx)
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
id := devicecompliancepolicy.NewDeviceManagementDeviceCompliancePolicyID("deviceCompliancePolicyIdValue")

payload := devicecompliancepolicy.CreateDeviceCompliancePolicyScheduleActionsForRuleRequest{
	// ...
}


read, err := client.CreateDeviceCompliancePolicyScheduleActionsForRule(ctx, id, payload)
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
id := devicecompliancepolicy.NewDeviceManagementDeviceCompliancePolicyID("deviceCompliancePolicyIdValue")

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


### Example Usage: `DeviceCompliancePolicyClient.GetDeviceCompliancePoliciesNoncompliantDevicesToRetire`

```go
ctx := context.TODO()

payload := devicecompliancepolicy.GetDeviceCompliancePoliciesNoncompliantDevicesToRetireRequest{
	// ...
}


read, err := client.GetDeviceCompliancePoliciesNoncompliantDevicesToRetire(ctx, payload)
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
id := devicecompliancepolicy.NewDeviceManagementDeviceCompliancePolicyID("deviceCompliancePolicyIdValue")

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


### Example Usage: `DeviceCompliancePolicyClient.ListDeviceCompliancePolicyHasPayloadLinks`

```go
ctx := context.TODO()

payload := devicecompliancepolicy.ListDeviceCompliancePolicyHasPayloadLinksRequest{
	// ...
}


// alternatively `client.ListDeviceCompliancePolicyHasPayloadLinks(ctx, payload, devicecompliancepolicy.DefaultListDeviceCompliancePolicyHasPayloadLinksOperationOptions())` can be used to do batched pagination
items, err := client.ListDeviceCompliancePolicyHasPayloadLinksComplete(ctx, payload, devicecompliancepolicy.DefaultListDeviceCompliancePolicyHasPayloadLinksOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeviceCompliancePolicyClient.SetDeviceCompliancePoliciesScheduledRetireState`

```go
ctx := context.TODO()

payload := devicecompliancepolicy.SetDeviceCompliancePoliciesScheduledRetireStateRequest{
	// ...
}


read, err := client.SetDeviceCompliancePoliciesScheduledRetireState(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceCompliancePolicyClient.UpdateDeviceCompliancePolicy`

```go
ctx := context.TODO()
id := devicecompliancepolicy.NewDeviceManagementDeviceCompliancePolicyID("deviceCompliancePolicyIdValue")

payload := devicecompliancepolicy.DeviceCompliancePolicy{
	// ...
}


read, err := client.UpdateDeviceCompliancePolicy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceCompliancePolicyClient.ValidateDeviceCompliancePoliciesComplianceScript`

```go
ctx := context.TODO()

payload := devicecompliancepolicy.ValidateDeviceCompliancePoliciesComplianceScriptRequest{
	// ...
}


read, err := client.ValidateDeviceCompliancePoliciesComplianceScript(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
