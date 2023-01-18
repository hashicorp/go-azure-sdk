
## `github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/notificationchannels` Documentation

The `notificationchannels` SDK allows for interaction with the Azure Resource Manager Service `devtestlab` (API Version `2018-09-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/notificationchannels"
```


### Client Initialization

```go
client := notificationchannels.NewNotificationChannelsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NotificationChannelsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := notificationchannels.NewNotificationChannelID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "nameValue")

payload := notificationchannels.NotificationChannel{
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


### Example Usage: `NotificationChannelsClient.Delete`

```go
ctx := context.TODO()
id := notificationchannels.NewNotificationChannelID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "nameValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NotificationChannelsClient.Get`

```go
ctx := context.TODO()
id := notificationchannels.NewNotificationChannelID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "nameValue")

read, err := client.Get(ctx, id, notificationchannels.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NotificationChannelsClient.List`

```go
ctx := context.TODO()
id := notificationchannels.NewLabID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue")

// alternatively `client.List(ctx, id, notificationchannels.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, notificationchannels.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NotificationChannelsClient.Notify`

```go
ctx := context.TODO()
id := notificationchannels.NewNotificationChannelID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "nameValue")

payload := notificationchannels.NotifyParameters{
	// ...
}


read, err := client.Notify(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NotificationChannelsClient.Update`

```go
ctx := context.TODO()
id := notificationchannels.NewNotificationChannelID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "nameValue")

payload := notificationchannels.UpdateResource{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
