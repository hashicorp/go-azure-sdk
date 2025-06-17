
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/authenticationhardwareoathmethod` Documentation

The `authenticationhardwareoathmethod` SDK allows for interaction with Microsoft Graph `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/authenticationhardwareoathmethod"
```


### Client Initialization

```go
client := authenticationhardwareoathmethod.NewAuthenticationHardwareOathMethodClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AuthenticationHardwareOathMethodClient.AssignAuthenticationHardwareOathMethodsAndActivate`

```go
ctx := context.TODO()

payload := authenticationhardwareoathmethod.AssignAuthenticationHardwareOathMethodsAndActivateRequest{
	// ...
}


read, err := client.AssignAuthenticationHardwareOathMethodsAndActivate(ctx, payload, authenticationhardwareoathmethod.DefaultAssignAuthenticationHardwareOathMethodsAndActivateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationHardwareOathMethodClient.AssignAuthenticationHardwareOathMethodsAndActivateBySerialNumber`

```go
ctx := context.TODO()

payload := authenticationhardwareoathmethod.AssignAuthenticationHardwareOathMethodsAndActivateBySerialNumberRequest{
	// ...
}


read, err := client.AssignAuthenticationHardwareOathMethodsAndActivateBySerialNumber(ctx, payload, authenticationhardwareoathmethod.DefaultAssignAuthenticationHardwareOathMethodsAndActivateBySerialNumberOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationHardwareOathMethodClient.CreateAuthenticationHardwareOathMethod`

```go
ctx := context.TODO()

payload := authenticationhardwareoathmethod.HardwareOathAuthenticationMethod{
	// ...
}


read, err := client.CreateAuthenticationHardwareOathMethod(ctx, payload, authenticationhardwareoathmethod.DefaultCreateAuthenticationHardwareOathMethodOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationHardwareOathMethodClient.CreateAuthenticationHardwareOathMethodActivate`

```go
ctx := context.TODO()
id := authenticationhardwareoathmethod.NewMeAuthenticationHardwareOathMethodID("hardwareOathAuthenticationMethodId")

payload := authenticationhardwareoathmethod.CreateAuthenticationHardwareOathMethodActivateRequest{
	// ...
}


read, err := client.CreateAuthenticationHardwareOathMethodActivate(ctx, id, payload, authenticationhardwareoathmethod.DefaultCreateAuthenticationHardwareOathMethodActivateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationHardwareOathMethodClient.CreateAuthenticationHardwareOathMethodDeactivate`

```go
ctx := context.TODO()
id := authenticationhardwareoathmethod.NewMeAuthenticationHardwareOathMethodID("hardwareOathAuthenticationMethodId")

read, err := client.CreateAuthenticationHardwareOathMethodDeactivate(ctx, id, authenticationhardwareoathmethod.DefaultCreateAuthenticationHardwareOathMethodDeactivateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationHardwareOathMethodClient.DeleteAuthenticationHardwareOathMethod`

```go
ctx := context.TODO()
id := authenticationhardwareoathmethod.NewMeAuthenticationHardwareOathMethodID("hardwareOathAuthenticationMethodId")

read, err := client.DeleteAuthenticationHardwareOathMethod(ctx, id, authenticationhardwareoathmethod.DefaultDeleteAuthenticationHardwareOathMethodOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationHardwareOathMethodClient.GetAuthenticationHardwareOathMethod`

```go
ctx := context.TODO()
id := authenticationhardwareoathmethod.NewMeAuthenticationHardwareOathMethodID("hardwareOathAuthenticationMethodId")

read, err := client.GetAuthenticationHardwareOathMethod(ctx, id, authenticationhardwareoathmethod.DefaultGetAuthenticationHardwareOathMethodOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationHardwareOathMethodClient.GetAuthenticationHardwareOathMethodsCount`

```go
ctx := context.TODO()


read, err := client.GetAuthenticationHardwareOathMethodsCount(ctx, authenticationhardwareoathmethod.DefaultGetAuthenticationHardwareOathMethodsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AuthenticationHardwareOathMethodClient.ListAuthenticationHardwareOathMethods`

```go
ctx := context.TODO()


// alternatively `client.ListAuthenticationHardwareOathMethods(ctx, authenticationhardwareoathmethod.DefaultListAuthenticationHardwareOathMethodsOperationOptions())` can be used to do batched pagination
items, err := client.ListAuthenticationHardwareOathMethodsComplete(ctx, authenticationhardwareoathmethod.DefaultListAuthenticationHardwareOathMethodsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
