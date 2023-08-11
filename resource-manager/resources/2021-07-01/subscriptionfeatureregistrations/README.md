
## `github.com/hashicorp/go-azure-sdk/resource-manager/resources/2021-07-01/subscriptionfeatureregistrations` Documentation

The `subscriptionfeatureregistrations` SDK allows for interaction with the Azure Resource Manager Service `resources` (API Version `2021-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/resources/2021-07-01/subscriptionfeatureregistrations"
```


### Client Initialization

```go
client := subscriptionfeatureregistrations.NewSubscriptionFeatureRegistrationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SubscriptionFeatureRegistrationsClient.SubscriptionFeatureRegistrationsCreateOrUpdate`

```go
ctx := context.TODO()
id := subscriptionfeatureregistrations.NewSubscriptionFeatureRegistrationID("12345678-1234-9876-4563-123456789012", "featureProviderValue", "subscriptionFeatureRegistrationValue")

payload := subscriptionfeatureregistrations.SubscriptionFeatureRegistration{
	// ...
}


read, err := client.SubscriptionFeatureRegistrationsCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SubscriptionFeatureRegistrationsClient.SubscriptionFeatureRegistrationsDelete`

```go
ctx := context.TODO()
id := subscriptionfeatureregistrations.NewSubscriptionFeatureRegistrationID("12345678-1234-9876-4563-123456789012", "featureProviderValue", "subscriptionFeatureRegistrationValue")

read, err := client.SubscriptionFeatureRegistrationsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SubscriptionFeatureRegistrationsClient.SubscriptionFeatureRegistrationsGet`

```go
ctx := context.TODO()
id := subscriptionfeatureregistrations.NewSubscriptionFeatureRegistrationID("12345678-1234-9876-4563-123456789012", "featureProviderValue", "subscriptionFeatureRegistrationValue")

read, err := client.SubscriptionFeatureRegistrationsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SubscriptionFeatureRegistrationsClient.SubscriptionFeatureRegistrationsListAllBySubscription`

```go
ctx := context.TODO()
id := subscriptionfeatureregistrations.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.SubscriptionFeatureRegistrationsListAllBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.SubscriptionFeatureRegistrationsListAllBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SubscriptionFeatureRegistrationsClient.SubscriptionFeatureRegistrationsListBySubscription`

```go
ctx := context.TODO()
id := subscriptionfeatureregistrations.NewFeatureProviderID("12345678-1234-9876-4563-123456789012", "featureProviderValue")

// alternatively `client.SubscriptionFeatureRegistrationsListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.SubscriptionFeatureRegistrationsListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
