
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/associatedtenant` Documentation

The `associatedtenant` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/associatedtenant"
```


### Client Initialization

```go
client := associatedtenant.NewAssociatedTenantClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AssociatedTenantClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := associatedtenant.NewAssociatedTenantID("billingAccountValue", "associatedTenantValue")

payload := associatedtenant.AssociatedTenant{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AssociatedTenantClient.Delete`

```go
ctx := context.TODO()
id := associatedtenant.NewAssociatedTenantID("billingAccountValue", "associatedTenantValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `AssociatedTenantClient.Get`

```go
ctx := context.TODO()
id := associatedtenant.NewAssociatedTenantID("billingAccountValue", "associatedTenantValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AssociatedTenantClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := associatedtenant.NewBillingAccountID("billingAccountValue")

// alternatively `client.ListByBillingAccount(ctx, id, associatedtenant.DefaultListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id, associatedtenant.DefaultListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
