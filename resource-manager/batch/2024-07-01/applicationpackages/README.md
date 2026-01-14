
## `github.com/hashicorp/go-azure-sdk/resource-manager/batch/2024-07-01/applicationpackages` Documentation

The `applicationpackages` SDK allows for interaction with Azure Resource Manager `batch` (API Version `2024-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/batch/2024-07-01/applicationpackages"
```


### Client Initialization

```go
client := applicationpackages.NewApplicationPackagesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApplicationPackagesClient.ApplicationPackageActivate`

```go
ctx := context.TODO()
id := applicationpackages.NewVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountName", "applicationName", "versionName")

payload := applicationpackages.ActivateApplicationPackageParameters{
	// ...
}


read, err := client.ApplicationPackageActivate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationPackagesClient.ApplicationPackageCreate`

```go
ctx := context.TODO()
id := applicationpackages.NewVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountName", "applicationName", "versionName")

payload := applicationpackages.ApplicationPackage{
	// ...
}


read, err := client.ApplicationPackageCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationPackagesClient.ApplicationPackageDelete`

```go
ctx := context.TODO()
id := applicationpackages.NewVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountName", "applicationName", "versionName")

read, err := client.ApplicationPackageDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationPackagesClient.ApplicationPackageGet`

```go
ctx := context.TODO()
id := applicationpackages.NewVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountName", "applicationName", "versionName")

read, err := client.ApplicationPackageGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationPackagesClient.ApplicationPackageList`

```go
ctx := context.TODO()
id := applicationpackages.NewApplicationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountName", "applicationName")

// alternatively `client.ApplicationPackageList(ctx, id, applicationpackages.DefaultApplicationPackageListOperationOptions())` can be used to do batched pagination
items, err := client.ApplicationPackageListComplete(ctx, id, applicationpackages.DefaultApplicationPackageListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
