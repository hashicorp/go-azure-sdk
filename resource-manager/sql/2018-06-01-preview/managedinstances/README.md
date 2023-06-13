
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/managedinstances` Documentation

The `managedinstances` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2018-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2018-06-01-preview/managedinstances"
```


### Client Initialization

```go
client := managedinstances.NewManagedInstancesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedInstancesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := managedinstances.NewManagedInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue")

payload := managedinstances.ManagedInstance{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ManagedInstancesClient.Delete`

```go
ctx := context.TODO()
id := managedinstances.NewManagedInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ManagedInstancesClient.Get`

```go
ctx := context.TODO()
id := managedinstances.NewManagedInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedInstancesClient.List`

```go
ctx := context.TODO()
id := managedinstances.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedInstancesClient.ListByInstancePool`

```go
ctx := context.TODO()
id := managedinstances.NewInstancePoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "instancePoolValue")

// alternatively `client.ListByInstancePool(ctx, id)` can be used to do batched pagination
items, err := client.ListByInstancePoolComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedInstancesClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := managedinstances.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedInstancesClient.Update`

```go
ctx := context.TODO()
id := managedinstances.NewManagedInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue")

payload := managedinstances.ManagedInstanceUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
