
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


### Example Usage: `DeviceManagementScriptClient.AssignDeviceManagementDeviceManagementScript`

```go
ctx := context.TODO()
id := devicemanagementscript.NewDeviceManagementDeviceManagementScriptID("deviceManagementScriptIdValue")

payload := devicemanagementscript.AssignDeviceManagementDeviceManagementScriptRequest{
	// ...
}


read, err := client.AssignDeviceManagementDeviceManagementScript(ctx, id, payload)
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


read, err := client.CreateDeviceManagementScript(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementScriptClient.CreateDeviceManagementScriptHasPayloadLink`

```go
ctx := context.TODO()

payload := devicemanagementscript.CreateDeviceManagementScriptHasPayloadLinkRequest{
	// ...
}


read, err := client.CreateDeviceManagementScriptHasPayloadLink(ctx, payload)
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
id := devicemanagementscript.NewDeviceManagementDeviceManagementScriptID("deviceManagementScriptIdValue")

read, err := client.DeleteDeviceManagementScript(ctx, id)
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
id := devicemanagementscript.NewDeviceManagementDeviceManagementScriptID("deviceManagementScriptIdValue")

read, err := client.GetDeviceManagementScript(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementScriptClient.GetDeviceManagementScriptCount`

```go
ctx := context.TODO()


read, err := client.GetDeviceManagementScriptCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceManagementScriptClient.ListDeviceManagementScripts`

```go
ctx := context.TODO()


// alternatively `client.ListDeviceManagementScripts(ctx)` can be used to do batched pagination
items, err := client.ListDeviceManagementScriptsComplete(ctx)
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
id := devicemanagementscript.NewDeviceManagementDeviceManagementScriptID("deviceManagementScriptIdValue")

payload := devicemanagementscript.DeviceManagementScript{
	// ...
}


read, err := client.UpdateDeviceManagementScript(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
