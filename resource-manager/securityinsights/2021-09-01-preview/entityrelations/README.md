
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2021-09-01-preview/entityrelations` Documentation

The `entityrelations` SDK allows for interaction with the Azure Resource Manager Service `securityinsights` (API Version `2021-09-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2021-09-01-preview/entityrelations"
```


### Client Initialization

```go
client := entityrelations.NewEntityRelationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EntityRelationsClient.EntitiesRelationsList`

```go
ctx := context.TODO()
id := entityrelations.NewEntityID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "entityIdValue")

// alternatively `client.EntitiesRelationsList(ctx, id, entityrelations.DefaultEntitiesRelationsListOperationOptions())` can be used to do batched pagination
items, err := client.EntitiesRelationsListComplete(ctx, id, entityrelations.DefaultEntitiesRelationsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EntityRelationsClient.GetRelation`

```go
ctx := context.TODO()
id := entityrelations.NewRelationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "entityIdValue", "relationValue")

read, err := client.GetRelation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
