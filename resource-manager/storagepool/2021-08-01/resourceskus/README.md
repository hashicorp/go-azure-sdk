
## `github.com/hashicorp/go-azure-sdk/resource-manager/storagepool/2021-08-01/resourceskus` Documentation

The `resourceskus` SDK allows for interaction with the Azure Resource Manager Service `storagepool` (API Version `2021-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/storagepool/2021-08-01/resourceskus"
```


### Client Initialization

```go
client := resourceskus.NewResourceSkusClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `ResourceSkusClient.List`

```go
ctx := context.TODO()
id := resourceskus.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")
// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
