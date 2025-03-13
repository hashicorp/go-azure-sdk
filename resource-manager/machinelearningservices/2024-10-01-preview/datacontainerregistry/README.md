
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/datacontainerregistry` Documentation

The `datacontainerregistry` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2024-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/datacontainerregistry"
```


### Client Initialization

```go
client := datacontainerregistry.NewDataContainerRegistryClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataContainerRegistryClient.RegistryDataContainersCreateOrUpdate`

```go
ctx := context.TODO()
id := datacontainerregistry.NewDataID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName", "dataName")

payload := datacontainerregistry.DataContainerResource{
	// ...
}


if err := client.RegistryDataContainersCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DataContainerRegistryClient.RegistryDataContainersDelete`

```go
ctx := context.TODO()
id := datacontainerregistry.NewDataID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName", "dataName")

if err := client.RegistryDataContainersDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DataContainerRegistryClient.RegistryDataContainersGet`

```go
ctx := context.TODO()
id := datacontainerregistry.NewDataID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName", "dataName")

read, err := client.RegistryDataContainersGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataContainerRegistryClient.RegistryDataContainersList`

```go
ctx := context.TODO()
id := datacontainerregistry.NewRegistryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "registryName")

// alternatively `client.RegistryDataContainersList(ctx, id, datacontainerregistry.DefaultRegistryDataContainersListOperationOptions())` can be used to do batched pagination
items, err := client.RegistryDataContainersListComplete(ctx, id, datacontainerregistry.DefaultRegistryDataContainersListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
