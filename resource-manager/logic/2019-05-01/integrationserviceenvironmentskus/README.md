
## `github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/integrationserviceenvironmentskus` Documentation

The `integrationserviceenvironmentskus` SDK allows for interaction with the Azure Resource Manager Service `logic` (API Version `2019-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/logic/2019-05-01/integrationserviceenvironmentskus"
```


### Client Initialization

```go
client := integrationserviceenvironmentskus.NewIntegrationServiceEnvironmentSkusClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IntegrationServiceEnvironmentSkusClient.List`

```go
ctx := context.TODO()
id := integrationserviceenvironmentskus.NewIntegrationServiceEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "integrationServiceEnvironmentValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
