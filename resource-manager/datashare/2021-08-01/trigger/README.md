
## `github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2021-08-01/trigger` Documentation

The `trigger` SDK allows for interaction with Azure Resource Manager `datashare` (API Version `2021-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2021-08-01/trigger"
```


### Client Initialization

```go
client := trigger.NewTriggerClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TriggerClient.Create`

```go
ctx := context.TODO()
id := trigger.NewTriggerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "shareSubscriptionName", "triggerName")

payload := trigger.Trigger{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `TriggerClient.Delete`

```go
ctx := context.TODO()
id := trigger.NewTriggerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "shareSubscriptionName", "triggerName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `TriggerClient.Get`

```go
ctx := context.TODO()
id := trigger.NewTriggerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "shareSubscriptionName", "triggerName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TriggerClient.ListByShareSubscription`

```go
ctx := context.TODO()
id := trigger.NewShareSubscriptionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "shareSubscriptionName")

// alternatively `client.ListByShareSubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListByShareSubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
