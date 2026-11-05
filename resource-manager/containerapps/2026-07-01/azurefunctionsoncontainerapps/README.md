
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2026-07-01/azurefunctionsoncontainerapps` Documentation

The `azurefunctionsoncontainerapps` SDK allows for interaction with Azure Resource Manager `containerapps` (API Version `2026-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2026-07-01/azurefunctionsoncontainerapps"
```


### Client Initialization

```go
client := azurefunctionsoncontainerapps.NewAzureFunctionsOnContainerAppsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AzureFunctionsOnContainerAppsClient.ContainerAppsFunctionsGet`

```go
ctx := context.TODO()
id := azurefunctionsoncontainerapps.NewFunctionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "functionName")

read, err := client.ContainerAppsFunctionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AzureFunctionsOnContainerAppsClient.ContainerAppsFunctionsList`

```go
ctx := context.TODO()
id := azurefunctionsoncontainerapps.NewContainerAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName")

// alternatively `client.ContainerAppsFunctionsList(ctx, id)` can be used to do batched pagination
items, err := client.ContainerAppsFunctionsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AzureFunctionsOnContainerAppsClient.ContainerAppsRevisionFunctionsGet`

```go
ctx := context.TODO()
id := azurefunctionsoncontainerapps.NewRevisionFunctionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "revisionName", "functionName")

read, err := client.ContainerAppsRevisionFunctionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AzureFunctionsOnContainerAppsClient.ContainerAppsRevisionFunctionsList`

```go
ctx := context.TODO()
id := azurefunctionsoncontainerapps.NewRevisionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "revisionName")

// alternatively `client.ContainerAppsRevisionFunctionsList(ctx, id)` can be used to do batched pagination
items, err := client.ContainerAppsRevisionFunctionsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
