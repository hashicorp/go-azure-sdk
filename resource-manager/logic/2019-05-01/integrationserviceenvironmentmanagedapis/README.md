
## `github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/integrationserviceenvironmentmanagedapis` Documentation

The `integrationserviceenvironmentmanagedapis` SDK allows for interaction with Azure Resource Manager `logic` (API Version `2019-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/integrationserviceenvironmentmanagedapis"
```


### Client Initialization

```go
client := integrationserviceenvironmentmanagedapis.NewIntegrationServiceEnvironmentManagedApisClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IntegrationServiceEnvironmentManagedApisClient.IntegrationServiceEnvironmentManagedApiOperationsList`

```go
ctx := context.TODO()
id := integrationserviceenvironmentmanagedapis.NewManagedApiID("12345678-1234-9876-4563-123456789012", "example-resource-group", "integrationServiceEnvironmentValue", "managedApiValue")

// alternatively `client.IntegrationServiceEnvironmentManagedApiOperationsList(ctx, id)` can be used to do batched pagination
items, err := client.IntegrationServiceEnvironmentManagedApiOperationsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IntegrationServiceEnvironmentManagedApisClient.List`

```go
ctx := context.TODO()
id := integrationserviceenvironmentmanagedapis.NewIntegrationServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "integrationServiceEnvironmentValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
