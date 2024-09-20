
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/javacomponents` Documentation

The `javacomponents` SDK allows for interaction with Azure Resource Manager `containerapps` (API Version `2024-02-02-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/javacomponents"
```


### Client Initialization

```go
client := javacomponents.NewJavaComponentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JavaComponentsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := javacomponents.NewJavaComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "environmentName", "name")

payload := javacomponents.JavaComponent{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `JavaComponentsClient.Delete`

```go
ctx := context.TODO()
id := javacomponents.NewJavaComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "environmentName", "name")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `JavaComponentsClient.Get`

```go
ctx := context.TODO()
id := javacomponents.NewJavaComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "environmentName", "name")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JavaComponentsClient.List`

```go
ctx := context.TODO()
id := javacomponents.NewManagedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "environmentName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JavaComponentsClient.Update`

```go
ctx := context.TODO()
id := javacomponents.NewJavaComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "environmentName", "name")

payload := javacomponents.JavaComponent{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
