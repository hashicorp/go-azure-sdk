
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/authenticationhardwareoathmethod` Documentation

The `authenticationhardwareoathmethod` SDK allows for interaction with Microsoft Graph `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/authenticationhardwareoathmethod"
```


### Client Initialization

```go
client := authenticationhardwareoathmethod.NewAuthenticationHardwareOathMethodClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AuthenticationHardwareOathMethodClient.AssignAuthenticationHardwareOathMethodsAndActivate`

```go
ctx := context.TODO()
id := authenticationhardwareoathmethod.NewUserID("userId")

payload := authenticationhardwareoathmethod.AssignAuthenticationHardwareOathMethodsAndActivateRequest{
	// ...
}


read, err := client.AssignAuthenticationHardwareOathMethodsAndActivate(ctx, id, payload, authenticationhardwareoathmethod.DefaultAssignAuthenticationHardwareOathMethodsAndActivateOperationOptions())
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
id := authenticationhardwareoathmethod.NewUserID("userId")

payload := authenticationhardwareoathmethod.AssignAuthenticationHardwareOathMethodsAndActivateBySerialNumberRequest{
	// ...
}


read, err := client.AssignAuthenticationHardwareOathMethodsAndActivateBySerialNumber(ctx, id, payload, authenticationhardwareoathmethod.DefaultAssignAuthenticationHardwareOathMethodsAndActivateBySerialNumberOperationOptions())
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
id := authenticationhardwareoathmethod.NewUserID("userId")

payload := authenticationhardwareoathmethod.HardwareOathAuthenticationMethod{
	// ...
}


read, err := client.CreateAuthenticationHardwareOathMethod(ctx, id, payload, authenticationhardwareoathmethod.DefaultCreateAuthenticationHardwareOathMethodOperationOptions())
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
id := authenticationhardwareoathmethod.NewUserIdAuthenticationHardwareOathMethodID("userId", "hardwareOathAuthenticationMethodId")

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
id := authenticationhardwareoathmethod.NewUserIdAuthenticationHardwareOathMethodID("userId", "hardwareOathAuthenticationMethodId")

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
id := authenticationhardwareoathmethod.NewUserIdAuthenticationHardwareOathMethodID("userId", "hardwareOathAuthenticationMethodId")

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
id := authenticationhardwareoathmethod.NewUserIdAuthenticationHardwareOathMethodID("userId", "hardwareOathAuthenticationMethodId")

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
id := authenticationhardwareoathmethod.NewUserID("userId")

read, err := client.GetAuthenticationHardwareOathMethodsCount(ctx, id, authenticationhardwareoathmethod.DefaultGetAuthenticationHardwareOathMethodsCountOperationOptions())
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
id := authenticationhardwareoathmethod.NewUserID("userId")

// alternatively `client.ListAuthenticationHardwareOathMethods(ctx, id, authenticationhardwareoathmethod.DefaultListAuthenticationHardwareOathMethodsOperationOptions())` can be used to do batched pagination
items, err := client.ListAuthenticationHardwareOathMethodsComplete(ctx, id, authenticationhardwareoathmethod.DefaultListAuthenticationHardwareOathMethodsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
