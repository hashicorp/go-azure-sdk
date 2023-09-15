
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2023-11-15/apimconfig` Documentation

The `apimconfig` SDK allows for interaction with the Azure Resource Manager Service `security` (API Version `2023-11-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2023-11-15/apimconfig"
```


### Client Initialization

```go
client := apimconfig.NewAPIMConfigClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `APIMConfigClient.APICollectionsGetByAzureApiManagementService`

```go
ctx := context.TODO()
id := apimconfig.NewApiCollectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "apiIdValue")

read, err := client.APICollectionsGetByAzureApiManagementService(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `APIMConfigClient.APICollectionsListByAzureApiManagementService`

```go
ctx := context.TODO()
id := apimconfig.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue")

// alternatively `client.APICollectionsListByAzureApiManagementService(ctx, id)` can be used to do batched pagination
items, err := client.APICollectionsListByAzureApiManagementServiceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `APIMConfigClient.APICollectionsOffboardAzureApiManagementApi`

```go
ctx := context.TODO()
id := apimconfig.NewApiCollectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "apiIdValue")

read, err := client.APICollectionsOffboardAzureApiManagementApi(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `APIMConfigClient.APICollectionsOnboardAzureApiManagementApi`

```go
ctx := context.TODO()
id := apimconfig.NewApiCollectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "apiIdValue")

if err := client.APICollectionsOnboardAzureApiManagementApiThenPoll(ctx, id); err != nil {
	// handle the error
}
```
