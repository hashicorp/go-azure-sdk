
## `github.com/hashicorp/go-azure-sdk/resource-manager/mixedreality/2025-01-01/key` Documentation

The `key` SDK allows for interaction with Azure Resource Manager `mixedreality` (API Version `2025-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mixedreality/2025-01-01/key"
```


### Client Initialization

```go
client := key.NewKeyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `KeyClient.RemoteRenderingAccountsListKeys`

```go
ctx := context.TODO()
id := key.NewRemoteRenderingAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "remoteRenderingAccountName")

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
id := key.NewRemoteRenderingAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "remoteRenderingAccountName")

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
