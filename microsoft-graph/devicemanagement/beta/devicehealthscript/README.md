
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicehealthscript` Documentation

The `devicehealthscript` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicehealthscript"
```


### Client Initialization

```go
client := devicehealthscript.NewDeviceHealthScriptClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeviceHealthScriptClient.AssignDeviceManagementDeviceHealthScript`

```go
ctx := context.TODO()
id := devicehealthscript.NewDeviceManagementDeviceHealthScriptID("deviceHealthScriptIdValue")

payload := devicehealthscript.AssignDeviceManagementDeviceHealthScriptRequest{
	// ...
}


read, err := client.AssignDeviceManagementDeviceHealthScript(ctx, id, payload)
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


read, err := client.CreateDeviceHealthScript(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceHealthScriptClient.CreateDeviceHealthScriptEnableGlobalScript`

```go
ctx := context.TODO()


read, err := client.CreateDeviceHealthScriptEnableGlobalScript(ctx)
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
id := devicehealthscript.NewDeviceManagementDeviceHealthScriptID("deviceHealthScriptIdValue")

read, err := client.DeleteDeviceHealthScript(ctx, id)
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
id := devicehealthscript.NewDeviceManagementDeviceHealthScriptID("deviceHealthScriptIdValue")

read, err := client.GetDeviceHealthScript(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceHealthScriptClient.GetDeviceHealthScriptCount`

```go
ctx := context.TODO()


read, err := client.GetDeviceHealthScriptCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceHealthScriptClient.GetDeviceManagementDeviceHealthScriptGlobalScriptHighestAvailableVersion`

```go
ctx := context.TODO()
id := devicehealthscript.NewDeviceManagementDeviceHealthScriptID("deviceHealthScriptIdValue")

read, err := client.GetDeviceManagementDeviceHealthScriptGlobalScriptHighestAvailableVersion(ctx, id)
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


// alternatively `client.ListDeviceHealthScripts(ctx)` can be used to do batched pagination
items, err := client.ListDeviceHealthScriptsComplete(ctx)
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
id := devicehealthscript.NewDeviceManagementDeviceHealthScriptID("deviceHealthScriptIdValue")

payload := devicehealthscript.DeviceHealthScript{
	// ...
}


read, err := client.UpdateDeviceHealthScript(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceHealthScriptClient.UpdateDeviceManagementDeviceHealthScriptGlobalScript`

```go
ctx := context.TODO()
id := devicehealthscript.NewDeviceManagementDeviceHealthScriptID("deviceHealthScriptIdValue")

payload := devicehealthscript.UpdateDeviceManagementDeviceHealthScriptGlobalScriptRequest{
	// ...
}


read, err := client.UpdateDeviceManagementDeviceHealthScriptGlobalScript(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
