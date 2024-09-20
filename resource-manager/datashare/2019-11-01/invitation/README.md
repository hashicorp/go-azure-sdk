
## `github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/invitation` Documentation

The `invitation` SDK allows for interaction with Azure Resource Manager `datashare` (API Version `2019-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datashare/2019-11-01/invitation"
```


### Client Initialization

```go
client := invitation.NewInvitationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InvitationClient.Create`

```go
ctx := context.TODO()
id := invitation.NewInvitationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "shareName", "invitationName")

payload := invitation.Invitation{
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


### Example Usage: `InvitationClient.Delete`

```go
ctx := context.TODO()
id := invitation.NewInvitationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "shareName", "invitationName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InvitationClient.Get`

```go
ctx := context.TODO()
id := invitation.NewInvitationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "shareName", "invitationName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InvitationClient.ListByShare`

```go
ctx := context.TODO()
id := invitation.NewShareID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "shareName")

// alternatively `client.ListByShare(ctx, id, invitation.DefaultListByShareOperationOptions())` can be used to do batched pagination
items, err := client.ListByShareComplete(ctx, id, invitation.DefaultListByShareOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
