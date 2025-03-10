
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2025-03-01/contentpackages` Documentation

The `contentpackages` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2025-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2025-03-01/contentpackages"
```


### Client Initialization

```go
client := contentpackages.NewContentPackagesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ContentPackagesClient.ContentPackageInstall`

```go
ctx := context.TODO()
id := contentpackages.NewContentPackageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "packageId")

payload := contentpackages.PackageModel{
	// ...
}


read, err := client.ContentPackageInstall(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContentPackagesClient.ContentPackageUninstall`

```go
ctx := context.TODO()
id := contentpackages.NewContentPackageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "packageId")

read, err := client.ContentPackageUninstall(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContentPackagesClient.Get`

```go
ctx := context.TODO()
id := contentpackages.NewContentPackageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "packageId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContentPackagesClient.List`

```go
ctx := context.TODO()
id := contentpackages.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.List(ctx, id, contentpackages.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, contentpackages.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
