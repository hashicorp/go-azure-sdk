
## `github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2021-08-01/share` Documentation

The `share` SDK allows for interaction with the Azure Resource Manager Service `datashare` (API Version `2021-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2021-08-01/share"
```


### Client Initialization

```go
client := share.NewShareClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ShareClient.Create`

```go
ctx := context.TODO()
id := share.NewShareID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "shareValue")

payload := share.Share{
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


### Example Usage: `ShareClient.Delete`

```go
ctx := context.TODO()
id := share.NewShareID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "shareValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ShareClient.Get`

```go
ctx := context.TODO()
id := share.NewShareID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "shareValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ShareClient.ListByAccount`

```go
ctx := context.TODO()
id := share.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

// alternatively `client.ListByAccount(ctx, id, share.DefaultListByAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByAccountComplete(ctx, id, share.DefaultListByAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ShareClient.ListSynchronizationDetails`

```go
ctx := context.TODO()
id := share.NewShareID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "shareValue")

payload := share.ShareSynchronization{
	// ...
}


// alternatively `client.ListSynchronizationDetails(ctx, id, payload, share.DefaultListSynchronizationDetailsOperationOptions())` can be used to do batched pagination
items, err := client.ListSynchronizationDetailsComplete(ctx, id, payload, share.DefaultListSynchronizationDetailsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ShareClient.ListSynchronizations`

```go
ctx := context.TODO()
id := share.NewShareID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "shareValue")

// alternatively `client.ListSynchronizations(ctx, id, share.DefaultListSynchronizationsOperationOptions())` can be used to do batched pagination
items, err := client.ListSynchronizationsComplete(ctx, id, share.DefaultListSynchronizationsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ShareClient.ProviderShareSubscriptionsAdjust`

```go
ctx := context.TODO()
id := share.NewProviderShareSubscriptionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "shareValue", "providerShareSubscriptionIdValue")

payload := share.ProviderShareSubscription{
	// ...
}


read, err := client.ProviderShareSubscriptionsAdjust(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ShareClient.ProviderShareSubscriptionsGetByShare`

```go
ctx := context.TODO()
id := share.NewProviderShareSubscriptionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "shareValue", "providerShareSubscriptionIdValue")

read, err := client.ProviderShareSubscriptionsGetByShare(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ShareClient.ProviderShareSubscriptionsListByShare`

```go
ctx := context.TODO()
id := share.NewShareID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "shareValue")

// alternatively `client.ProviderShareSubscriptionsListByShare(ctx, id)` can be used to do batched pagination
items, err := client.ProviderShareSubscriptionsListByShareComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ShareClient.ProviderShareSubscriptionsReinstate`

```go
ctx := context.TODO()
id := share.NewProviderShareSubscriptionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "shareValue", "providerShareSubscriptionIdValue")

payload := share.ProviderShareSubscription{
	// ...
}


read, err := client.ProviderShareSubscriptionsReinstate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ShareClient.ProviderShareSubscriptionsRevoke`

```go
ctx := context.TODO()
id := share.NewProviderShareSubscriptionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "shareValue", "providerShareSubscriptionIdValue")

if err := client.ProviderShareSubscriptionsRevokeThenPoll(ctx, id); err != nil {
	// handle the error
}
```
