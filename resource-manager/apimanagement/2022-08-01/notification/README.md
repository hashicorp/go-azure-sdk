
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/notification` Documentation

The `notification` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2022-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/notification"
```


### Client Initialization

```go
client := notification.NewNotificationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NotificationClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := notification.NewNotificationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "AccountClosedPublisher")

read, err := client.CreateOrUpdate(ctx, id, notification.DefaultCreateOrUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NotificationClient.Get`

```go
ctx := context.TODO()
id := notification.NewNotificationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "AccountClosedPublisher")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NotificationClient.ListByService`

```go
ctx := context.TODO()
id := notification.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue")

// alternatively `client.ListByService(ctx, id, notification.DefaultListByServiceOperationOptions())` can be used to do batched pagination
items, err := client.ListByServiceComplete(ctx, id, notification.DefaultListByServiceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
