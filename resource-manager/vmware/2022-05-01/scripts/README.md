
## `github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2022-05-01/scripts` Documentation

The `scripts` SDK allows for interaction with Azure Resource Manager `vmware` (API Version `2022-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2022-05-01/scripts"
```


### Client Initialization

```go
client := scripts.NewScriptsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ScriptsClient.ScriptCmdletsGet`

```go
ctx := context.TODO()
id := scripts.NewScriptCmdletID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "scriptPackageName", "scriptCmdletName")

read, err := client.ScriptCmdletsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ScriptsClient.ScriptCmdletsList`

```go
ctx := context.TODO()
id := scripts.NewScriptPackageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "scriptPackageName")

// alternatively `client.ScriptCmdletsList(ctx, id)` can be used to do batched pagination
items, err := client.ScriptCmdletsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ScriptsClient.ScriptExecutionsCreateOrUpdate`

```go
ctx := context.TODO()
id := scripts.NewScriptExecutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "scriptExecutionName")

payload := scripts.ScriptExecution{
	// ...
}


if err := client.ScriptExecutionsCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ScriptsClient.ScriptExecutionsDelete`

```go
ctx := context.TODO()
id := scripts.NewScriptExecutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "scriptExecutionName")

if err := client.ScriptExecutionsDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ScriptsClient.ScriptExecutionsGet`

```go
ctx := context.TODO()
id := scripts.NewScriptExecutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "scriptExecutionName")

read, err := client.ScriptExecutionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ScriptsClient.ScriptExecutionsGetExecutionLogs`

```go
ctx := context.TODO()
id := scripts.NewScriptExecutionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "scriptExecutionName")
var payload []ScriptOutputStreamType

read, err := client.ScriptExecutionsGetExecutionLogs(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ScriptsClient.ScriptExecutionsList`

```go
ctx := context.TODO()
id := scripts.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName")

// alternatively `client.ScriptExecutionsList(ctx, id)` can be used to do batched pagination
items, err := client.ScriptExecutionsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ScriptsClient.ScriptPackagesGet`

```go
ctx := context.TODO()
id := scripts.NewScriptPackageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "scriptPackageName")

read, err := client.ScriptPackagesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ScriptsClient.ScriptPackagesList`

```go
ctx := context.TODO()
id := scripts.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName")

// alternatively `client.ScriptPackagesList(ctx, id)` can be used to do batched pagination
items, err := client.ScriptPackagesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
