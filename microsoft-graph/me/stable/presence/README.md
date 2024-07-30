
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/presence` Documentation

The `presence` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/presence"
```


### Client Initialization

```go
client := presence.NewPresenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PresenceClient.CreatePresenceClearPresence`

```go
ctx := context.TODO()

payload := presence.CreatePresenceClearPresenceRequest{
	// ...
}


read, err := client.CreatePresenceClearPresence(ctx, payload)
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


read, err := client.CreatePresenceClearUserPreferredPresence(ctx)
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


read, err := client.DeletePresence(ctx)
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


read, err := client.GetPresence(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PresenceClient.SetMePresencePresence`

```go
ctx := context.TODO()

payload := presence.SetMePresencePresenceRequest{
	// ...
}


read, err := client.SetMePresencePresence(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PresenceClient.SetMePresenceStatusMessage`

```go
ctx := context.TODO()

payload := presence.SetMePresenceStatusMessageRequest{
	// ...
}


read, err := client.SetMePresenceStatusMessage(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PresenceClient.SetMePresenceUserPreferredPresence`

```go
ctx := context.TODO()

payload := presence.SetMePresenceUserPreferredPresenceRequest{
	// ...
}


read, err := client.SetMePresenceUserPreferredPresence(ctx, payload)
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
