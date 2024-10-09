
## `github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2024-05-01/privatelinkresources` Documentation

The `privatelinkresources` SDK allows for interaction with Azure Resource Manager `appconfiguration` (API Version `2024-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2024-05-01/privatelinkresources"
```


### Client Initialization

```go
client := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateLinkResourcesClient.Get`

```go
ctx := context.TODO()
id := privatelinkresources.NewPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "configurationStoreName", "privateLinkResourceName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateLinkResourcesClient.ListByConfigurationStore`

```go
ctx := context.TODO()
id := privatelinkresources.NewConfigurationStoreID("12345678-1234-9876-4563-123456789012", "example-resource-group", "configurationStoreName")

// alternatively `client.ListByConfigurationStore(ctx, id)` can be used to do batched pagination
items, err := client.ListByConfigurationStoreComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
