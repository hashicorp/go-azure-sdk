
## `github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2026-06-01/newrelics` Documentation

The `newrelics` SDK allows for interaction with Azure Resource Manager `newrelic` (API Version `2026-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/newrelic/2026-06-01/newrelics"
```


### Client Initialization

```go
client := newrelics.NewNewRelicsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NewRelicsClient.AccountsList`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.AccountsList(ctx, id, newrelics.DefaultAccountsListOperationOptions())` can be used to do batched pagination
items, err := client.AccountsListComplete(ctx, id, newrelics.DefaultAccountsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NewRelicsClient.OrganizationsList`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.OrganizationsList(ctx, id, newrelics.DefaultOrganizationsListOperationOptions())` can be used to do batched pagination
items, err := client.OrganizationsListComplete(ctx, id, newrelics.DefaultOrganizationsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NewRelicsClient.PlansList`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.PlansList(ctx, id, newrelics.DefaultPlansListOperationOptions())` can be used to do batched pagination
items, err := client.PlansListComplete(ctx, id, newrelics.DefaultPlansListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NewRelicsClient.SaaSActivateResource`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := newrelics.ActivateSaaSParameterRequest{
	// ...
}


read, err := client.SaaSActivateResource(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
