
## `github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/usersession` Documentation

The `usersession` SDK allows for interaction with Azure Resource Manager `desktopvirtualization` (API Version `2024-04-08-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/desktopvirtualization/2024-04-08-preview/usersession"
```


### Client Initialization

```go
client := usersession.NewUserSessionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserSessionClient.Delete`

```go
ctx := context.TODO()
id := usersession.NewUserSessionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostPoolName", "sessionHostName", "userSessionId")

read, err := client.Delete(ctx, id, usersession.DefaultDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserSessionClient.Disconnect`

```go
ctx := context.TODO()
id := usersession.NewUserSessionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostPoolName", "sessionHostName", "userSessionId")

read, err := client.Disconnect(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserSessionClient.Get`

```go
ctx := context.TODO()
id := usersession.NewUserSessionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostPoolName", "sessionHostName", "userSessionId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserSessionClient.List`

```go
ctx := context.TODO()
id := usersession.NewSessionHostID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostPoolName", "sessionHostName")

// alternatively `client.List(ctx, id, usersession.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, usersession.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserSessionClient.ListByHostPool`

```go
ctx := context.TODO()
id := usersession.NewHostPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostPoolName")

// alternatively `client.ListByHostPool(ctx, id, usersession.DefaultListByHostPoolOperationOptions())` can be used to do batched pagination
items, err := client.ListByHostPoolComplete(ctx, id, usersession.DefaultListByHostPoolOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserSessionClient.SendMessage`

```go
ctx := context.TODO()
id := usersession.NewUserSessionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hostPoolName", "sessionHostName", "userSessionId")

payload := usersession.SendMessage{
	// ...
}


read, err := client.SendMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
