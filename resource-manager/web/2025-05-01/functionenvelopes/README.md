
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/functionenvelopes` Documentation

The `functionenvelopes` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/functionenvelopes"
```


### Client Initialization

```go
client := functionenvelopes.NewFunctionEnvelopesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FunctionEnvelopesClient.WebAppsCreateFunction`

```go
ctx := context.TODO()
id := functionenvelopes.NewFunctionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "functionName")

payload := functionenvelopes.FunctionEnvelope{
	// ...
}


if err := client.WebAppsCreateFunctionThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `FunctionEnvelopesClient.WebAppsCreateOrUpdateFunctionSecret`

```go
ctx := context.TODO()
id := functionenvelopes.NewKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "functionName", "keyName")

payload := functionenvelopes.KeyInfo{
	// ...
}


read, err := client.WebAppsCreateOrUpdateFunctionSecret(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FunctionEnvelopesClient.WebAppsDeleteFunction`

```go
ctx := context.TODO()
id := functionenvelopes.NewFunctionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "functionName")

read, err := client.WebAppsDeleteFunction(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FunctionEnvelopesClient.WebAppsDeleteFunctionSecret`

```go
ctx := context.TODO()
id := functionenvelopes.NewKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "functionName", "keyName")

read, err := client.WebAppsDeleteFunctionSecret(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FunctionEnvelopesClient.WebAppsGetFunction`

```go
ctx := context.TODO()
id := functionenvelopes.NewFunctionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "functionName")

read, err := client.WebAppsGetFunction(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FunctionEnvelopesClient.WebAppsListFunctionKeys`

```go
ctx := context.TODO()
id := functionenvelopes.NewFunctionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "functionName")

read, err := client.WebAppsListFunctionKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FunctionEnvelopesClient.WebAppsListFunctionSecrets`

```go
ctx := context.TODO()
id := functionenvelopes.NewFunctionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "functionName")

read, err := client.WebAppsListFunctionSecrets(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FunctionEnvelopesClient.WebAppsListFunctions`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.WebAppsListFunctions(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListFunctionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
