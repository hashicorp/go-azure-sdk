
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2026-07-01/sandboxgroups` Documentation

The `sandboxgroups` SDK allows for interaction with Azure Resource Manager `containerapps` (API Version `2026-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2026-07-01/sandboxgroups"
```


### Client Initialization

```go
client := sandboxgroups.NewSandboxGroupsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SandboxGroupsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := sandboxgroups.NewSandboxGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sandboxGroupName")

payload := sandboxgroups.SandboxGroup{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SandboxGroupsClient.Delete`

```go
ctx := context.TODO()
id := sandboxgroups.NewSandboxGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sandboxGroupName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SandboxGroupsClient.Get`

```go
ctx := context.TODO()
id := sandboxgroups.NewSandboxGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sandboxGroupName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SandboxGroupsClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SandboxGroupsClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SandboxGroupsClient.Update`

```go
ctx := context.TODO()
id := sandboxgroups.NewSandboxGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sandboxGroupName")

payload := sandboxgroups.SandboxGroupPatch{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
