
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/presence` Documentation

The `presence` SDK allows for interaction with Microsoft Graph `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/presence"
```


### Client Initialization

```go
client := presence.NewPresenceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PresenceClient.ClearPresencePresence`

```go
ctx := context.TODO()
id := presence.NewUserID("userId")

payload := presence.ClearPresencePresenceRequest{
	// ...
}


read, err := client.ClearPresencePresence(ctx, id, payload, presence.DefaultClearPresencePresenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PresenceClient.ClearPresenceUserPreferredPresence`

```go
ctx := context.TODO()
id := presence.NewUserID("userId")

read, err := client.ClearPresenceUserPreferredPresence(ctx, id, presence.DefaultClearPresenceUserPreferredPresenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PresenceClient.DeletePresence`

```go
ctx := context.TODO()
id := presence.NewUserID("userId")

read, err := client.DeletePresence(ctx, id, presence.DefaultDeletePresenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PresenceClient.GetPresence`

```go
ctx := context.TODO()
id := presence.NewUserID("userId")

read, err := client.GetPresence(ctx, id, presence.DefaultGetPresenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PresenceClient.SetPresencePresence`

```go
ctx := context.TODO()
id := presence.NewUserID("userId")

payload := presence.SetPresencePresenceRequest{
	// ...
}


read, err := client.SetPresencePresence(ctx, id, payload, presence.DefaultSetPresencePresenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PresenceClient.SetPresenceStatusMessage`

```go
ctx := context.TODO()
id := presence.NewUserID("userId")

payload := presence.SetPresenceStatusMessageRequest{
	// ...
}


read, err := client.SetPresenceStatusMessage(ctx, id, payload, presence.DefaultSetPresenceStatusMessageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PresenceClient.SetPresenceUserPreferredPresence`

```go
ctx := context.TODO()
id := presence.NewUserID("userId")

payload := presence.SetPresenceUserPreferredPresenceRequest{
	// ...
}


read, err := client.SetPresenceUserPreferredPresence(ctx, id, payload, presence.DefaultSetPresenceUserPreferredPresenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PresenceClient.UpdatePresence`

```go
ctx := context.TODO()
id := presence.NewUserID("userId")

payload := presence.Presence{
	// ...
}


read, err := client.UpdatePresence(ctx, id, payload, presence.DefaultUpdatePresenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
