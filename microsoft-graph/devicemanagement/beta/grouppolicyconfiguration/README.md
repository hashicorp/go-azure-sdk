
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


### Example Usage: `GroupPolicyConfigurationClient.AssignDeviceManagementGroupPolicyConfigurations`

```go
ctx := context.TODO()
id := grouppolicyconfiguration.NewDeviceManagementGroupPolicyConfigurationID("groupPolicyConfigurationIdValue")

payload := grouppolicyconfiguration.AssignDeviceManagementGroupPolicyConfigurationsRequest{
	// ...
}


// alternatively `client.AssignDeviceManagementGroupPolicyConfigurations(ctx, id, payload)` can be used to do batched pagination
items, err := client.AssignDeviceManagementGroupPolicyConfigurationsComplete(ctx, id, payload)
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

read, err := client.DeleteGroupPolicyConfiguration(ctx, id)
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

read, err := client.GetGroupPolicyConfiguration(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyConfigurationClient.GetGroupPolicyConfigurationCount`

```go
ctx := context.TODO()


read, err := client.GetGroupPolicyConfigurationCount(ctx)
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


// alternatively `client.ListGroupPolicyConfigurations(ctx)` can be used to do batched pagination
items, err := client.ListGroupPolicyConfigurationsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupPolicyConfigurationClient.UpdateDeviceManagementGroupPolicyConfigurationDefinitionValue`

```go
ctx := context.TODO()
id := grouppolicyconfiguration.NewDeviceManagementGroupPolicyConfigurationID("groupPolicyConfigurationIdValue")

payload := grouppolicyconfiguration.UpdateDeviceManagementGroupPolicyConfigurationDefinitionValueRequest{
	// ...
}


read, err := client.UpdateDeviceManagementGroupPolicyConfigurationDefinitionValue(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
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
