
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userpresence` Documentation

The `userpresence` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userpresence"
```


### Client Initialization

```go
client := userpresence.NewUserPresenceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserPresenceClient.CreateUserByIdPresenceClearPresence`

```go
ctx := context.TODO()
id := userpresence.NewUserID("userIdValue")

payload := userpresence.CreateUserByIdPresenceClearPresenceRequest{
	// ...
}


read, err := client.CreateUserByIdPresenceClearPresence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPresenceClient.CreateUserByIdPresenceClearUserPreferredPresence`

```go
ctx := context.TODO()
id := userpresence.NewUserID("userIdValue")

read, err := client.CreateUserByIdPresenceClearUserPreferredPresence(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPresenceClient.DeleteUserByIdPresence`

```go
ctx := context.TODO()
id := userpresence.NewUserID("userIdValue")

read, err := client.DeleteUserByIdPresence(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPresenceClient.GetUserByIdPresence`

```go
ctx := context.TODO()
id := userpresence.NewUserID("userIdValue")

read, err := client.GetUserByIdPresence(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPresenceClient.SetUserByIdPresencePresence`

```go
ctx := context.TODO()
id := userpresence.NewUserID("userIdValue")

payload := userpresence.SetUserByIdPresencePresenceRequest{
	// ...
}


read, err := client.SetUserByIdPresencePresence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPresenceClient.SetUserByIdPresenceStatusMessage`

```go
ctx := context.TODO()
id := userpresence.NewUserID("userIdValue")

payload := userpresence.SetUserByIdPresenceStatusMessageRequest{
	// ...
}


read, err := client.SetUserByIdPresenceStatusMessage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPresenceClient.SetUserByIdPresenceUserPreferredPresence`

```go
ctx := context.TODO()
id := userpresence.NewUserID("userIdValue")

payload := userpresence.SetUserByIdPresenceUserPreferredPresenceRequest{
	// ...
}


read, err := client.SetUserByIdPresenceUserPreferredPresence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserPresenceClient.UpdateUserByIdPresence`

```go
ctx := context.TODO()
id := userpresence.NewUserID("userIdValue")

payload := userpresence.Presence{
	// ...
}


read, err := client.UpdateUserByIdPresence(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
