
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userauthenticationmethod` Documentation

The `userauthenticationmethod` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userauthenticationmethod"
```


### Client Initialization

```go
client := userauthenticationmethod.NewUserAuthenticationMethodClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserAuthenticationMethodClient.CreateUserByIdAuthenticationMethod`

```go
ctx := context.TODO()
id := userauthenticationmethod.NewUserID("userIdValue")

payload := userauthenticationmethod.AuthenticationMethod{
	// ...
}


read, err := client.CreateUserByIdAuthenticationMethod(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserAuthenticationMethodClient.CreateUserByIdAuthenticationMethodByIdDisableSmsSignIn`

```go
ctx := context.TODO()
id := userauthenticationmethod.NewUserAuthenticationMethodID("userIdValue", "authenticationMethodIdValue")

read, err := client.CreateUserByIdAuthenticationMethodByIdDisableSmsSignIn(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserAuthenticationMethodClient.CreateUserByIdAuthenticationMethodByIdEnableSmsSignIn`

```go
ctx := context.TODO()
id := userauthenticationmethod.NewUserAuthenticationMethodID("userIdValue", "authenticationMethodIdValue")

read, err := client.CreateUserByIdAuthenticationMethodByIdEnableSmsSignIn(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserAuthenticationMethodClient.CreateUserByIdAuthenticationMethodByIdResetPassword`

```go
ctx := context.TODO()
id := userauthenticationmethod.NewUserAuthenticationMethodID("userIdValue", "authenticationMethodIdValue")

payload := userauthenticationmethod.CreateUserByIdAuthenticationMethodByIdResetPasswordRequest{
	// ...
}


read, err := client.CreateUserByIdAuthenticationMethodByIdResetPassword(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserAuthenticationMethodClient.GetUserByIdAuthenticationMethodById`

```go
ctx := context.TODO()
id := userauthenticationmethod.NewUserAuthenticationMethodID("userIdValue", "authenticationMethodIdValue")

read, err := client.GetUserByIdAuthenticationMethodById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserAuthenticationMethodClient.GetUserByIdAuthenticationMethodCount`

```go
ctx := context.TODO()
id := userauthenticationmethod.NewUserID("userIdValue")

read, err := client.GetUserByIdAuthenticationMethodCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserAuthenticationMethodClient.ListUserByIdAuthenticationMethods`

```go
ctx := context.TODO()
id := userauthenticationmethod.NewUserID("userIdValue")

// alternatively `client.ListUserByIdAuthenticationMethods(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdAuthenticationMethodsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserAuthenticationMethodClient.UpdateUserByIdAuthenticationMethodById`

```go
ctx := context.TODO()
id := userauthenticationmethod.NewUserAuthenticationMethodID("userIdValue", "authenticationMethodIdValue")

payload := userauthenticationmethod.AuthenticationMethod{
	// ...
}


read, err := client.UpdateUserByIdAuthenticationMethodById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
