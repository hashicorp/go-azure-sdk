
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicehealthscript` Documentation

The `devicehealthscript` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicehealthscript"
```


### Client Initialization

```go
client := devicehealthscript.NewDeviceHealthScriptClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeviceHealthScriptClient.AssignDeviceHealthScript`

```go
ctx := context.TODO()
id := devicehealthscript.NewDeviceManagementDeviceHealthScriptID("deviceHealthScriptId")

payload := devicehealthscript.AssignDeviceHealthScriptRequest{
	// ...
}


read, err := client.AssignDeviceHealthScript(ctx, id, payload, devicehealthscript.DefaultAssignDeviceHealthScriptOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceHealthScriptClient.CreateDeviceHealthScript`

```go
ctx := context.TODO()

payload := devicehealthscript.DeviceHealthScript{
	// ...
}


read, err := client.CreateDeviceHealthScript(ctx, payload, devicehealthscript.DefaultCreateDeviceHealthScriptOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceHealthScriptClient.DeleteDeviceHealthScript`

```go
ctx := context.TODO()
id := devicehealthscript.NewDeviceManagementDeviceHealthScriptID("deviceHealthScriptId")

read, err := client.DeleteDeviceHealthScript(ctx, id, devicehealthscript.DefaultDeleteDeviceHealthScriptOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceHealthScriptClient.EnableDeviceHealthScriptsGlobalScripts`

```go
ctx := context.TODO()


read, err := client.EnableDeviceHealthScriptsGlobalScripts(ctx, devicehealthscript.DefaultEnableDeviceHealthScriptsGlobalScriptsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceHealthScriptClient.GetDeviceHealthScript`

```go
ctx := context.TODO()
id := devicehealthscript.NewDeviceManagementDeviceHealthScriptID("deviceHealthScriptId")

read, err := client.GetDeviceHealthScript(ctx, id, devicehealthscript.DefaultGetDeviceHealthScriptOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceHealthScriptClient.GetDeviceHealthScriptGlobalScriptHighestAvailableVersion`

```go
ctx := context.TODO()
id := devicehealthscript.NewDeviceManagementDeviceHealthScriptID("deviceHealthScriptId")

read, err := client.GetDeviceHealthScriptGlobalScriptHighestAvailableVersion(ctx, id, devicehealthscript.DefaultGetDeviceHealthScriptGlobalScriptHighestAvailableVersionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceHealthScriptClient.GetDeviceHealthScriptsCount`

```go
ctx := context.TODO()


read, err := client.GetDeviceHealthScriptsCount(ctx, devicehealthscript.DefaultGetDeviceHealthScriptsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceHealthScriptClient.ListDeviceHealthScripts`

```go
ctx := context.TODO()


// alternatively `client.ListDeviceHealthScripts(ctx, devicehealthscript.DefaultListDeviceHealthScriptsOperationOptions())` can be used to do batched pagination
items, err := client.ListDeviceHealthScriptsComplete(ctx, devicehealthscript.DefaultListDeviceHealthScriptsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeviceHealthScriptClient.UpdateDeviceHealthScript`

```go
ctx := context.TODO()
id := devicehealthscript.NewDeviceManagementDeviceHealthScriptID("deviceHealthScriptId")

payload := devicehealthscript.DeviceHealthScript{
	// ...
}


read, err := client.UpdateDeviceHealthScript(ctx, id, payload, devicehealthscript.DefaultUpdateDeviceHealthScriptOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceHealthScriptClient.UpdateDeviceHealthScriptGlobalScript`

```go
ctx := context.TODO()
id := devicehealthscript.NewDeviceManagementDeviceHealthScriptID("deviceHealthScriptId")

payload := devicehealthscript.UpdateDeviceHealthScriptGlobalScriptRequest{
	// ...
}


read, err := client.UpdateDeviceHealthScriptGlobalScript(ctx, id, payload, devicehealthscript.DefaultUpdateDeviceHealthScriptGlobalScriptOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
