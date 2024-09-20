
## `github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/sharesubscription` Documentation

The `sharesubscription` SDK allows for interaction with Azure Resource Manager `datashare` (API Version `2019-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/sharesubscription"
```


### Client Initialization

```go
client := sharesubscription.NewShareSubscriptionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ShareSubscriptionClient.CancelSynchronization`

```go
ctx := context.TODO()
id := sharesubscription.NewShareSubscriptionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "shareSubscriptionName")

payload := sharesubscription.ShareSubscriptionSynchronization{
	// ...
}


if err := client.CancelSynchronizationThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ShareSubscriptionClient.ConsumerSourceDataSetsListByShareSubscription`

```go
ctx := context.TODO()
id := sharesubscription.NewShareSubscriptionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "shareSubscriptionName")

// alternatively `client.ConsumerSourceDataSetsListByShareSubscription(ctx, id)` can be used to do batched pagination
items, err := client.ConsumerSourceDataSetsListByShareSubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ShareSubscriptionClient.Create`

```go
ctx := context.TODO()
id := sharesubscription.NewShareSubscriptionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "shareSubscriptionName")

payload := sharesubscription.ShareSubscription{
	// ...
}


read, err := client.Create(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ShareSubscriptionClient.Delete`

```go
ctx := context.TODO()
id := sharesubscription.NewShareSubscriptionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "shareSubscriptionName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ShareSubscriptionClient.Get`

```go
ctx := context.TODO()
id := sharesubscription.NewShareSubscriptionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "shareSubscriptionName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ShareSubscriptionClient.ListByAccount`

```go
ctx := context.TODO()
id := sharesubscription.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName")

// alternatively `client.ListByAccount(ctx, id, sharesubscription.DefaultListByAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByAccountComplete(ctx, id, sharesubscription.DefaultListByAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ShareSubscriptionClient.ListSourceShareSynchronizationSettings`

```go
ctx := context.TODO()
id := sharesubscription.NewShareSubscriptionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "shareSubscriptionName")

// alternatively `client.ListSourceShareSynchronizationSettings(ctx, id)` can be used to do batched pagination
items, err := client.ListSourceShareSynchronizationSettingsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ShareSubscriptionClient.ListSynchronizationDetails`

```go
ctx := context.TODO()
id := sharesubscription.NewShareSubscriptionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "shareSubscriptionName")

payload := sharesubscription.ShareSubscriptionSynchronization{
	// ...
}


// alternatively `client.ListSynchronizationDetails(ctx, id, payload, sharesubscription.DefaultListSynchronizationDetailsOperationOptions())` can be used to do batched pagination
items, err := client.ListSynchronizationDetailsComplete(ctx, id, payload, sharesubscription.DefaultListSynchronizationDetailsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ShareSubscriptionClient.ListSynchronizations`

```go
ctx := context.TODO()
id := sharesubscription.NewShareSubscriptionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "shareSubscriptionName")

// alternatively `client.ListSynchronizations(ctx, id, sharesubscription.DefaultListSynchronizationsOperationOptions())` can be used to do batched pagination
items, err := client.ListSynchronizationsComplete(ctx, id, sharesubscription.DefaultListSynchronizationsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ShareSubscriptionClient.Synchronize`

```go
ctx := context.TODO()
id := sharesubscription.NewShareSubscriptionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "shareSubscriptionName")

payload := sharesubscription.Synchronize{
	// ...
}


if err := client.SynchronizeThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
