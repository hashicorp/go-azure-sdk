
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/configurationpolicy` Documentation

The `configurationpolicy` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/configurationpolicy"
```


### Client Initialization

```go
client := configurationpolicy.NewConfigurationPolicyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConfigurationPolicyClient.AssignDeviceManagementConfigurationPolicies`

```go
ctx := context.TODO()
id := configurationpolicy.NewDeviceManagementConfigurationPolicyID("deviceManagementConfigurationPolicyIdValue")

payload := configurationpolicy.AssignDeviceManagementConfigurationPoliciesRequest{
	// ...
}


// alternatively `client.AssignDeviceManagementConfigurationPolicies(ctx, id, payload)` can be used to do batched pagination
items, err := client.AssignDeviceManagementConfigurationPoliciesComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ConfigurationPolicyClient.CreateConfigurationPolicy`

```go
ctx := context.TODO()

payload := configurationpolicy.DeviceManagementConfigurationPolicy{
	// ...
}


read, err := client.CreateConfigurationPolicy(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConfigurationPolicyClient.CreateConfigurationPolicyCreateCopy`

```go
ctx := context.TODO()
id := configurationpolicy.NewDeviceManagementConfigurationPolicyID("deviceManagementConfigurationPolicyIdValue")

payload := configurationpolicy.CreateConfigurationPolicyCreateCopyRequest{
	// ...
}


read, err := client.CreateConfigurationPolicyCreateCopy(ctx, id, payload)
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
id := configurationpolicy.NewDeviceManagementConfigurationPolicyID("deviceManagementConfigurationPolicyIdValue")

payload := configurationpolicy.CreateConfigurationPolicyReorderRequest{
	// ...
}


read, err := client.CreateConfigurationPolicyReorder(ctx, id, payload)
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
id := configurationpolicy.NewDeviceManagementConfigurationPolicyID("deviceManagementConfigurationPolicyIdValue")

read, err := client.DeleteConfigurationPolicy(ctx, id)
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
id := configurationpolicy.NewDeviceManagementConfigurationPolicyID("deviceManagementConfigurationPolicyIdValue")

read, err := client.GetConfigurationPolicy(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConfigurationPolicyClient.GetConfigurationPolicyCount`

```go
ctx := context.TODO()


read, err := client.GetConfigurationPolicyCount(ctx)
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


// alternatively `client.ListConfigurationPolicies(ctx)` can be used to do batched pagination
items, err := client.ListConfigurationPoliciesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ConfigurationPolicyClient.UpdateConfigurationPolicy`

```go
ctx := context.TODO()
id := configurationpolicy.NewDeviceManagementConfigurationPolicyID("deviceManagementConfigurationPolicyIdValue")

payload := configurationpolicy.DeviceManagementConfigurationPolicy{
	// ...
}


read, err := client.UpdateConfigurationPolicy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
