
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/authenticationmethod` Documentation

The `authenticationmethod` SDK allows for interaction with Microsoft Graph `users` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/authenticationmethod"
```


### Client Initialization

```go
client := authenticationmethod.NewAuthenticationMethodClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AuthenticationMethodClient.CreateAuthenticationMethod`

```go
ctx := context.TODO()
id := authenticationmethod.NewUserID("userId")

payload := authenticationmethod.AuthenticationMethod{
	// ...
}


read, err := client.CreateAuthenticationMethod(ctx, id, payload, authenticationmethod.DefaultCreateAuthenticationMethodOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationMethodClient.GetAuthenticationMethod`

```go
ctx := context.TODO()
id := authenticationmethod.NewUserIdAuthenticationMethodID("userId", "authenticationMethodId")

read, err := client.GetAuthenticationMethod(ctx, id, authenticationmethod.DefaultGetAuthenticationMethodOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationMethodClient.GetAuthenticationMethodsCount`

```go
ctx := context.TODO()
id := authenticationmethod.NewUserID("userId")

read, err := client.GetAuthenticationMethodsCount(ctx, id, authenticationmethod.DefaultGetAuthenticationMethodsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationMethodClient.ListAuthenticationMethods`

```go
ctx := context.TODO()
id := authenticationmethod.NewUserID("userId")

// alternatively `client.ListAuthenticationMethods(ctx, id, authenticationmethod.DefaultListAuthenticationMethodsOperationOptions())` can be used to do batched pagination
items, err := client.ListAuthenticationMethodsComplete(ctx, id, authenticationmethod.DefaultListAuthenticationMethodsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AuthenticationMethodClient.ResetAuthenticationMethodPassword`

```go
ctx := context.TODO()
id := authenticationmethod.NewUserIdAuthenticationMethodID("userId", "authenticationMethodId")

payload := authenticationmethod.ResetAuthenticationMethodPasswordRequest{
	// ...
}


read, err := client.ResetAuthenticationMethodPassword(ctx, id, payload, authenticationmethod.DefaultResetAuthenticationMethodPasswordOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationMethodClient.UpdateAuthenticationMethod`

```go
ctx := context.TODO()
id := authenticationmethod.NewUserIdAuthenticationMethodID("userId", "authenticationMethodId")

payload := authenticationmethod.AuthenticationMethod{
	// ...
}


read, err := client.UpdateAuthenticationMethod(ctx, id, payload, authenticationmethod.DefaultUpdateAuthenticationMethodOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
