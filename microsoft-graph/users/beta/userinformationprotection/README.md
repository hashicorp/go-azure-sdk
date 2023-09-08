
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userinformationprotection` Documentation

The `userinformationprotection` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userinformationprotection"
```


### Client Initialization

```go
client := userinformationprotection.NewUserInformationProtectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserInformationProtectionClient.CreateUserByIdInformationProtectionDecryptBuffer`

```go
ctx := context.TODO()
id := userinformationprotection.NewUserID("userIdValue")

payload := userinformationprotection.CreateUserByIdInformationProtectionDecryptBufferRequest{
	// ...
}


read, err := client.CreateUserByIdInformationProtectionDecryptBuffer(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionClient.CreateUserByIdInformationProtectionEncryptBuffer`

```go
ctx := context.TODO()
id := userinformationprotection.NewUserID("userIdValue")

payload := userinformationprotection.CreateUserByIdInformationProtectionEncryptBufferRequest{
	// ...
}


read, err := client.CreateUserByIdInformationProtectionEncryptBuffer(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionClient.CreateUserByIdInformationProtectionSignDigest`

```go
ctx := context.TODO()
id := userinformationprotection.NewUserID("userIdValue")

payload := userinformationprotection.CreateUserByIdInformationProtectionSignDigestRequest{
	// ...
}


read, err := client.CreateUserByIdInformationProtectionSignDigest(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionClient.CreateUserByIdInformationProtectionVerifySignature`

```go
ctx := context.TODO()
id := userinformationprotection.NewUserID("userIdValue")

payload := userinformationprotection.CreateUserByIdInformationProtectionVerifySignatureRequest{
	// ...
}


read, err := client.CreateUserByIdInformationProtectionVerifySignature(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionClient.DeleteUserByIdInformationProtection`

```go
ctx := context.TODO()
id := userinformationprotection.NewUserID("userIdValue")

read, err := client.DeleteUserByIdInformationProtection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionClient.GetUserByIdInformationProtection`

```go
ctx := context.TODO()
id := userinformationprotection.NewUserID("userIdValue")

read, err := client.GetUserByIdInformationProtection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionClient.UpdateUserByIdInformationProtection`

```go
ctx := context.TODO()
id := userinformationprotection.NewUserID("userIdValue")

payload := userinformationprotection.InformationProtection{
	// ...
}


read, err := client.UpdateUserByIdInformationProtection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
