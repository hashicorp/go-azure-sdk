
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


### Example Usage: `PresenceClient.CreatePresenceClearPresence`

```go
ctx := context.TODO()
id := presence.NewUserID("userIdValue")

payload := presence.CreatePresenceClearPresenceRequest{
	// ...
}


read, err := client.CreatePresenceClearPresence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PresenceClient.CreatePresenceClearUserPreferredPresence`

```go
ctx := context.TODO()
id := presence.NewUserID("userIdValue")

read, err := client.CreatePresenceClearUserPreferredPresence(ctx, id)
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

read, err := client.DeletePresence(ctx, id)
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

read, err := client.GetPresence(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PresenceClient.SetUserPresencePresence`

```go
ctx := context.TODO()
id := presence.NewUserID("userIdValue")

payload := presence.SetUserPresencePresenceRequest{
	// ...
}


read, err := client.SetUserPresencePresence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PresenceClient.SetUserPresenceStatusMessage`

```go
ctx := context.TODO()
id := presence.NewUserID("userIdValue")

payload := presence.SetUserPresenceStatusMessageRequest{
	// ...
}


read, err := client.SetUserPresenceStatusMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PresenceClient.SetUserPresenceUserPreferredPresence`

```go
ctx := context.TODO()
id := presence.NewUserID("userIdValue")

payload := presence.SetUserPresenceUserPreferredPresenceRequest{
	// ...
}


read, err := client.SetUserPresenceUserPreferredPresence(ctx, id, payload)
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
