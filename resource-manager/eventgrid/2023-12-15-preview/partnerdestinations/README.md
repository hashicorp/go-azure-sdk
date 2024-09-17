
## `github.com/hashicorp/go-azure-sdk/resource-manager/eventgrid/2023-12-15-preview/partnerdestinations` Documentation

The `partnerdestinations` SDK allows for interaction with Azure Resource Manager `eventgrid` (API Version `2023-12-15-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/eventgrid/2023-12-15-preview/partnerdestinations"
```


### Client Initialization

```go
client := partnerdestinations.NewPartnerDestinationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PartnerDestinationsClient.Activate`

```go
ctx := context.TODO()
id := partnerdestinations.NewPartnerDestinationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "partnerDestinationValue")

read, err := client.Activate(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PartnerDestinationsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := partnerdestinations.NewPartnerDestinationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "partnerDestinationValue")

payload := partnerdestinations.PartnerDestination{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `PartnerDestinationsClient.Delete`

```go
ctx := context.TODO()
id := partnerdestinations.NewPartnerDestinationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "partnerDestinationValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PartnerDestinationsClient.Get`

```go
ctx := context.TODO()
id := partnerdestinations.NewPartnerDestinationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "partnerDestinationValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PartnerDestinationsClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id, partnerdestinations.DefaultListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id, partnerdestinations.DefaultListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PartnerDestinationsClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id, partnerdestinations.DefaultListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id, partnerdestinations.DefaultListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PartnerDestinationsClient.Update`

```go
ctx := context.TODO()
id := partnerdestinations.NewPartnerDestinationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "partnerDestinationValue")

payload := partnerdestinations.PartnerDestinationUpdateParameters{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
