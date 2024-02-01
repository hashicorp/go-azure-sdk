
## `github.com/hashicorp/go-azure-sdk/resource-manager/resources/2021-07-01/subscriptionfeatureregistrations` Documentation

The `subscriptionfeatureregistrations` SDK allows for interaction with the Azure Resource Manager Service `resources` (API Version `2021-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/resources/2021-07-01/subscriptionfeatureregistrations"
```


### Client Initialization

```go
client := subscriptionfeatureregistrations.NewSubscriptionFeatureRegistrationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SubscriptionFeatureRegistrationsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := subscriptionfeatureregistrations.NewSubscriptionFeatureRegistrationID("12345678-1234-9876-4563-123456789012", "featureProviderValue", "subscriptionFeatureRegistrationValue")

payload := subscriptionfeatureregistrations.SubscriptionFeatureRegistration{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SubscriptionFeatureRegistrationsClient.Delete`

```go
ctx := context.TODO()
id := subscriptionfeatureregistrations.NewSubscriptionFeatureRegistrationID("12345678-1234-9876-4563-123456789012", "featureProviderValue", "subscriptionFeatureRegistrationValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SubscriptionFeatureRegistrationsClient.Get`

```go
ctx := context.TODO()
id := subscriptionfeatureregistrations.NewSubscriptionFeatureRegistrationID("12345678-1234-9876-4563-123456789012", "featureProviderValue", "subscriptionFeatureRegistrationValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SubscriptionFeatureRegistrationsClient.ListAllBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListAllBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListAllBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SubscriptionFeatureRegistrationsClient.ListBySubscription`

```go
ctx := context.TODO()
id := subscriptionfeatureregistrations.NewFeatureProviderID("12345678-1234-9876-4563-123456789012", "featureProviderValue")

// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
