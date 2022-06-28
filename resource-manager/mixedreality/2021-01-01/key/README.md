
## `github.com/hashicorp/go-azure-sdk/resource-manager/mixedreality/2021-01-01/key` Documentation

The `key` SDK allows for interaction with the Azure Resource Manager Service `mixedreality` (API Version `2021-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mixedreality/2021-01-01/key"
```


### Client Initialization

```go
client := key.NewKeyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `KeyClient.RemoteRenderingAccountsListKeys`

```go
ctx := context.TODO()
id := key.NewRemoteRenderingAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

read, err := client.RemoteRenderingAccountsListKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `KeyClient.RemoteRenderingAccountsRegenerateKeys`

```go
ctx := context.TODO()
id := key.NewRemoteRenderingAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

payload := key.AccountKeyRegenerateRequest{
	// ...
}


read, err := client.RemoteRenderingAccountsRegenerateKeys(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `KeyClient.SpatialAnchorsAccountsListKeys`

```go
ctx := context.TODO()
id := key.NewSpatialAnchorsAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

read, err := client.SpatialAnchorsAccountsListKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `KeyClient.SpatialAnchorsAccountsRegenerateKeys`

```go
ctx := context.TODO()
id := key.NewSpatialAnchorsAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

payload := key.AccountKeyRegenerateRequest{
	// ...
}


read, err := client.SpatialAnchorsAccountsRegenerateKeys(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
