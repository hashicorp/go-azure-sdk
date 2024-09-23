
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/configurationpolicy` Documentation

The `configurationpolicy` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/configurationpolicy"
```


### Client Initialization

```go
client := configurationpolicy.NewConfigurationPolicyClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConfigurationPolicyClient.AssignConfigurationPolicies`

```go
ctx := context.TODO()
id := configurationpolicy.NewDeviceManagementConfigurationPolicyID("deviceManagementConfigurationPolicyId")

payload := configurationpolicy.AssignConfigurationPoliciesRequest{
	// ...
}


// alternatively `client.AssignConfigurationPolicies(ctx, id, payload, configurationpolicy.DefaultAssignConfigurationPoliciesOperationOptions())` can be used to do batched pagination
items, err := client.AssignConfigurationPoliciesComplete(ctx, id, payload, configurationpolicy.DefaultAssignConfigurationPoliciesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ConfigurationPolicyClient.ClearConfigurationPolicyEnrollmentTimeDeviceMembershipTarget`

```go
ctx := context.TODO()
id := configurationpolicy.NewDeviceManagementConfigurationPolicyID("deviceManagementConfigurationPolicyId")

read, err := client.ClearConfigurationPolicyEnrollmentTimeDeviceMembershipTarget(ctx, id, configurationpolicy.DefaultClearConfigurationPolicyEnrollmentTimeDeviceMembershipTargetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConfigurationPolicyClient.CreateConfigurationPolicy`

```go
ctx := context.TODO()

payload := configurationpolicy.DeviceManagementConfigurationPolicy{
	// ...
}


read, err := client.CreateConfigurationPolicy(ctx, payload, configurationpolicy.DefaultCreateConfigurationPolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConfigurationPolicyClient.CreateConfigurationPolicyCopy`

```go
ctx := context.TODO()
id := configurationpolicy.NewDeviceManagementConfigurationPolicyID("deviceManagementConfigurationPolicyId")

payload := configurationpolicy.CreateConfigurationPolicyCopyRequest{
	// ...
}


read, err := client.CreateConfigurationPolicyCopy(ctx, id, payload, configurationpolicy.DefaultCreateConfigurationPolicyCopyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConfigurationPolicyClient.CreateConfigurationPolicyReorder`

```go
ctx := context.TODO()
id := configurationpolicy.NewDeviceManagementConfigurationPolicyID("deviceManagementConfigurationPolicyId")

payload := configurationpolicy.CreateConfigurationPolicyReorderRequest{
	// ...
}


read, err := client.CreateConfigurationPolicyReorder(ctx, id, payload, configurationpolicy.DefaultCreateConfigurationPolicyReorderOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConfigurationPolicyClient.CreateConfigurationPolicyRetrieveEnrollmentTimeDeviceMembershipTarget`

```go
ctx := context.TODO()
id := configurationpolicy.NewDeviceManagementConfigurationPolicyID("deviceManagementConfigurationPolicyId")

read, err := client.CreateConfigurationPolicyRetrieveEnrollmentTimeDeviceMembershipTarget(ctx, id, configurationpolicy.DefaultCreateConfigurationPolicyRetrieveEnrollmentTimeDeviceMembershipTargetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConfigurationPolicyClient.DeleteConfigurationPolicy`

```go
ctx := context.TODO()
id := configurationpolicy.NewDeviceManagementConfigurationPolicyID("deviceManagementConfigurationPolicyId")

read, err := client.DeleteConfigurationPolicy(ctx, id, configurationpolicy.DefaultDeleteConfigurationPolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConfigurationPolicyClient.GetConfigurationPoliciesCount`

```go
ctx := context.TODO()


read, err := client.GetConfigurationPoliciesCount(ctx, configurationpolicy.DefaultGetConfigurationPoliciesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConfigurationPolicyClient.GetConfigurationPolicy`

```go
ctx := context.TODO()
id := configurationpolicy.NewDeviceManagementConfigurationPolicyID("deviceManagementConfigurationPolicyId")

read, err := client.GetConfigurationPolicy(ctx, id, configurationpolicy.DefaultGetConfigurationPolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConfigurationPolicyClient.ListConfigurationPolicies`

```go
ctx := context.TODO()


// alternatively `client.ListConfigurationPolicies(ctx, configurationpolicy.DefaultListConfigurationPoliciesOperationOptions())` can be used to do batched pagination
items, err := client.ListConfigurationPoliciesComplete(ctx, configurationpolicy.DefaultListConfigurationPoliciesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ConfigurationPolicyClient.SetConfigurationPolicyEnrollmentTimeDeviceMembershipTarget`

```go
ctx := context.TODO()
id := configurationpolicy.NewDeviceManagementConfigurationPolicyID("deviceManagementConfigurationPolicyId")

payload := configurationpolicy.SetConfigurationPolicyEnrollmentTimeDeviceMembershipTargetRequest{
	// ...
}


read, err := client.SetConfigurationPolicyEnrollmentTimeDeviceMembershipTarget(ctx, id, payload, configurationpolicy.DefaultSetConfigurationPolicyEnrollmentTimeDeviceMembershipTargetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConfigurationPolicyClient.UpdateConfigurationPolicy`

```go
ctx := context.TODO()
id := configurationpolicy.NewDeviceManagementConfigurationPolicyID("deviceManagementConfigurationPolicyId")

payload := configurationpolicy.DeviceManagementConfigurationPolicy{
	// ...
}


read, err := client.UpdateConfigurationPolicy(ctx, id, payload, configurationpolicy.DefaultUpdateConfigurationPolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
