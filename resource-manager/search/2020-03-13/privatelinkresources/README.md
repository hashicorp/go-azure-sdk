
## `github.com/hashicorp/go-azure-sdk/resource-manager/search/2020-03-13/privatelinkresources` Documentation

The `privatelinkresources` SDK allows for interaction with the Azure Resource Manager Service `search` (API Version `2020-03-13`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/search/2020-03-13/privatelinkresources"
```


### Client Initialization

```go
client := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateLinkResourcesClient.ListSupported`

```go
ctx := context.TODO()
id := privatelinkresources.NewSearchServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "searchServiceValue")

read, err := client.ListSupported(ctx, id, privatelinkresources.DefaultListSupportedOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
