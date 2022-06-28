
## `github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/namespacesprivatelinkresources` Documentation

The `namespacesprivatelinkresources` SDK allows for interaction with the Azure Resource Manager Service `eventhub` (API Version `2018-01-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/namespacesprivatelinkresources"
```


### Client Initialization

```go
client := namespacesprivatelinkresources.NewNamespacesPrivateLinkResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `NamespacesPrivateLinkResourcesClient.PrivateLinkResourcesGet`

```go
ctx := context.TODO()
id := namespacesprivatelinkresources.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue")
read, err := client.PrivateLinkResourcesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
