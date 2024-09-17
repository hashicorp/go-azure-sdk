
## `github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/privatelinkresources` Documentation

The `privatelinkresources` SDK allows for interaction with Azure Resource Manager `storagesync` (API Version `2022-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/storagesync/2022-09-01/privatelinkresources"
```


### Client Initialization

```go
client := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateLinkResourcesClient.ListByStorageSyncService`

```go
ctx := context.TODO()
id := privatelinkresources.NewStorageSyncServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageSyncServiceValue")

read, err := client.ListByStorageSyncService(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
