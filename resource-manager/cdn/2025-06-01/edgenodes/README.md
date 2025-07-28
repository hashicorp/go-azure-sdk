
## `github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/edgenodes` Documentation

The `edgenodes` SDK allows for interaction with Azure Resource Manager `cdn` (API Version `2025-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/edgenodes"
```


### Client Initialization

```go
client := edgenodes.NewEdgenodesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EdgenodesClient.List`

```go
ctx := context.TODO()


// alternatively `client.List(ctx)` can be used to do batched pagination
items, err := client.ListComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
