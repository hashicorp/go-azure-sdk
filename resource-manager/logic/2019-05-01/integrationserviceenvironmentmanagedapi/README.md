
## `github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/integrationserviceenvironmentmanagedapi` Documentation

The `integrationserviceenvironmentmanagedapi` SDK allows for interaction with the Azure Resource Manager Service `logic` (API Version `2019-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/integrationserviceenvironmentmanagedapi"
```


### Client Initialization

```go
client := integrationserviceenvironmentmanagedapi.NewIntegrationServiceEnvironmentManagedApiClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IntegrationServiceEnvironmentManagedApiClient.Delete`

```go
ctx := context.TODO()
id := integrationserviceenvironmentmanagedapi.NewManagedApiID("12345678-1234-9876-4563-123456789012", "example-resource-group", "integrationServiceEnvironmentValue", "apiValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `IntegrationServiceEnvironmentManagedApiClient.Get`

```go
ctx := context.TODO()
id := integrationserviceenvironmentmanagedapi.NewManagedApiID("12345678-1234-9876-4563-123456789012", "example-resource-group", "integrationServiceEnvironmentValue", "apiValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IntegrationServiceEnvironmentManagedApiClient.Put`

```go
ctx := context.TODO()
id := integrationserviceenvironmentmanagedapi.NewManagedApiID("12345678-1234-9876-4563-123456789012", "example-resource-group", "integrationServiceEnvironmentValue", "apiValue")

payload := integrationserviceenvironmentmanagedapi.IntegrationServiceEnvironmentManagedApi{
	// ...
}


if err := client.PutThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
