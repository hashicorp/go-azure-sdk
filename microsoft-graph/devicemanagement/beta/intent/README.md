
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/intent` Documentation

The `intent` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/intent"
```


### Client Initialization

```go
client := intent.NewIntentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IntentClient.AssignIntent`

```go
ctx := context.TODO()
id := intent.NewDeviceManagementIntentID("deviceManagementIntentIdValue")

payload := intent.AssignIntentRequest{
	// ...
}


read, err := client.AssignIntent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IntentClient.CreateIntent`

```go
ctx := context.TODO()

payload := intent.DeviceManagementIntent{
	// ...
}


read, err := client.CreateIntent(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IntentClient.CreateIntentCopy`

```go
ctx := context.TODO()
id := intent.NewDeviceManagementIntentID("deviceManagementIntentIdValue")

payload := intent.CreateIntentCopyRequest{
	// ...
}


read, err := client.CreateIntentCopy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IntentClient.CreateIntentMigrateToTemplate`

```go
ctx := context.TODO()
id := intent.NewDeviceManagementIntentID("deviceManagementIntentIdValue")

payload := intent.CreateIntentMigrateToTemplateRequest{
	// ...
}


read, err := client.CreateIntentMigrateToTemplate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IntentClient.DeleteIntent`

```go
ctx := context.TODO()
id := intent.NewDeviceManagementIntentID("deviceManagementIntentIdValue")

read, err := client.DeleteIntent(ctx, id, intent.DefaultDeleteIntentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IntentClient.GetIntent`

```go
ctx := context.TODO()
id := intent.NewDeviceManagementIntentID("deviceManagementIntentIdValue")

read, err := client.GetIntent(ctx, id, intent.DefaultGetIntentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IntentClient.GetIntentsCount`

```go
ctx := context.TODO()


read, err := client.GetIntentsCount(ctx, intent.DefaultGetIntentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IntentClient.ListIntents`

```go
ctx := context.TODO()


// alternatively `client.ListIntents(ctx, intent.DefaultListIntentsOperationOptions())` can be used to do batched pagination
items, err := client.ListIntentsComplete(ctx, intent.DefaultListIntentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IntentClient.UpdateIntent`

```go
ctx := context.TODO()
id := intent.NewDeviceManagementIntentID("deviceManagementIntentIdValue")

payload := intent.DeviceManagementIntent{
	// ...
}


read, err := client.UpdateIntent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IntentClient.UpdateIntentSetting`

```go
ctx := context.TODO()
id := intent.NewDeviceManagementIntentID("deviceManagementIntentIdValue")

payload := intent.UpdateIntentSettingRequest{
	// ...
}


read, err := client.UpdateIntentSetting(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
