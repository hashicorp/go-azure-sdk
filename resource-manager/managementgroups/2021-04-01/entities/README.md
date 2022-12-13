
## `github.com/hashicorp/go-azure-sdk/resource-manager/managementgroups/2021-04-01/entities` Documentation

The `entities` SDK allows for interaction with the Azure Resource Manager Service `managementgroups` (API Version `2021-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/managementgroups/2021-04-01/entities"
```


### Client Initialization

```go
client := entities.NewEntitiesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EntitiesClient.List`

```go
ctx := context.TODO()


// alternatively `client.List(ctx, entities.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, entities.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
