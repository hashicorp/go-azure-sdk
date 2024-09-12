
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/presence` Documentation

The `presence` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/presence"
```


### Client Initialization

```go
client := presence.NewPresenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PresenceClient.ClearPresence`

```go
ctx := context.TODO()

payload := presence.ClearPresenceRequest{
	// ...
}


read, err := client.ClearPresence(ctx, payload)
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


read, err := client.ClearPresenceUserPreferredPresence(ctx)
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


### Example Usage: `PresenceClient.SetPresence`

```go
ctx := context.TODO()

payload := presence.SetPresenceRequest{
	// ...
}


read, err := client.SetPresence(ctx, payload)
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


read, err := client.SetPresenceStatusMessage(ctx, payload)
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


read, err := client.SetPresenceUserPreferredPresence(ctx, payload)
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


read, err := client.UpdatePresence(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
