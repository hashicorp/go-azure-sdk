
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/authenticationqrcodepinmethodpin` Documentation

The `authenticationqrcodepinmethodpin` SDK allows for interaction with Microsoft Graph `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/authenticationqrcodepinmethodpin"
```


### Client Initialization

```go
client := authenticationqrcodepinmethodpin.NewAuthenticationQrCodePinMethodPinClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AuthenticationQrCodePinMethodPinClient.DeleteAuthenticationQrCodePinMethodPin`

```go
ctx := context.TODO()
id := authenticationqrcodepinmethodpin.NewUserID("userId")

read, err := client.DeleteAuthenticationQrCodePinMethodPin(ctx, id, authenticationqrcodepinmethodpin.DefaultDeleteAuthenticationQrCodePinMethodPinOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationQrCodePinMethodPinClient.GetAuthenticationQrCodePinMethodPin`

```go
ctx := context.TODO()
id := authenticationqrcodepinmethodpin.NewUserID("userId")

read, err := client.GetAuthenticationQrCodePinMethodPin(ctx, id, authenticationqrcodepinmethodpin.DefaultGetAuthenticationQrCodePinMethodPinOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationQrCodePinMethodPinClient.UpdateAuthenticationQrCodePinMethodPin`

```go
ctx := context.TODO()
id := authenticationqrcodepinmethodpin.NewUserID("userId")

payload := authenticationqrcodepinmethodpin.QrPin{
	// ...
}


read, err := client.UpdateAuthenticationQrCodePinMethodPin(ctx, id, payload, authenticationqrcodepinmethodpin.DefaultUpdateAuthenticationQrCodePinMethodPinOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationQrCodePinMethodPinClient.UpdateAuthenticationQrCodePinMethodPinPin`

```go
ctx := context.TODO()
id := authenticationqrcodepinmethodpin.NewUserID("userId")

payload := authenticationqrcodepinmethodpin.UpdateAuthenticationQrCodePinMethodPinPinRequest{
	// ...
}


read, err := client.UpdateAuthenticationQrCodePinMethodPinPin(ctx, id, payload, authenticationqrcodepinmethodpin.DefaultUpdateAuthenticationQrCodePinMethodPinPinOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
