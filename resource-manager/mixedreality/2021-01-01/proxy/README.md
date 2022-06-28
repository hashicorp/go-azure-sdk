
## `github.com/hashicorp/go-azure-sdk/resource-manager/mixedreality/2021-01-01/proxy` Documentation

The `proxy` SDK allows for interaction with the Azure Resource Manager Service `mixedreality` (API Version `2021-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mixedreality/2021-01-01/proxy"
```


### Client Initialization

```go
client := proxy.NewProxyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `ProxyClient.CheckNameAvailabilityLocal`

```go
ctx := context.TODO()
id := proxy.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := proxy.CheckNameAvailabilityRequest{
	// ...
}

read, err := client.CheckNameAvailabilityLocal(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProxyClient.RemoteRenderingAccountsListBySubscription`

```go
ctx := context.TODO()
id := proxy.NewSubscriptionID()
// alternatively `client.RemoteRenderingAccountsListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.RemoteRenderingAccountsListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ProxyClient.SpatialAnchorsAccountsListBySubscription`

```go
ctx := context.TODO()
id := proxy.NewSubscriptionID()
// alternatively `client.SpatialAnchorsAccountsListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.SpatialAnchorsAccountsListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
