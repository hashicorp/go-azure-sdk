
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/containerappspatches` Documentation

The `containerappspatches` SDK allows for interaction with Azure Resource Manager `containerapps` (API Version `2024-02-02-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/containerappspatches"
```


### Client Initialization

```go
client := containerappspatches.NewContainerAppsPatchesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ContainerAppsPatchesClient.Apply`

```go
ctx := context.TODO()
id := containerappspatches.NewPatchID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "patchName")

if err := client.ApplyThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ContainerAppsPatchesClient.Delete`

```go
ctx := context.TODO()
id := containerappspatches.NewPatchID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "patchName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ContainerAppsPatchesClient.Get`

```go
ctx := context.TODO()
id := containerappspatches.NewPatchID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "patchName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContainerAppsPatchesClient.ListByContainerApp`

```go
ctx := context.TODO()
id := containerappspatches.NewContainerAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName")

// alternatively `client.ListByContainerApp(ctx, id, containerappspatches.DefaultListByContainerAppOperationOptions())` can be used to do batched pagination
items, err := client.ListByContainerAppComplete(ctx, id, containerappspatches.DefaultListByContainerAppOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ContainerAppsPatchesClient.SkipConfigure`

```go
ctx := context.TODO()
id := containerappspatches.NewPatchID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "patchName")

payload := containerappspatches.PatchSkipConfig{
	// ...
}


if err := client.SkipConfigureThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
