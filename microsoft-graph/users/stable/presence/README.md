
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/presence` Documentation

The `presence` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/presence"
```


### Client Initialization

```go
client := presence.NewPresenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PresenceClient.ClearPresence`

```go
ctx := context.TODO()
id := presence.NewUserID("userIdValue")

payload := presence.ClearPresenceRequest{
	// ...
}


read, err := client.ClearPresence(ctx, id, payload)
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
id := presence.NewUserID("userIdValue")

read, err := client.ClearPresenceUserPreferredPresence(ctx, id)
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
id := presence.NewUserID("userIdValue")

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
id := presence.NewUserID("userIdValue")

read, err := client.GetPresence(ctx, id, presence.DefaultGetPresenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PresenceClient.SetPresence`

```go
ctx := context.TODO()
id := presence.NewUserID("userIdValue")

payload := presence.SetPresenceRequest{
	// ...
}


read, err := client.SetPresence(ctx, id, payload)
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
id := presence.NewUserID("userIdValue")

payload := presence.SetPresenceStatusMessageRequest{
	// ...
}


read, err := client.SetPresenceStatusMessage(ctx, id, payload)
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
id := presence.NewUserID("userIdValue")

payload := presence.SetPresenceUserPreferredPresenceRequest{
	// ...
}


read, err := client.SetPresenceUserPreferredPresence(ctx, id, payload)
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
id := presence.NewUserID("userIdValue")

payload := presence.Presence{
	// ...
}


read, err := client.UpdatePresence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
