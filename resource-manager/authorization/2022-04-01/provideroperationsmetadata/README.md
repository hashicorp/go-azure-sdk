
## `github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2022-04-01/provideroperationsmetadata` Documentation

The `provideroperationsmetadata` SDK allows for interaction with the Azure Resource Manager Service `authorization` (API Version `2022-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/authorization/2022-04-01/provideroperationsmetadata"
```


### Client Initialization

```go
client := provideroperationsmetadata.NewProviderOperationsMetadataClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ProviderOperationsMetadataClient.Get`

```go
ctx := context.TODO()
id := provideroperationsmetadata.NewProviderOperationID("providerOperationValue")

read, err := client.Get(ctx, id, provideroperationsmetadata.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProviderOperationsMetadataClient.List`

```go
ctx := context.TODO()


// alternatively `client.List(ctx, provideroperationsmetadata.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, provideroperationsmetadata.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
