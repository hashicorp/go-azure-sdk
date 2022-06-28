
## `github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/regions` Documentation

The `regions` SDK allows for interaction with the Azure Resource Manager Service `eventhub` (API Version `2018-01-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/regions"
```


### Client Initialization

```go
client := regions.NewRegionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RegionsClient.ListBySku`

```go
ctx := context.TODO()
id := regions.NewSkuID("12345678-1234-9876-4563-123456789012", "skuValue")

// alternatively `client.ListBySku(ctx, id)` can be used to do batched pagination
items, err := client.ListBySkuComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
