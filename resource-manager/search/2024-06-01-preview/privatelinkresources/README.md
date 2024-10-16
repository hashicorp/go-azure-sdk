
## `github.com/hashicorp/go-azure-sdk/resource-manager/search/2024-06-01-preview/privatelinkresources` Documentation

The `privatelinkresources` SDK allows for interaction with Azure Resource Manager `search` (API Version `2024-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/search/2024-06-01-preview/privatelinkresources"
```


### Client Initialization

```go
client := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateLinkResourcesClient.ListSupported`

```go
ctx := context.TODO()
id := privatelinkresources.NewSearchServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "searchServiceName")

read, err := client.ListSupported(ctx, id, privatelinkresources.DefaultListSupportedOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
