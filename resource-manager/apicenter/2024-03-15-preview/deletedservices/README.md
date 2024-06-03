
## `github.com/hashicorp/go-azure-sdk/resource-manager/apicenter/2024-03-15-preview/deletedservices` Documentation

The `deletedservices` SDK allows for interaction with the Azure Resource Manager Service `apicenter` (API Version `2024-03-15-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/apicenter/2024-03-15-preview/deletedservices"
```


### Client Initialization

```go
client := deletedservices.NewDeletedServicesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeletedServicesClient.Delete`

```go
ctx := context.TODO()
id := deletedservices.NewDeletedServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deletedServiceValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedServicesClient.Get`

```go
ctx := context.TODO()
id := deletedservices.NewDeletedServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "deletedServiceValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedServicesClient.List`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.List(ctx, id, deletedservices.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, deletedservices.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeletedServicesClient.ListBySubscription`

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
