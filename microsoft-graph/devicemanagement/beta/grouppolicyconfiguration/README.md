
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicyconfiguration` Documentation

The `grouppolicyconfiguration` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicyconfiguration"
```


### Client Initialization

```go
client := grouppolicyconfiguration.NewGroupPolicyConfigurationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupPolicyConfigurationClient.AssignGroupPolicyConfigurations`

```go
ctx := context.TODO()
id := grouppolicyconfiguration.NewDeviceManagementGroupPolicyConfigurationID("groupPolicyConfigurationIdValue")

payload := grouppolicyconfiguration.AssignGroupPolicyConfigurationsRequest{
	// ...
}


// alternatively `client.AssignGroupPolicyConfigurations(ctx, id, payload, grouppolicyconfiguration.DefaultAssignGroupPolicyConfigurationsOperationOptions())` can be used to do batched pagination
items, err := client.AssignGroupPolicyConfigurationsComplete(ctx, id, payload, grouppolicyconfiguration.DefaultAssignGroupPolicyConfigurationsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupPolicyConfigurationClient.CreateGroupPolicyConfiguration`

```go
ctx := context.TODO()

payload := grouppolicyconfiguration.GroupPolicyConfiguration{
	// ...
}


read, err := client.CreateGroupPolicyConfiguration(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyConfigurationClient.DeleteGroupPolicyConfiguration`

```go
ctx := context.TODO()
id := grouppolicyconfiguration.NewDeviceManagementGroupPolicyConfigurationID("groupPolicyConfigurationIdValue")

read, err := client.DeleteGroupPolicyConfiguration(ctx, id, grouppolicyconfiguration.DefaultDeleteGroupPolicyConfigurationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyConfigurationClient.GetGroupPolicyConfiguration`

```go
ctx := context.TODO()
id := grouppolicyconfiguration.NewDeviceManagementGroupPolicyConfigurationID("groupPolicyConfigurationIdValue")

read, err := client.GetGroupPolicyConfiguration(ctx, id, grouppolicyconfiguration.DefaultGetGroupPolicyConfigurationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyConfigurationClient.GetGroupPolicyConfigurationsCount`

```go
ctx := context.TODO()


read, err := client.GetGroupPolicyConfigurationsCount(ctx, grouppolicyconfiguration.DefaultGetGroupPolicyConfigurationsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyConfigurationClient.ListGroupPolicyConfigurations`

```go
ctx := context.TODO()


// alternatively `client.ListGroupPolicyConfigurations(ctx, grouppolicyconfiguration.DefaultListGroupPolicyConfigurationsOperationOptions())` can be used to do batched pagination
items, err := client.ListGroupPolicyConfigurationsComplete(ctx, grouppolicyconfiguration.DefaultListGroupPolicyConfigurationsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupPolicyConfigurationClient.UpdateGroupPolicyConfiguration`

```go
ctx := context.TODO()
id := grouppolicyconfiguration.NewDeviceManagementGroupPolicyConfigurationID("groupPolicyConfigurationIdValue")

payload := grouppolicyconfiguration.GroupPolicyConfiguration{
	// ...
}


read, err := client.UpdateGroupPolicyConfiguration(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyConfigurationClient.UpdateGroupPolicyConfigurationDefinitionValue`

```go
ctx := context.TODO()
id := grouppolicyconfiguration.NewDeviceManagementGroupPolicyConfigurationID("groupPolicyConfigurationIdValue")

payload := grouppolicyconfiguration.UpdateGroupPolicyConfigurationDefinitionValueRequest{
	// ...
}


read, err := client.UpdateGroupPolicyConfigurationDefinitionValue(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
