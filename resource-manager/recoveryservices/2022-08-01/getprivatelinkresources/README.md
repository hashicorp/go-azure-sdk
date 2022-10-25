
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2022-08-01/getprivatelinkresources` Documentation

The `getprivatelinkresources` SDK allows for interaction with the Azure Resource Manager Service `recoveryservices` (API Version `2022-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2022-08-01/getprivatelinkresources"
```


### Client Initialization

```go
client := getprivatelinkresources.NewGetPrivateLinkResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GetPrivateLinkResourcesClient.PrivateLinkResourcesGet`

```go
ctx := context.TODO()
id := getprivatelinkresources.NewPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "privateLinkResourceValue")

read, err := client.PrivateLinkResourcesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
