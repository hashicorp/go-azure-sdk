
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/containerappsbuilds` Documentation

The `containerappsbuilds` SDK allows for interaction with the Azure Resource Manager Service `containerapps` (API Version `2024-02-02-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/containerappsbuilds"
```


### Client Initialization

```go
client := containerappsbuilds.NewContainerAppsBuildsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ContainerAppsBuildsClient.ByContainerAppList`

```go
ctx := context.TODO()
id := containerappsbuilds.NewContainerAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppValue")

// alternatively `client.ByContainerAppList(ctx, id)` can be used to do batched pagination
items, err := client.ByContainerAppListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ContainerAppsBuildsClient.Delete`

```go
ctx := context.TODO()
id := containerappsbuilds.NewContainerAppBuildID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppValue", "buildValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ContainerAppsBuildsClient.Get`

```go
ctx := context.TODO()
id := containerappsbuilds.NewContainerAppBuildID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppValue", "buildValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
