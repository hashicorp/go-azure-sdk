
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecustomattributeshellscript` Documentation

The `devicecustomattributeshellscript` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecustomattributeshellscript"
```


### Client Initialization

```go
client := devicecustomattributeshellscript.NewDeviceCustomAttributeShellScriptClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeviceCustomAttributeShellScriptClient.AssignDeviceManagementDeviceCustomAttributeShellScript`

```go
ctx := context.TODO()
id := devicecustomattributeshellscript.NewDeviceManagementDeviceCustomAttributeShellScriptID("deviceCustomAttributeShellScriptIdValue")

payload := devicecustomattributeshellscript.AssignDeviceManagementDeviceCustomAttributeShellScriptRequest{
	// ...
}


read, err := client.AssignDeviceManagementDeviceCustomAttributeShellScript(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceCustomAttributeShellScriptClient.CreateDeviceCustomAttributeShellScript`

```go
ctx := context.TODO()

payload := devicecustomattributeshellscript.DeviceCustomAttributeShellScript{
	// ...
}


read, err := client.CreateDeviceCustomAttributeShellScript(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceCustomAttributeShellScriptClient.DeleteDeviceCustomAttributeShellScript`

```go
ctx := context.TODO()
id := devicecustomattributeshellscript.NewDeviceManagementDeviceCustomAttributeShellScriptID("deviceCustomAttributeShellScriptIdValue")

read, err := client.DeleteDeviceCustomAttributeShellScript(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceCustomAttributeShellScriptClient.GetDeviceCustomAttributeShellScript`

```go
ctx := context.TODO()
id := devicecustomattributeshellscript.NewDeviceManagementDeviceCustomAttributeShellScriptID("deviceCustomAttributeShellScriptIdValue")

read, err := client.GetDeviceCustomAttributeShellScript(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceCustomAttributeShellScriptClient.GetDeviceCustomAttributeShellScriptCount`

```go
ctx := context.TODO()


read, err := client.GetDeviceCustomAttributeShellScriptCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceCustomAttributeShellScriptClient.ListDeviceCustomAttributeShellScripts`

```go
ctx := context.TODO()


// alternatively `client.ListDeviceCustomAttributeShellScripts(ctx)` can be used to do batched pagination
items, err := client.ListDeviceCustomAttributeShellScriptsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeviceCustomAttributeShellScriptClient.UpdateDeviceCustomAttributeShellScript`

```go
ctx := context.TODO()
id := devicecustomattributeshellscript.NewDeviceManagementDeviceCustomAttributeShellScriptID("deviceCustomAttributeShellScriptIdValue")

payload := devicecustomattributeshellscript.DeviceCustomAttributeShellScript{
	// ...
}


read, err := client.UpdateDeviceCustomAttributeShellScript(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
