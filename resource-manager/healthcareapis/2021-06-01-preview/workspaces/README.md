
## `github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/workspaces` Documentation

The `workspaces` SDK allows for interaction with the Azure Resource Manager Service `healthcareapis` (API Version `2021-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/workspaces"
```


### Client Initialization

```go
client := workspaces.NewWorkspacesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `WorkspacesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := workspaces.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

payload := workspaces.Workspace{
	// ...
}

future, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `WorkspacesClient.Delete`

```go
ctx := context.TODO()
id := workspaces.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")
future, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `WorkspacesClient.Get`

```go
ctx := context.TODO()
id := workspaces.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspacesClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := workspaces.NewResourceGroupID()
// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkspacesClient.ListBySubscription`

```go
ctx := context.TODO()
id := workspaces.NewSubscriptionID()
// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkspacesClient.Update`

```go
ctx := context.TODO()
id := workspaces.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

payload := workspaces.ResourceTags{
	// ...
}

future, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```
