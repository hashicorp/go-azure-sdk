
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-11-01/contentproductpackages` Documentation

The `contentproductpackages` SDK allows for interaction with the Azure Resource Manager Service `securityinsights` (API Version `2023-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-11-01/contentproductpackages"
```


### Client Initialization

```go
client := contentproductpackages.NewContentProductPackagesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ContentProductPackagesClient.ProductPackageGet`

```go
ctx := context.TODO()
id := contentproductpackages.NewContentProductPackageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "packageIdValue")

read, err := client.ProductPackageGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContentProductPackagesClient.ProductPackagesList`

```go
ctx := context.TODO()
id := contentproductpackages.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

// alternatively `client.ProductPackagesList(ctx, id, contentproductpackages.DefaultProductPackagesListOperationOptions())` can be used to do batched pagination
items, err := client.ProductPackagesListComplete(ctx, id, contentproductpackages.DefaultProductPackagesListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
