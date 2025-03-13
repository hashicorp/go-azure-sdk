
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/dataversionregistry` Documentation

The `dataversionregistry` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2024-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/dataversionregistry"
```


### Client Initialization

```go
client := dataversionregistry.NewDataVersionRegistryClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataVersionRegistryClient.RegistryDataVersionsCreateOrGetStartPendingUpload`

```go
ctx := context.TODO()
id := dataversionregistry.NewVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName", "dataName", "versionName")

payload := dataversionregistry.PendingUploadRequestDto{
	// ...
}


read, err := client.RegistryDataVersionsCreateOrGetStartPendingUpload(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataVersionRegistryClient.RegistryDataVersionsCreateOrUpdate`

```go
ctx := context.TODO()
id := dataversionregistry.NewVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName", "dataName", "versionName")

payload := dataversionregistry.DataVersionBaseResource{
	// ...
}


if err := client.RegistryDataVersionsCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DataVersionRegistryClient.RegistryDataVersionsDelete`

```go
ctx := context.TODO()
id := dataversionregistry.NewVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName", "dataName", "versionName")

if err := client.RegistryDataVersionsDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DataVersionRegistryClient.RegistryDataVersionsGet`

```go
ctx := context.TODO()
id := dataversionregistry.NewVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName", "dataName", "versionName")

read, err := client.RegistryDataVersionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataVersionRegistryClient.RegistryDataVersionsList`

```go
ctx := context.TODO()
id := dataversionregistry.NewDataID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName", "dataName")

// alternatively `client.RegistryDataVersionsList(ctx, id, dataversionregistry.DefaultRegistryDataVersionsListOperationOptions())` can be used to do batched pagination
items, err := client.RegistryDataVersionsListComplete(ctx, id, dataversionregistry.DefaultRegistryDataVersionsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
