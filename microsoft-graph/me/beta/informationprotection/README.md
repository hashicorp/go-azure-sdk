
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/informationprotection` Documentation

The `informationprotection` SDK allows for interaction with Microsoft Graph `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/informationprotection"
```


### Client Initialization

```go
client := informationprotection.NewInformationProtectionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InformationProtectionClient.CreateInformationProtectionDecryptBuffer`

```go
ctx := context.TODO()

payload := informationprotection.CreateInformationProtectionDecryptBufferRequest{
	// ...
}


read, err := client.CreateInformationProtectionDecryptBuffer(ctx, payload, informationprotection.DefaultCreateInformationProtectionDecryptBufferOperationOptions())
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

payload := informationprotection.CreateInformationProtectionEncryptBufferRequest{
	// ...
}


read, err := client.CreateInformationProtectionEncryptBuffer(ctx, payload, informationprotection.DefaultCreateInformationProtectionEncryptBufferOperationOptions())
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

payload := informationprotection.CreateInformationProtectionSignDigestRequest{
	// ...
}


read, err := client.CreateInformationProtectionSignDigest(ctx, payload, informationprotection.DefaultCreateInformationProtectionSignDigestOperationOptions())
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

payload := informationprotection.CreateInformationProtectionVerifySignatureRequest{
	// ...
}


read, err := client.CreateInformationProtectionVerifySignature(ctx, payload, informationprotection.DefaultCreateInformationProtectionVerifySignatureOperationOptions())
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


read, err := client.DeleteInformationProtection(ctx, informationprotection.DefaultDeleteInformationProtectionOperationOptions())
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


read, err := client.GetInformationProtection(ctx, informationprotection.DefaultGetInformationProtectionOperationOptions())
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

payload := informationprotection.InformationProtection{
	// ...
}


read, err := client.UpdateInformationProtection(ctx, payload, informationprotection.DefaultUpdateInformationProtectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
