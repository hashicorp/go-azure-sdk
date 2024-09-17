
## `github.com/hashicorp/go-azure-sdk/resource-manager/storagepool/2021-08-01/diskpoolzones` Documentation

The `diskpoolzones` SDK allows for interaction with Azure Resource Manager `storagepool` (API Version `2021-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/storagepool/2021-08-01/diskpoolzones"
```


### Client Initialization

```go
client := diskpoolzones.NewDiskPoolZonesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DiskPoolZonesClient.List`

```go
ctx := context.TODO()
id := diskpoolzones.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
