
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-11-01/servertrustgroups` Documentation

The `servertrustgroups` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2021-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-11-01/servertrustgroups"
```


### Client Initialization

```go
client := servertrustgroups.NewServerTrustGroupsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServerTrustGroupsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := servertrustgroups.NewServerTrustGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue", "serverTrustGroupValue")

payload := servertrustgroups.ServerTrustGroup{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ServerTrustGroupsClient.Delete`

```go
ctx := context.TODO()
id := servertrustgroups.NewServerTrustGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue", "serverTrustGroupValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ServerTrustGroupsClient.Get`

```go
ctx := context.TODO()
id := servertrustgroups.NewServerTrustGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue", "serverTrustGroupValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServerTrustGroupsClient.ListByInstance`

```go
ctx := context.TODO()
id := servertrustgroups.NewManagedInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue")

// alternatively `client.ListByInstance(ctx, id)` can be used to do batched pagination
items, err := client.ListByInstanceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ServerTrustGroupsClient.ListByLocation`

```go
ctx := context.TODO()
id := servertrustgroups.NewProviderLocationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue")

// alternatively `client.ListByLocation(ctx, id)` can be used to do batched pagination
items, err := client.ListByLocationComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
