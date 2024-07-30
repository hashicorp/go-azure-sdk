
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceshellscript` Documentation

The `deviceshellscript` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deviceshellscript"
```


### Client Initialization

```go
client := deviceshellscript.NewDeviceShellScriptClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeviceShellScriptClient.AssignDeviceManagementDeviceShellScript`

```go
ctx := context.TODO()
id := deviceshellscript.NewDeviceManagementDeviceShellScriptID("deviceShellScriptIdValue")

payload := deviceshellscript.AssignDeviceManagementDeviceShellScriptRequest{
	// ...
}


read, err := client.AssignDeviceManagementDeviceShellScript(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceShellScriptClient.CreateDeviceShellScript`

```go
ctx := context.TODO()

payload := deviceshellscript.DeviceShellScript{
	// ...
}


read, err := client.CreateDeviceShellScript(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceShellScriptClient.DeleteDeviceShellScript`

```go
ctx := context.TODO()
id := deviceshellscript.NewDeviceManagementDeviceShellScriptID("deviceShellScriptIdValue")

read, err := client.DeleteDeviceShellScript(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceShellScriptClient.GetDeviceShellScript`

```go
ctx := context.TODO()
id := deviceshellscript.NewDeviceManagementDeviceShellScriptID("deviceShellScriptIdValue")

read, err := client.GetDeviceShellScript(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceShellScriptClient.GetDeviceShellScriptCount`

```go
ctx := context.TODO()


read, err := client.GetDeviceShellScriptCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceShellScriptClient.ListDeviceShellScripts`

```go
ctx := context.TODO()


// alternatively `client.ListDeviceShellScripts(ctx)` can be used to do batched pagination
items, err := client.ListDeviceShellScriptsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeviceShellScriptClient.UpdateDeviceShellScript`

```go
ctx := context.TODO()
id := deviceshellscript.NewDeviceManagementDeviceShellScriptID("deviceShellScriptIdValue")

payload := deviceshellscript.DeviceShellScript{
	// ...
}


read, err := client.UpdateDeviceShellScript(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
