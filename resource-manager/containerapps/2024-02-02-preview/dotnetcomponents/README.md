
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/dotnetcomponents` Documentation

The `dotnetcomponents` SDK allows for interaction with Azure Resource Manager `containerapps` (API Version `2024-02-02-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/dotnetcomponents"
```


### Client Initialization

```go
client := dotnetcomponents.NewDotNetComponentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DotNetComponentsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := dotnetcomponents.NewDotNetComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedEnvironmentName", "dotNetComponentName")

payload := dotnetcomponents.DotNetComponent{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DotNetComponentsClient.Delete`

```go
ctx := context.TODO()
id := dotnetcomponents.NewDotNetComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedEnvironmentName", "dotNetComponentName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DotNetComponentsClient.Get`

```go
ctx := context.TODO()
id := dotnetcomponents.NewDotNetComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedEnvironmentName", "dotNetComponentName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DotNetComponentsClient.List`

```go
ctx := context.TODO()
id := dotnetcomponents.NewManagedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedEnvironmentName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DotNetComponentsClient.Update`

```go
ctx := context.TODO()
id := dotnetcomponents.NewDotNetComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedEnvironmentName", "dotNetComponentName")

payload := dotnetcomponents.DotNetComponent{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
