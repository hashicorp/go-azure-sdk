
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsqualityupdatepolicy` Documentation

The `windowsqualityupdatepolicy` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/windowsqualityupdatepolicy"
```


### Client Initialization

```go
client := windowsqualityupdatepolicy.NewWindowsQualityUpdatePolicyClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WindowsQualityUpdatePolicyClient.AssignWindowsQualityUpdatePolicy`

```go
ctx := context.TODO()
id := windowsqualityupdatepolicy.NewDeviceManagementWindowsQualityUpdatePolicyID("windowsQualityUpdatePolicyId")

payload := windowsqualityupdatepolicy.AssignWindowsQualityUpdatePolicyRequest{
	// ...
}


read, err := client.AssignWindowsQualityUpdatePolicy(ctx, id, payload, windowsqualityupdatepolicy.DefaultAssignWindowsQualityUpdatePolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsQualityUpdatePolicyClient.CreateWindowsQualityUpdatePolicy`

```go
ctx := context.TODO()

payload := windowsqualityupdatepolicy.WindowsQualityUpdatePolicy{
	// ...
}


read, err := client.CreateWindowsQualityUpdatePolicy(ctx, payload, windowsqualityupdatepolicy.DefaultCreateWindowsQualityUpdatePolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsQualityUpdatePolicyClient.DeleteWindowsQualityUpdatePolicy`

```go
ctx := context.TODO()
id := windowsqualityupdatepolicy.NewDeviceManagementWindowsQualityUpdatePolicyID("windowsQualityUpdatePolicyId")

read, err := client.DeleteWindowsQualityUpdatePolicy(ctx, id, windowsqualityupdatepolicy.DefaultDeleteWindowsQualityUpdatePolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsQualityUpdatePolicyClient.GetWindowsQualityUpdatePoliciesCount`

```go
ctx := context.TODO()


read, err := client.GetWindowsQualityUpdatePoliciesCount(ctx, windowsqualityupdatepolicy.DefaultGetWindowsQualityUpdatePoliciesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsQualityUpdatePolicyClient.GetWindowsQualityUpdatePolicy`

```go
ctx := context.TODO()
id := windowsqualityupdatepolicy.NewDeviceManagementWindowsQualityUpdatePolicyID("windowsQualityUpdatePolicyId")

read, err := client.GetWindowsQualityUpdatePolicy(ctx, id, windowsqualityupdatepolicy.DefaultGetWindowsQualityUpdatePolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WindowsQualityUpdatePolicyClient.ListWindowsQualityUpdatePolicies`

```go
ctx := context.TODO()


// alternatively `client.ListWindowsQualityUpdatePolicies(ctx, windowsqualityupdatepolicy.DefaultListWindowsQualityUpdatePoliciesOperationOptions())` can be used to do batched pagination
items, err := client.ListWindowsQualityUpdatePoliciesComplete(ctx, windowsqualityupdatepolicy.DefaultListWindowsQualityUpdatePoliciesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WindowsQualityUpdatePolicyClient.UpdateWindowsQualityUpdatePolicy`

```go
ctx := context.TODO()
id := windowsqualityupdatepolicy.NewDeviceManagementWindowsQualityUpdatePolicyID("windowsQualityUpdatePolicyId")

payload := windowsqualityupdatepolicy.WindowsQualityUpdatePolicy{
	// ...
}


read, err := client.UpdateWindowsQualityUpdatePolicy(ctx, id, payload, windowsqualityupdatepolicy.DefaultUpdateWindowsQualityUpdatePolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
