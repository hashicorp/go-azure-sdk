
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/presence` Documentation

The `presence` SDK allows for interaction with Microsoft Graph `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/presence"
```


### Client Initialization

```go
client := presence.NewPresenceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PresenceClient.ClearPresencePresence`

```go
ctx := context.TODO()

payload := presence.ClearPresencePresenceRequest{
	// ...
}


read, err := client.ClearPresencePresence(ctx, payload, presence.DefaultClearPresencePresenceOperationOptions())
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


read, err := client.ClearPresenceUserPreferredPresence(ctx, presence.DefaultClearPresenceUserPreferredPresenceOperationOptions())
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


read, err := client.DeletePresence(ctx, presence.DefaultDeletePresenceOperationOptions())
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


read, err := client.GetPresence(ctx, presence.DefaultGetPresenceOperationOptions())
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

payload := presence.SetPresencePresenceRequest{
	// ...
}


read, err := client.SetPresencePresence(ctx, payload, presence.DefaultSetPresencePresenceOperationOptions())
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

payload := presence.SetPresenceStatusMessageRequest{
	// ...
}


read, err := client.SetPresenceStatusMessage(ctx, payload, presence.DefaultSetPresenceStatusMessageOperationOptions())
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

payload := presence.SetPresenceUserPreferredPresenceRequest{
	// ...
}


read, err := client.SetPresenceUserPreferredPresence(ctx, payload, presence.DefaultSetPresenceUserPreferredPresenceOperationOptions())
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

payload := presence.Presence{
	// ...
}


read, err := client.UpdatePresence(ctx, payload, presence.DefaultUpdatePresenceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
