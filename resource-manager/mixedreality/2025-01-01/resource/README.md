
## `github.com/hashicorp/go-azure-sdk/resource-manager/mixedreality/2025-01-01/resource` Documentation

The `resource` SDK allows for interaction with Azure Resource Manager `mixedreality` (API Version `2025-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/mixedreality/2025-01-01/resource"
```


### Client Initialization

```go
client := resource.NewResourceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ResourceClient.RemoteRenderingAccountsCreate`

```go
ctx := context.TODO()
id := resource.NewRemoteRenderingAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "remoteRenderingAccountName")

payload := resource.RemoteRenderingAccount{
	// ...
}


read, err := client.RemoteRenderingAccountsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourceClient.RemoteRenderingAccountsDelete`

```go
ctx := context.TODO()
id := resource.NewRemoteRenderingAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "remoteRenderingAccountName")

read, err := client.RemoteRenderingAccountsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourceClient.RemoteRenderingAccountsGet`

```go
ctx := context.TODO()
id := resource.NewRemoteRenderingAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "remoteRenderingAccountName")

read, err := client.RemoteRenderingAccountsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourceClient.RemoteRenderingAccountsListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.RemoteRenderingAccountsListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.RemoteRenderingAccountsListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ResourceClient.RemoteRenderingAccountsListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.RemoteRenderingAccountsListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.RemoteRenderingAccountsListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ResourceClient.RemoteRenderingAccountsUpdate`

```go
ctx := context.TODO()
id := resource.NewRemoteRenderingAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "remoteRenderingAccountName")

payload := resource.RemoteRenderingAccount{
	// ...
}


read, err := client.RemoteRenderingAccountsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
