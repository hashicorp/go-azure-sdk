
## `github.com/hashicorp/go-azure-sdk/resource-manager/storage/2021-04-01/skus` Documentation

The `skus` SDK allows for interaction with the Azure Resource Manager Service `storage` (API Version `2021-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/storage/2021-04-01/skus"
```


### Client Initialization

```go
client := skus.NewSkusClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SkusClient.List`

```go
ctx := context.TODO()
id := skus.NewSubscriptionID()

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
