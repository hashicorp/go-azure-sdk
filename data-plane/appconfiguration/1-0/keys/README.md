
## `github.com/hashicorp/go-azure-sdk/data-plane/appconfiguration/1.0/keys` Documentation

The `keys` SDK allows for interaction with <unknown source data type> `appconfiguration` (API Version `1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/appconfiguration/1.0/keys"
```


### Client Initialization

```go
client := keys.NewKeysClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `KeysClient.CheckKeys`

```go
ctx := context.TODO()


read, err := client.CheckKeys(ctx, keys.DefaultCheckKeysOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `KeysClient.GetKeys`

```go
ctx := context.TODO()


// alternatively `client.GetKeys(ctx, keys.DefaultGetKeysOperationOptions())` can be used to do batched pagination
items, err := client.GetKeysComplete(ctx, keys.DefaultGetKeysOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
