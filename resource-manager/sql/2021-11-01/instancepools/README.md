
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-11-01/instancepools` Documentation

The `instancepools` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2021-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-11-01/instancepools"
```


### Client Initialization

```go
client := instancepools.NewInstancePoolsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InstancePoolsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := instancepools.NewInstancePoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "instancePoolValue")

payload := instancepools.InstancePool{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `InstancePoolsClient.Delete`

```go
ctx := context.TODO()
id := instancepools.NewInstancePoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "instancePoolValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `InstancePoolsClient.Get`

```go
ctx := context.TODO()
id := instancepools.NewInstancePoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "instancePoolValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InstancePoolsClient.List`

```go
ctx := context.TODO()
id := instancepools.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `InstancePoolsClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := instancepools.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `InstancePoolsClient.Update`

```go
ctx := context.TODO()
id := instancepools.NewInstancePoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "instancePoolValue")

payload := instancepools.InstancePoolUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
