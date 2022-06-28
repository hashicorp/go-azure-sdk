
## `github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2022-05-01/privatelinkresources` Documentation

The `privatelinkresources` SDK allows for interaction with the Azure Resource Manager Service `appconfiguration` (API Version `2022-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2022-05-01/privatelinkresources"
```


### Client Initialization

```go
client := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `PrivateLinkResourcesClient.Get`

```go
ctx := context.TODO()
id := privatelinkresources.NewPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "configStoreValue", "groupValue")
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
id := privatelinkresources.NewConfigurationStoreID("12345678-1234-9876-4563-123456789012", "example-resource-group", "configStoreValue")
// alternatively `client.ListByConfigurationStore(ctx, id)` can be used to do batched pagination
items, err := client.ListByConfigurationStoreComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
