
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicemanagementscript` Documentation

The `devicemanagementscript` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicemanagementscript"
```


### Client Initialization

```go
client := devicemanagementscript.NewDeviceManagementScriptClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeviceManagementScriptClient.AssignScript`

```go
ctx := context.TODO()
id := devicemanagementscript.NewDeviceManagementDeviceManagementScriptID("deviceManagementScriptIdValue")

payload := devicemanagementscript.AssignScriptRequest{
	// ...
}


read, err := client.AssignScript(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementScriptClient.CreateScript`

```go
ctx := context.TODO()

payload := devicemanagementscript.DeviceManagementScript{
	// ...
}


read, err := client.CreateScript(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementScriptClient.DeleteScript`

```go
ctx := context.TODO()
id := devicemanagementscript.NewDeviceManagementDeviceManagementScriptID("deviceManagementScriptIdValue")

read, err := client.DeleteScript(ctx, id, devicemanagementscript.DefaultDeleteScriptOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementScriptClient.GetScript`

```go
ctx := context.TODO()
id := devicemanagementscript.NewDeviceManagementDeviceManagementScriptID("deviceManagementScriptIdValue")

read, err := client.GetScript(ctx, id, devicemanagementscript.DefaultGetScriptOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementScriptClient.GetScriptsCount`

```go
ctx := context.TODO()


read, err := client.GetScriptsCount(ctx, devicemanagementscript.DefaultGetScriptsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementScriptClient.ListScriptHasPayloadLinks`

```go
ctx := context.TODO()

payload := devicemanagementscript.ListScriptHasPayloadLinksRequest{
	// ...
}


// alternatively `client.ListScriptHasPayloadLinks(ctx, payload, devicemanagementscript.DefaultListScriptHasPayloadLinksOperationOptions())` can be used to do batched pagination
items, err := client.ListScriptHasPayloadLinksComplete(ctx, payload, devicemanagementscript.DefaultListScriptHasPayloadLinksOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeviceManagementScriptClient.ListScripts`

```go
ctx := context.TODO()


// alternatively `client.ListScripts(ctx, devicemanagementscript.DefaultListScriptsOperationOptions())` can be used to do batched pagination
items, err := client.ListScriptsComplete(ctx, devicemanagementscript.DefaultListScriptsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeviceManagementScriptClient.UpdateScript`

```go
ctx := context.TODO()
id := devicemanagementscript.NewDeviceManagementDeviceManagementScriptID("deviceManagementScriptIdValue")

payload := devicemanagementscript.DeviceManagementScript{
	// ...
}


read, err := client.UpdateScript(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
