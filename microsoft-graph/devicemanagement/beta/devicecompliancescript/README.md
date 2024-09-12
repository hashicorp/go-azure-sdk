
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecompliancescript` Documentation

The `devicecompliancescript` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/devicecompliancescript"
```


### Client Initialization

```go
client := devicecompliancescript.NewDeviceComplianceScriptClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeviceComplianceScriptClient.AssignDeviceComplianceScript`

```go
ctx := context.TODO()
id := devicecompliancescript.NewDeviceManagementDeviceComplianceScriptID("deviceComplianceScriptIdValue")

payload := devicecompliancescript.AssignDeviceComplianceScriptRequest{
	// ...
}


read, err := client.AssignDeviceComplianceScript(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceComplianceScriptClient.CreateDeviceComplianceScript`

```go
ctx := context.TODO()

payload := devicecompliancescript.DeviceComplianceScript{
	// ...
}


read, err := client.CreateDeviceComplianceScript(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceComplianceScriptClient.DeleteDeviceComplianceScript`

```go
ctx := context.TODO()
id := devicecompliancescript.NewDeviceManagementDeviceComplianceScriptID("deviceComplianceScriptIdValue")

read, err := client.DeleteDeviceComplianceScript(ctx, id, devicecompliancescript.DefaultDeleteDeviceComplianceScriptOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceComplianceScriptClient.GetDeviceComplianceScript`

```go
ctx := context.TODO()
id := devicecompliancescript.NewDeviceManagementDeviceComplianceScriptID("deviceComplianceScriptIdValue")

read, err := client.GetDeviceComplianceScript(ctx, id, devicecompliancescript.DefaultGetDeviceComplianceScriptOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceComplianceScriptClient.GetDeviceComplianceScriptsCount`

```go
ctx := context.TODO()


read, err := client.GetDeviceComplianceScriptsCount(ctx, devicecompliancescript.DefaultGetDeviceComplianceScriptsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceComplianceScriptClient.ListDeviceComplianceScripts`

```go
ctx := context.TODO()


// alternatively `client.ListDeviceComplianceScripts(ctx, devicecompliancescript.DefaultListDeviceComplianceScriptsOperationOptions())` can be used to do batched pagination
items, err := client.ListDeviceComplianceScriptsComplete(ctx, devicecompliancescript.DefaultListDeviceComplianceScriptsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeviceComplianceScriptClient.UpdateDeviceComplianceScript`

```go
ctx := context.TODO()
id := devicecompliancescript.NewDeviceManagementDeviceComplianceScriptID("deviceComplianceScriptIdValue")

payload := devicecompliancescript.DeviceComplianceScript{
	// ...
}


read, err := client.UpdateDeviceComplianceScript(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
