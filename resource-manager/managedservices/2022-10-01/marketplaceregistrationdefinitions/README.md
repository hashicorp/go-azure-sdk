
## `github.com/hashicorp/go-azure-sdk/resource-manager/managedservices/2022-10-01/marketplaceregistrationdefinitions` Documentation

The `marketplaceregistrationdefinitions` SDK allows for interaction with the Azure Resource Manager Service `managedservices` (API Version `2022-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/managedservices/2022-10-01/marketplaceregistrationdefinitions"
```


### Client Initialization

```go
client := marketplaceregistrationdefinitions.NewMarketplaceRegistrationDefinitionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MarketplaceRegistrationDefinitionsClient.Get`

```go
ctx := context.TODO()
id := marketplaceregistrationdefinitions.NewScopedMarketplaceRegistrationDefinitionID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "marketplaceIdentifierValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MarketplaceRegistrationDefinitionsClient.List`

```go
ctx := context.TODO()
id := marketplaceregistrationdefinitions.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.List(ctx, id, marketplaceregistrationdefinitions.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, marketplaceregistrationdefinitions.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MarketplaceRegistrationDefinitionsClient.WithoutScopeGet`

```go
ctx := context.TODO()
id := marketplaceregistrationdefinitions.NewMarketplaceRegistrationDefinitionID("marketplaceIdentifierValue")

read, err := client.WithoutScopeGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MarketplaceRegistrationDefinitionsClient.WithoutScopeList`

```go
ctx := context.TODO()


// alternatively `client.WithoutScopeList(ctx, marketplaceregistrationdefinitions.DefaultWithoutScopeListOperationOptions())` can be used to do batched pagination
items, err := client.WithoutScopeListComplete(ctx, marketplaceregistrationdefinitions.DefaultWithoutScopeListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
