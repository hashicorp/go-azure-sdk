
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/builds` Documentation

The `builds` SDK allows for interaction with the Azure Resource Manager Service `containerapps` (API Version `2024-02-02-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/builds"
```


### Client Initialization

```go
client := builds.NewBuildsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BuildsClient.BuildAuthTokenList`

```go
ctx := context.TODO()
id := builds.NewBuildID("12345678-1234-9876-4563-123456789012", "example-resource-group", "builderValue", "buildValue")

read, err := client.BuildAuthTokenList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BuildsClient.ByBuilderResourceList`

```go
ctx := context.TODO()
id := builds.NewBuilderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "builderValue")

// alternatively `client.ByBuilderResourceList(ctx, id)` can be used to do batched pagination
items, err := client.ByBuilderResourceListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BuildsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := builds.NewBuildID("12345678-1234-9876-4563-123456789012", "example-resource-group", "builderValue", "buildValue")

payload := builds.BuildResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BuildsClient.Delete`

```go
ctx := context.TODO()
id := builds.NewBuildID("12345678-1234-9876-4563-123456789012", "example-resource-group", "builderValue", "buildValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `BuildsClient.Get`

```go
ctx := context.TODO()
id := builds.NewBuildID("12345678-1234-9876-4563-123456789012", "example-resource-group", "builderValue", "buildValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
