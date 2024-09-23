
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/authenticationmethod` Documentation

The `authenticationmethod` SDK allows for interaction with Microsoft Graph `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/authenticationmethod"
```


### Client Initialization

```go
client := authenticationmethod.NewAuthenticationMethodClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AuthenticationMethodClient.CreateAuthenticationMethod`

```go
ctx := context.TODO()

payload := authenticationmethod.AuthenticationMethod{
	// ...
}


read, err := client.CreateAuthenticationMethod(ctx, payload, authenticationmethod.DefaultCreateAuthenticationMethodOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationMethodClient.DisableAuthenticationMethodSmsSignIn`

```go
ctx := context.TODO()
id := authenticationmethod.NewMeAuthenticationMethodID("authenticationMethodId")

read, err := client.DisableAuthenticationMethodSmsSignIn(ctx, id, authenticationmethod.DefaultDisableAuthenticationMethodSmsSignInOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationMethodClient.EnableAuthenticationMethodSmsSignIn`

```go
ctx := context.TODO()
id := authenticationmethod.NewMeAuthenticationMethodID("authenticationMethodId")

read, err := client.EnableAuthenticationMethodSmsSignIn(ctx, id, authenticationmethod.DefaultEnableAuthenticationMethodSmsSignInOperationOptions())
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
id := authenticationmethod.NewMeAuthenticationMethodID("authenticationMethodId")

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


read, err := client.GetAuthenticationMethodsCount(ctx, authenticationmethod.DefaultGetAuthenticationMethodsCountOperationOptions())
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


// alternatively `client.ListAuthenticationMethods(ctx, authenticationmethod.DefaultListAuthenticationMethodsOperationOptions())` can be used to do batched pagination
items, err := client.ListAuthenticationMethodsComplete(ctx, authenticationmethod.DefaultListAuthenticationMethodsOperationOptions())
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
id := authenticationmethod.NewMeAuthenticationMethodID("authenticationMethodId")

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
id := authenticationmethod.NewMeAuthenticationMethodID("authenticationMethodId")

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
