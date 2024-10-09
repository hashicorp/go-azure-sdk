
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/containerappssourcecontrols` Documentation

The `containerappssourcecontrols` SDK allows for interaction with Azure Resource Manager `containerapps` (API Version `2024-02-02-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/containerappssourcecontrols"
```


### Client Initialization

```go
client := containerappssourcecontrols.NewContainerAppsSourceControlsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ContainerAppsSourceControlsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := containerappssourcecontrols.NewSourceControlID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "sourceControlName")

payload := containerappssourcecontrols.SourceControl{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload, containerappssourcecontrols.DefaultCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `ContainerAppsSourceControlsClient.Delete`

```go
ctx := context.TODO()
id := containerappssourcecontrols.NewSourceControlID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "sourceControlName")

if err := client.DeleteThenPoll(ctx, id, containerappssourcecontrols.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `ContainerAppsSourceControlsClient.Get`

```go
ctx := context.TODO()
id := containerappssourcecontrols.NewSourceControlID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "sourceControlName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContainerAppsSourceControlsClient.ListByContainerApp`

```go
ctx := context.TODO()
id := containerappssourcecontrols.NewContainerAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName")

// alternatively `client.ListByContainerApp(ctx, id)` can be used to do batched pagination
items, err := client.ListByContainerAppComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
