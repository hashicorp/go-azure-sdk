
## `github.com/hashicorp/go-azure-sdk/resource-manager/customproviders/2018-09-01-preview/associations` Documentation

The `associations` SDK allows for interaction with the Azure Resource Manager Service `customproviders` (API Version `2018-09-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/customproviders/2018-09-01-preview/associations"
```


### Client Initialization

```go
client := associations.NewAssociationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AssociationsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := associations.NewScopedAssociationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "associationValue")

payload := associations.Association{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AssociationsClient.Delete`

```go
ctx := context.TODO()
id := associations.NewScopedAssociationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "associationValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `AssociationsClient.Get`

```go
ctx := context.TODO()
id := associations.NewScopedAssociationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "associationValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AssociationsClient.ListAll`

```go
ctx := context.TODO()
id := associations.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.ListAll(ctx, id)` can be used to do batched pagination
items, err := client.ListAllComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
