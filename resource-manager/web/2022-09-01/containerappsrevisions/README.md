
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/containerappsrevisions` Documentation

The `containerappsrevisions` SDK allows for interaction with Azure Resource Manager `web` (API Version `2022-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/containerappsrevisions"
```


### Client Initialization

```go
client := containerappsrevisions.NewContainerAppsRevisionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ContainerAppsRevisionsClient.ActivateRevision`

```go
ctx := context.TODO()
id := containerappsrevisions.NewRevisionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "name")

read, err := client.ActivateRevision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContainerAppsRevisionsClient.DeactivateRevision`

```go
ctx := context.TODO()
id := containerappsrevisions.NewRevisionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "name")

read, err := client.DeactivateRevision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContainerAppsRevisionsClient.GetRevision`

```go
ctx := context.TODO()
id := containerappsrevisions.NewRevisionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "name")

read, err := client.GetRevision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContainerAppsRevisionsClient.ListRevisions`

```go
ctx := context.TODO()
id := containerappsrevisions.NewProviderContainerAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "name")

// alternatively `client.ListRevisions(ctx, id)` can be used to do batched pagination
items, err := client.ListRevisionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ContainerAppsRevisionsClient.RestartRevision`

```go
ctx := context.TODO()
id := containerappsrevisions.NewRevisionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "name")

read, err := client.RestartRevision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
