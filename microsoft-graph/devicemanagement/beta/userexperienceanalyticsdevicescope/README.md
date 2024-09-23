
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsdevicescope` Documentation

The `userexperienceanalyticsdevicescope` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/userexperienceanalyticsdevicescope"
```


### Client Initialization

```go
client := userexperienceanalyticsdevicescope.NewUserExperienceAnalyticsDeviceScopeClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserExperienceAnalyticsDeviceScopeClient.CreateUserExperienceAnalyticsDeviceScope`

```go
ctx := context.TODO()

payload := userexperienceanalyticsdevicescope.UserExperienceAnalyticsDeviceScope{
	// ...
}


read, err := client.CreateUserExperienceAnalyticsDeviceScope(ctx, payload, userexperienceanalyticsdevicescope.DefaultCreateUserExperienceAnalyticsDeviceScopeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserExperienceAnalyticsDeviceScopeClient.CreateUserExperienceAnalyticsDeviceScopeTriggerAction`

```go
ctx := context.TODO()
id := userexperienceanalyticsdevicescope.NewDeviceManagementUserExperienceAnalyticsDeviceScopeID("userExperienceAnalyticsDeviceScopeId")

payload := userexperienceanalyticsdevicescope.CreateUserExperienceAnalyticsDeviceScopeTriggerActionRequest{
	// ...
}


read, err := client.CreateUserExperienceAnalyticsDeviceScopeTriggerAction(ctx, id, payload, userexperienceanalyticsdevicescope.DefaultCreateUserExperienceAnalyticsDeviceScopeTriggerActionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserExperienceAnalyticsDeviceScopeClient.DeleteUserExperienceAnalyticsDeviceScope`

```go
ctx := context.TODO()
id := userexperienceanalyticsdevicescope.NewDeviceManagementUserExperienceAnalyticsDeviceScopeID("userExperienceAnalyticsDeviceScopeId")

read, err := client.DeleteUserExperienceAnalyticsDeviceScope(ctx, id, userexperienceanalyticsdevicescope.DefaultDeleteUserExperienceAnalyticsDeviceScopeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserExperienceAnalyticsDeviceScopeClient.GetUserExperienceAnalyticsDeviceScope`

```go
ctx := context.TODO()
id := userexperienceanalyticsdevicescope.NewDeviceManagementUserExperienceAnalyticsDeviceScopeID("userExperienceAnalyticsDeviceScopeId")

read, err := client.GetUserExperienceAnalyticsDeviceScope(ctx, id, userexperienceanalyticsdevicescope.DefaultGetUserExperienceAnalyticsDeviceScopeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserExperienceAnalyticsDeviceScopeClient.GetUserExperienceAnalyticsDeviceScopesCount`

```go
ctx := context.TODO()


read, err := client.GetUserExperienceAnalyticsDeviceScopesCount(ctx, userexperienceanalyticsdevicescope.DefaultGetUserExperienceAnalyticsDeviceScopesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserExperienceAnalyticsDeviceScopeClient.ListUserExperienceAnalyticsDeviceScopes`

```go
ctx := context.TODO()


// alternatively `client.ListUserExperienceAnalyticsDeviceScopes(ctx, userexperienceanalyticsdevicescope.DefaultListUserExperienceAnalyticsDeviceScopesOperationOptions())` can be used to do batched pagination
items, err := client.ListUserExperienceAnalyticsDeviceScopesComplete(ctx, userexperienceanalyticsdevicescope.DefaultListUserExperienceAnalyticsDeviceScopesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserExperienceAnalyticsDeviceScopeClient.UpdateUserExperienceAnalyticsDeviceScope`

```go
ctx := context.TODO()
id := userexperienceanalyticsdevicescope.NewDeviceManagementUserExperienceAnalyticsDeviceScopeID("userExperienceAnalyticsDeviceScopeId")

payload := userexperienceanalyticsdevicescope.UserExperienceAnalyticsDeviceScope{
	// ...
}


read, err := client.UpdateUserExperienceAnalyticsDeviceScope(ctx, id, payload, userexperienceanalyticsdevicescope.DefaultUpdateUserExperienceAnalyticsDeviceScopeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
