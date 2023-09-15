
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2023-11-15/d4apicollection` Documentation

The `d4apicollection` SDK allows for interaction with the Azure Resource Manager Service `security` (API Version `2023-11-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2023-11-15/d4apicollection"
```


### Client Initialization

```go
client := d4apicollection.NewD4APICollectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `D4APICollectionClient.APICollectionsGetByAzureApiManagementService`

```go
ctx := context.TODO()
id := d4apicollection.NewApiCollectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "apiIdValue")

read, err := client.APICollectionsGetByAzureApiManagementService(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
