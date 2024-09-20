
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicemanagementscript` Documentation

The `devicemanagementscript` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicemanagementscript"
```


### Client Initialization

```go
client := devicemanagementscript.NewDeviceManagementScriptClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeviceManagementScriptClient.AssignDeviceManagementScript`

```go
ctx := context.TODO()
id := devicemanagementscript.NewDeviceManagementDeviceManagementScriptID("deviceManagementScriptId")

payload := devicemanagementscript.AssignDeviceManagementScriptRequest{
	// ...
}


read, err := client.AssignDeviceManagementScript(ctx, id, payload, devicemanagementscript.DefaultAssignDeviceManagementScriptOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementScriptClient.CreateDeviceManagementScript`

```go
ctx := context.TODO()

payload := devicemanagementscript.DeviceManagementScript{
	// ...
}


read, err := client.CreateDeviceManagementScript(ctx, payload, devicemanagementscript.DefaultCreateDeviceManagementScriptOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementScriptClient.DeleteDeviceManagementScript`

```go
ctx := context.TODO()
id := devicemanagementscript.NewDeviceManagementDeviceManagementScriptID("deviceManagementScriptId")

read, err := client.DeleteDeviceManagementScript(ctx, id, devicemanagementscript.DefaultDeleteDeviceManagementScriptOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementScriptClient.GetDeviceManagementScript`

```go
ctx := context.TODO()
id := devicemanagementscript.NewDeviceManagementDeviceManagementScriptID("deviceManagementScriptId")

read, err := client.GetDeviceManagementScript(ctx, id, devicemanagementscript.DefaultGetDeviceManagementScriptOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementScriptClient.GetDeviceManagementScriptsCount`

```go
ctx := context.TODO()


read, err := client.GetDeviceManagementScriptsCount(ctx, devicemanagementscript.DefaultGetDeviceManagementScriptsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementScriptClient.ListDeviceManagementScriptHasPayloadLinks`

```go
ctx := context.TODO()

payload := devicemanagementscript.ListDeviceManagementScriptHasPayloadLinksRequest{
	// ...
}


// alternatively `client.ListDeviceManagementScriptHasPayloadLinks(ctx, payload, devicemanagementscript.DefaultListDeviceManagementScriptHasPayloadLinksOperationOptions())` can be used to do batched pagination
items, err := client.ListDeviceManagementScriptHasPayloadLinksComplete(ctx, payload, devicemanagementscript.DefaultListDeviceManagementScriptHasPayloadLinksOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeviceManagementScriptClient.ListDeviceManagementScripts`

```go
ctx := context.TODO()


// alternatively `client.ListDeviceManagementScripts(ctx, devicemanagementscript.DefaultListDeviceManagementScriptsOperationOptions())` can be used to do batched pagination
items, err := client.ListDeviceManagementScriptsComplete(ctx, devicemanagementscript.DefaultListDeviceManagementScriptsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeviceManagementScriptClient.UpdateDeviceManagementScript`

```go
ctx := context.TODO()
id := devicemanagementscript.NewDeviceManagementDeviceManagementScriptID("deviceManagementScriptId")

payload := devicemanagementscript.DeviceManagementScript{
	// ...
}


read, err := client.UpdateDeviceManagementScript(ctx, id, payload, devicemanagementscript.DefaultUpdateDeviceManagementScriptOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
