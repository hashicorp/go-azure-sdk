
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/importedwindowsautopilotdeviceidentity` Documentation

The `importedwindowsautopilotdeviceidentity` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/importedwindowsautopilotdeviceidentity"
```


### Client Initialization

```go
client := importedwindowsautopilotdeviceidentity.NewImportedWindowsAutopilotDeviceIdentityClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ImportedWindowsAutopilotDeviceIdentityClient.CreateImportedWindowsAutopilotDeviceIdentity`

```go
ctx := context.TODO()

payload := importedwindowsautopilotdeviceidentity.ImportedWindowsAutopilotDeviceIdentity{
	// ...
}


read, err := client.CreateImportedWindowsAutopilotDeviceIdentity(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ImportedWindowsAutopilotDeviceIdentityClient.DeleteImportedWindowsAutopilotDeviceIdentity`

```go
ctx := context.TODO()
id := importedwindowsautopilotdeviceidentity.NewDeviceManagementImportedWindowsAutopilotDeviceIdentityID("importedWindowsAutopilotDeviceIdentityIdValue")

read, err := client.DeleteImportedWindowsAutopilotDeviceIdentity(ctx, id, importedwindowsautopilotdeviceidentity.DefaultDeleteImportedWindowsAutopilotDeviceIdentityOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ImportedWindowsAutopilotDeviceIdentityClient.GetImportedWindowsAutopilotDeviceIdentitiesCount`

```go
ctx := context.TODO()


read, err := client.GetImportedWindowsAutopilotDeviceIdentitiesCount(ctx, importedwindowsautopilotdeviceidentity.DefaultGetImportedWindowsAutopilotDeviceIdentitiesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ImportedWindowsAutopilotDeviceIdentityClient.GetImportedWindowsAutopilotDeviceIdentity`

```go
ctx := context.TODO()
id := importedwindowsautopilotdeviceidentity.NewDeviceManagementImportedWindowsAutopilotDeviceIdentityID("importedWindowsAutopilotDeviceIdentityIdValue")

read, err := client.GetImportedWindowsAutopilotDeviceIdentity(ctx, id, importedwindowsautopilotdeviceidentity.DefaultGetImportedWindowsAutopilotDeviceIdentityOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ImportedWindowsAutopilotDeviceIdentityClient.ListImportedWindowsAutopilotDeviceIdentities`

```go
ctx := context.TODO()


// alternatively `client.ListImportedWindowsAutopilotDeviceIdentities(ctx, importedwindowsautopilotdeviceidentity.DefaultListImportedWindowsAutopilotDeviceIdentitiesOperationOptions())` can be used to do batched pagination
items, err := client.ListImportedWindowsAutopilotDeviceIdentitiesComplete(ctx, importedwindowsautopilotdeviceidentity.DefaultListImportedWindowsAutopilotDeviceIdentitiesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ImportedWindowsAutopilotDeviceIdentityClient.ListImportedWindowsAutopilotDeviceIdentityImports`

```go
ctx := context.TODO()

payload := importedwindowsautopilotdeviceidentity.ListImportedWindowsAutopilotDeviceIdentityImportsRequest{
	// ...
}


// alternatively `client.ListImportedWindowsAutopilotDeviceIdentityImports(ctx, payload, importedwindowsautopilotdeviceidentity.DefaultListImportedWindowsAutopilotDeviceIdentityImportsOperationOptions())` can be used to do batched pagination
items, err := client.ListImportedWindowsAutopilotDeviceIdentityImportsComplete(ctx, payload, importedwindowsautopilotdeviceidentity.DefaultListImportedWindowsAutopilotDeviceIdentityImportsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ImportedWindowsAutopilotDeviceIdentityClient.UpdateImportedWindowsAutopilotDeviceIdentity`

```go
ctx := context.TODO()
id := importedwindowsautopilotdeviceidentity.NewDeviceManagementImportedWindowsAutopilotDeviceIdentityID("importedWindowsAutopilotDeviceIdentityIdValue")

payload := importedwindowsautopilotdeviceidentity.ImportedWindowsAutopilotDeviceIdentity{
	// ...
}


read, err := client.UpdateImportedWindowsAutopilotDeviceIdentity(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
