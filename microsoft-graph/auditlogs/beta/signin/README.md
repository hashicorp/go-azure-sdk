
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/auditlogs/beta/signin` Documentation

The `signin` SDK allows for interaction with Microsoft Graph `auditlogs` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/auditlogs/beta/signin"
```


### Client Initialization

```go
client := signin.NewSignInClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SignInClient.CreateSignIn`

```go
ctx := context.TODO()

payload := signin.SignIn{
	// ...
}


read, err := client.CreateSignIn(ctx, payload, signin.DefaultCreateSignInOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SignInClient.CreateSignInConfirmCompromised`

```go
ctx := context.TODO()

payload := signin.CreateSignInConfirmCompromisedRequest{
	// ...
}


read, err := client.CreateSignInConfirmCompromised(ctx, payload, signin.DefaultCreateSignInConfirmCompromisedOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SignInClient.CreateSignInConfirmSafe`

```go
ctx := context.TODO()

payload := signin.CreateSignInConfirmSafeRequest{
	// ...
}


read, err := client.CreateSignInConfirmSafe(ctx, payload, signin.DefaultCreateSignInConfirmSafeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SignInClient.DeleteSignIn`

```go
ctx := context.TODO()
id := signin.NewAuditLogSignInID("signInId")

read, err := client.DeleteSignIn(ctx, id, signin.DefaultDeleteSignInOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SignInClient.GetSignIn`

```go
ctx := context.TODO()
id := signin.NewAuditLogSignInID("signInId")

read, err := client.GetSignIn(ctx, id, signin.DefaultGetSignInOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SignInClient.GetSignInsCount`

```go
ctx := context.TODO()


read, err := client.GetSignInsCount(ctx, signin.DefaultGetSignInsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SignInClient.ListSignIns`

```go
ctx := context.TODO()


// alternatively `client.ListSignIns(ctx, signin.DefaultListSignInsOperationOptions())` can be used to do batched pagination
items, err := client.ListSignInsComplete(ctx, signin.DefaultListSignInsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SignInClient.UpdateSignIn`

```go
ctx := context.TODO()
id := signin.NewAuditLogSignInID("signInId")

payload := signin.SignIn{
	// ...
}


read, err := client.UpdateSignIn(ctx, id, payload, signin.DefaultUpdateSignInOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
