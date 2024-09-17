
## `github.com/hashicorp/go-azure-sdk/resource-manager/batch/2024-02-01/applicationpackage` Documentation

The `applicationpackage` SDK allows for interaction with Azure Resource Manager `batch` (API Version `2024-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/batch/2024-02-01/applicationpackage"
```


### Client Initialization

```go
client := applicationpackage.NewApplicationPackageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApplicationPackageClient.Activate`

```go
ctx := context.TODO()
id := applicationpackage.NewVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountValue", "applicationValue", "versionValue")

payload := applicationpackage.ActivateApplicationPackageParameters{
	// ...
}


read, err := client.Activate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationPackageClient.Create`

```go
ctx := context.TODO()
id := applicationpackage.NewVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountValue", "applicationValue", "versionValue")

payload := applicationpackage.ApplicationPackage{
	// ...
}


read, err := client.Create(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationPackageClient.Delete`

```go
ctx := context.TODO()
id := applicationpackage.NewVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountValue", "applicationValue", "versionValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationPackageClient.Get`

```go
ctx := context.TODO()
id := applicationpackage.NewVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountValue", "applicationValue", "versionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationPackageClient.List`

```go
ctx := context.TODO()
id := applicationpackage.NewApplicationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "batchAccountValue", "applicationValue")

// alternatively `client.List(ctx, id, applicationpackage.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, applicationpackage.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
