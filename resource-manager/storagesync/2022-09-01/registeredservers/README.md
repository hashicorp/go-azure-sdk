
## `github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/registeredservers` Documentation

The `registeredservers` SDK allows for interaction with Azure Resource Manager `storagesync` (API Version `2022-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/registeredservers"
```


### Client Initialization

```go
client := registeredservers.NewRegisteredServersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RegisteredServersClient.Create`

```go
ctx := context.TODO()
id := registeredservers.NewRegisteredServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "serverId")

payload := registeredservers.RegisteredServerCreateParameters{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `RegisteredServersClient.Delete`

```go
ctx := context.TODO()
id := registeredservers.NewRegisteredServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "serverId")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `RegisteredServersClient.Get`

```go
ctx := context.TODO()
id := registeredservers.NewRegisteredServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "serverId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RegisteredServersClient.ListByStorageSyncService`

```go
ctx := context.TODO()
id := registeredservers.NewStorageSyncServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName")

// alternatively `client.ListByStorageSyncService(ctx, id)` can be used to do batched pagination
items, err := client.ListByStorageSyncServiceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RegisteredServersClient.TriggerRollover`

```go
ctx := context.TODO()
id := registeredservers.NewRegisteredServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "serverId")

payload := registeredservers.TriggerRolloverRequest{
	// ...
}


if err := client.TriggerRolloverThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `RegisteredServersClient.Update`

```go
ctx := context.TODO()
id := registeredservers.NewRegisteredServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceName", "serverId")

payload := registeredservers.RegisteredServerUpdateParameters{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
