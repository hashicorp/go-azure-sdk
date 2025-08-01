
## `github.com/hashicorp/go-azure-sdk/resource-manager/communication/2025-05-01/emailservices` Documentation

The `emailservices` SDK allows for interaction with Azure Resource Manager `communication` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/communication/2025-05-01/emailservices"
```


### Client Initialization

```go
client := emailservices.NewEmailServicesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EmailServicesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := emailservices.NewEmailServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "emailServiceName")

payload := emailservices.EmailServiceResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `EmailServicesClient.Delete`

```go
ctx := context.TODO()
id := emailservices.NewEmailServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "emailServiceName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `EmailServicesClient.Get`

```go
ctx := context.TODO()
id := emailservices.NewEmailServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "emailServiceName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EmailServicesClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EmailServicesClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EmailServicesClient.ListVerifiedExchangeOnlineDomains`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.ListVerifiedExchangeOnlineDomains(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EmailServicesClient.Update`

```go
ctx := context.TODO()
id := emailservices.NewEmailServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "emailServiceName")

payload := emailservices.TaggedResource{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
