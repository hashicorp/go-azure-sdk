
## `github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/serverendpoints` Documentation

The `serverendpoints` SDK allows for interaction with Azure Resource Manager `storagesync` (API Version `2022-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/serverendpoints"
```


### Client Initialization

```go
client := serverendpoints.NewServerEndpointsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServerEndpointsClient.Create`

```go
ctx := context.TODO()
id := serverendpoints.NewServerEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "syncGroupName", "serverEndpointName")

payload := serverendpoints.ServerEndpointCreateParameters{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ServerEndpointsClient.Delete`

```go
ctx := context.TODO()
id := serverendpoints.NewServerEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "syncGroupName", "serverEndpointName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ServerEndpointsClient.Get`

```go
ctx := context.TODO()
id := serverendpoints.NewServerEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "syncGroupName", "serverEndpointName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServerEndpointsClient.ListBySyncGroup`

```go
ctx := context.TODO()
id := serverendpoints.NewSyncGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "syncGroupName")

// alternatively `client.ListBySyncGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListBySyncGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ServerEndpointsClient.RecallAction`

```go
ctx := context.TODO()
id := serverendpoints.NewServerEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "syncGroupName", "serverEndpointName")

payload := serverendpoints.RecallActionParameters{
	// ...
}


if err := client.RecallActionThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ServerEndpointsClient.Update`

```go
ctx := context.TODO()
id := serverendpoints.NewServerEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "syncGroupName", "serverEndpointName")

payload := serverendpoints.ServerEndpointUpdateParameters{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
