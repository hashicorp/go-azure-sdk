
## `github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2023-02-01/mhsmprivatelinkresources` Documentation

The `mhsmprivatelinkresources` SDK allows for interaction with the Azure Resource Manager Service `keyvault` (API Version `2023-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2023-02-01/mhsmprivatelinkresources"
```


### Client Initialization

```go
client := mhsmprivatelinkresources.NewMHSMPrivateLinkResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MHSMPrivateLinkResourcesClient.ListByMHSMResource`

```go
ctx := context.TODO()
id := mhsmprivatelinkresources.NewManagedHSMID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedHSMValue")

read, err := client.ListByMHSMResource(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
