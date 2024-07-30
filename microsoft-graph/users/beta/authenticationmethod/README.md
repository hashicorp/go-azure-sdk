
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/authenticationmethod` Documentation

The `authenticationmethod` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/authenticationmethod"
```


### Client Initialization

```go
client := authenticationmethod.NewAuthenticationMethodClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AuthenticationMethodClient.CreateAuthenticationMethod`

```go
ctx := context.TODO()
id := authenticationmethod.NewUserID("userIdValue")

payload := authenticationmethod.AuthenticationMethod{
	// ...
}


read, err := client.CreateAuthenticationMethod(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationMethodClient.CreateAuthenticationMethodDisableSmsSignIn`

```go
ctx := context.TODO()
id := authenticationmethod.NewUserIdAuthenticationMethodID("userIdValue", "authenticationMethodIdValue")

read, err := client.CreateAuthenticationMethodDisableSmsSignIn(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationMethodClient.CreateAuthenticationMethodEnableSmsSignIn`

```go
ctx := context.TODO()
id := authenticationmethod.NewUserIdAuthenticationMethodID("userIdValue", "authenticationMethodIdValue")

read, err := client.CreateAuthenticationMethodEnableSmsSignIn(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationMethodClient.CreateAuthenticationMethodResetPassword`

```go
ctx := context.TODO()
id := authenticationmethod.NewUserIdAuthenticationMethodID("userIdValue", "authenticationMethodIdValue")

payload := authenticationmethod.CreateAuthenticationMethodResetPasswordRequest{
	// ...
}


read, err := client.CreateAuthenticationMethodResetPassword(ctx, id, payload)
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
id := authenticationmethod.NewUserIdAuthenticationMethodID("userIdValue", "authenticationMethodIdValue")

read, err := client.GetAuthenticationMethod(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationMethodClient.GetAuthenticationMethodCount`

```go
ctx := context.TODO()
id := authenticationmethod.NewUserID("userIdValue")

read, err := client.GetAuthenticationMethodCount(ctx, id)
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
id := authenticationmethod.NewUserID("userIdValue")

// alternatively `client.ListAuthenticationMethods(ctx, id)` can be used to do batched pagination
items, err := client.ListAuthenticationMethodsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AuthenticationMethodClient.UpdateAuthenticationMethod`

```go
ctx := context.TODO()
id := authenticationmethod.NewUserIdAuthenticationMethodID("userIdValue", "authenticationMethodIdValue")

payload := authenticationmethod.AuthenticationMethod{
	// ...
}


read, err := client.UpdateAuthenticationMethod(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
