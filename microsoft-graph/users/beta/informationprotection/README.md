
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/informationprotection` Documentation

The `informationprotection` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/informationprotection"
```


### Client Initialization

```go
client := informationprotection.NewInformationProtectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InformationProtectionClient.CreateInformationProtectionDecryptBuffer`

```go
ctx := context.TODO()
id := informationprotection.NewUserID("userIdValue")

payload := informationprotection.CreateInformationProtectionDecryptBufferRequest{
	// ...
}


read, err := client.CreateInformationProtectionDecryptBuffer(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionClient.CreateInformationProtectionEncryptBuffer`

```go
ctx := context.TODO()
id := informationprotection.NewUserID("userIdValue")

payload := informationprotection.CreateInformationProtectionEncryptBufferRequest{
	// ...
}


read, err := client.CreateInformationProtectionEncryptBuffer(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionClient.CreateInformationProtectionSignDigest`

```go
ctx := context.TODO()
id := informationprotection.NewUserID("userIdValue")

payload := informationprotection.CreateInformationProtectionSignDigestRequest{
	// ...
}


read, err := client.CreateInformationProtectionSignDigest(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionClient.CreateInformationProtectionVerifySignature`

```go
ctx := context.TODO()
id := informationprotection.NewUserID("userIdValue")

payload := informationprotection.CreateInformationProtectionVerifySignatureRequest{
	// ...
}


read, err := client.CreateInformationProtectionVerifySignature(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionClient.DeleteInformationProtection`

```go
ctx := context.TODO()
id := informationprotection.NewUserID("userIdValue")

read, err := client.DeleteInformationProtection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionClient.GetInformationProtection`

```go
ctx := context.TODO()
id := informationprotection.NewUserID("userIdValue")

read, err := client.GetInformationProtection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionClient.UpdateInformationProtection`

```go
ctx := context.TODO()
id := informationprotection.NewUserID("userIdValue")

payload := informationprotection.InformationProtection{
	// ...
}


read, err := client.UpdateInformationProtection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
