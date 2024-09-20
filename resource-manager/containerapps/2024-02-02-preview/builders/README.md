
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/builders` Documentation

The `builders` SDK allows for interaction with Azure Resource Manager `containerapps` (API Version `2024-02-02-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/builders"
```


### Client Initialization

```go
client := builders.NewBuildersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BuildersClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := builders.NewBuilderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "builderName")

payload := builders.BuilderResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BuildersClient.Delete`

```go
ctx := context.TODO()
id := builders.NewBuilderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "builderName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `BuildersClient.Get`

```go
ctx := context.TODO()
id := builders.NewBuilderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "builderName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BuildersClient.ListByResourceGroup`

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


### Example Usage: `BuildersClient.ListBySubscription`

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


### Example Usage: `BuildersClient.Update`

```go
ctx := context.TODO()
id := builders.NewBuilderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "builderName")

payload := builders.BuilderResourceUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
