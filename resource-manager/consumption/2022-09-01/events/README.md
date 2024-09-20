
## `github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2022-09-01/events` Documentation

The `events` SDK allows for interaction with Azure Resource Manager `consumption` (API Version `2022-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2022-09-01/events"
```


### Client Initialization

```go
client := events.NewEventsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventsClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := events.NewBillingAccountID("billingAccountId")

// alternatively `client.ListByBillingAccount(ctx, id, events.DefaultListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id, events.DefaultListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EventsClient.ListByBillingProfile`

```go
ctx := context.TODO()
id := events.NewBillingProfileID("billingAccountId", "billingProfileId")

// alternatively `client.ListByBillingProfile(ctx, id, events.DefaultListByBillingProfileOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingProfileComplete(ctx, id, events.DefaultListByBillingProfileOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
