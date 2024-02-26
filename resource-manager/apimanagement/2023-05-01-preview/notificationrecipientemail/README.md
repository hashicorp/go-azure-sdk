
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2023-05-01-preview/notificationrecipientemail` Documentation

The `notificationrecipientemail` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2023-05-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2023-05-01-preview/notificationrecipientemail"
```


### Client Initialization

```go
client := notificationrecipientemail.NewNotificationRecipientEmailClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NotificationRecipientEmailClient.CheckEntityExists`

```go
ctx := context.TODO()
id := notificationrecipientemail.NewRecipientEmailID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "example", "recipientEmailValue")

read, err := client.CheckEntityExists(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NotificationRecipientEmailClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := notificationrecipientemail.NewRecipientEmailID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "example", "recipientEmailValue")

read, err := client.CreateOrUpdate(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NotificationRecipientEmailClient.Delete`

```go
ctx := context.TODO()
id := notificationrecipientemail.NewRecipientEmailID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "example", "recipientEmailValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NotificationRecipientEmailClient.ListByNotification`

```go
ctx := context.TODO()
id := notificationrecipientemail.NewNotificationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "example")

read, err := client.ListByNotification(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NotificationRecipientEmailClient.WorkspaceNotificationRecipientEmailCheckEntityExists`

```go
ctx := context.TODO()
id := notificationrecipientemail.NewNotificationRecipientEmailID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "workspaceIdValue", "example", "recipientEmailValue")

read, err := client.WorkspaceNotificationRecipientEmailCheckEntityExists(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NotificationRecipientEmailClient.WorkspaceNotificationRecipientEmailCreateOrUpdate`

```go
ctx := context.TODO()
id := notificationrecipientemail.NewNotificationRecipientEmailID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "workspaceIdValue", "example", "recipientEmailValue")

read, err := client.WorkspaceNotificationRecipientEmailCreateOrUpdate(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NotificationRecipientEmailClient.WorkspaceNotificationRecipientEmailDelete`

```go
ctx := context.TODO()
id := notificationrecipientemail.NewNotificationRecipientEmailID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "workspaceIdValue", "example", "recipientEmailValue")

read, err := client.WorkspaceNotificationRecipientEmailDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NotificationRecipientEmailClient.WorkspaceNotificationRecipientEmailListByNotification`

```go
ctx := context.TODO()
id := notificationrecipientemail.NewNotificationNotificationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "workspaceIdValue", "example")

read, err := client.WorkspaceNotificationRecipientEmailListByNotification(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
