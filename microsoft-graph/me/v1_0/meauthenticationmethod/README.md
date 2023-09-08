
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/meauthenticationmethod` Documentation

The `meauthenticationmethod` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/meauthenticationmethod"
```


### Client Initialization

```go
client := meauthenticationmethod.NewMeAuthenticationMethodClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeAuthenticationMethodClient.CreateMeAuthenticationMethod`

```go
ctx := context.TODO()

payload := meauthenticationmethod.AuthenticationMethod{
	// ...
}


read, err := client.CreateMeAuthenticationMethod(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeAuthenticationMethodClient.CreateMeAuthenticationMethodByIdResetPassword`

```go
ctx := context.TODO()
id := meauthenticationmethod.NewMeAuthenticationMethodID("authenticationMethodIdValue")

payload := meauthenticationmethod.CreateMeAuthenticationMethodByIdResetPasswordRequest{
	// ...
}


read, err := client.CreateMeAuthenticationMethodByIdResetPassword(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeAuthenticationMethodClient.GetMeAuthenticationMethodById`

```go
ctx := context.TODO()
id := meauthenticationmethod.NewMeAuthenticationMethodID("authenticationMethodIdValue")

read, err := client.GetMeAuthenticationMethodById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeAuthenticationMethodClient.GetMeAuthenticationMethodCount`

```go
ctx := context.TODO()


read, err := client.GetMeAuthenticationMethodCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeAuthenticationMethodClient.ListMeAuthenticationMethods`

```go
ctx := context.TODO()


// alternatively `client.ListMeAuthenticationMethods(ctx)` can be used to do batched pagination
items, err := client.ListMeAuthenticationMethodsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeAuthenticationMethodClient.UpdateMeAuthenticationMethodById`

```go
ctx := context.TODO()
id := meauthenticationmethod.NewMeAuthenticationMethodID("authenticationMethodIdValue")

payload := meauthenticationmethod.AuthenticationMethod{
	// ...
}


read, err := client.UpdateMeAuthenticationMethodById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
