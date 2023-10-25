
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-11-01/deletedservers` Documentation

The `deletedservers` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2021-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-11-01/deletedservers"
```


### Client Initialization

```go
client := deletedservers.NewDeletedServersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeletedServersClient.Get`

```go
ctx := context.TODO()
id := deletedservers.NewDeletedServerID("12345678-1234-9876-4563-123456789012", "locationValue", "deletedServerValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedServersClient.List`

```go
ctx := context.TODO()
id := deletedservers.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeletedServersClient.ListByLocation`

```go
ctx := context.TODO()
id := deletedservers.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

// alternatively `client.ListByLocation(ctx, id)` can be used to do batched pagination
items, err := client.ListByLocationComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeletedServersClient.Recover`

```go
ctx := context.TODO()
id := deletedservers.NewDeletedServerID("12345678-1234-9876-4563-123456789012", "locationValue", "deletedServerValue")

if err := client.RecoverThenPoll(ctx, id); err != nil {
	// handle the error
}
```
