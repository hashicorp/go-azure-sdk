
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/functionenvelopeoperationgroup` Documentation

The `functionenvelopeoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/functionenvelopeoperationgroup"
```


### Client Initialization

```go
client := functionenvelopeoperationgroup.NewFunctionEnvelopeOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FunctionEnvelopeOperationGroupClient.WebAppsCreateInstanceFunctionSlot`

```go
ctx := context.TODO()
id := functionenvelopeoperationgroup.NewSlotFunctionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "functionName")

payload := functionenvelopeoperationgroup.FunctionEnvelope{
	// ...
}


if err := client.WebAppsCreateInstanceFunctionSlotThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `FunctionEnvelopeOperationGroupClient.WebAppsCreateOrUpdateFunctionSecretSlot`

```go
ctx := context.TODO()
id := functionenvelopeoperationgroup.NewFunctionKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "functionName", "keyName")

payload := functionenvelopeoperationgroup.KeyInfo{
	// ...
}


read, err := client.WebAppsCreateOrUpdateFunctionSecretSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FunctionEnvelopeOperationGroupClient.WebAppsDeleteFunctionSecretSlot`

```go
ctx := context.TODO()
id := functionenvelopeoperationgroup.NewFunctionKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "functionName", "keyName")

read, err := client.WebAppsDeleteFunctionSecretSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FunctionEnvelopeOperationGroupClient.WebAppsDeleteInstanceFunctionSlot`

```go
ctx := context.TODO()
id := functionenvelopeoperationgroup.NewSlotFunctionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "functionName")

read, err := client.WebAppsDeleteInstanceFunctionSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FunctionEnvelopeOperationGroupClient.WebAppsGetInstanceFunctionSlot`

```go
ctx := context.TODO()
id := functionenvelopeoperationgroup.NewSlotFunctionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "functionName")

read, err := client.WebAppsGetInstanceFunctionSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FunctionEnvelopeOperationGroupClient.WebAppsListFunctionKeysSlot`

```go
ctx := context.TODO()
id := functionenvelopeoperationgroup.NewSlotFunctionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "functionName")

read, err := client.WebAppsListFunctionKeysSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FunctionEnvelopeOperationGroupClient.WebAppsListFunctionSecretsSlot`

```go
ctx := context.TODO()
id := functionenvelopeoperationgroup.NewSlotFunctionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "functionName")

read, err := client.WebAppsListFunctionSecretsSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FunctionEnvelopeOperationGroupClient.WebAppsListInstanceFunctionsSlot`

```go
ctx := context.TODO()
id := functionenvelopeoperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.WebAppsListInstanceFunctionsSlot(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListInstanceFunctionsSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
