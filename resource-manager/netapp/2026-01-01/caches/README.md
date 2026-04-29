
## `github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2026-01-01/caches` Documentation

The `caches` SDK allows for interaction with Azure Resource Manager `netapp` (API Version `2026-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2026-01-01/caches"
```


### Client Initialization

```go
client := caches.NewCachesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CachesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := caches.NewCacheID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName", "capacityPoolName", "cacheName")

payload := caches.Cache{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `CachesClient.Delete`

```go
ctx := context.TODO()
id := caches.NewCacheID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName", "capacityPoolName", "cacheName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `CachesClient.Get`

```go
ctx := context.TODO()
id := caches.NewCacheID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName", "capacityPoolName", "cacheName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CachesClient.List`

```go
ctx := context.TODO()
id := caches.NewCapacityPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName", "capacityPoolName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CachesClient.ListPeeringPassphrases`

```go
ctx := context.TODO()
id := caches.NewCacheID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName", "capacityPoolName", "cacheName")

read, err := client.ListPeeringPassphrases(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CachesClient.PoolChange`

```go
ctx := context.TODO()
id := caches.NewCacheID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName", "capacityPoolName", "cacheName")

payload := caches.PoolChangeRequest{
	// ...
}


if err := client.PoolChangeThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `CachesClient.ResetSmbPassword`

```go
ctx := context.TODO()
id := caches.NewCacheID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName", "capacityPoolName", "cacheName")

if err := client.ResetSmbPasswordThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `CachesClient.Update`

```go
ctx := context.TODO()
id := caches.NewCacheID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName", "capacityPoolName", "cacheName")

payload := caches.CacheUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
