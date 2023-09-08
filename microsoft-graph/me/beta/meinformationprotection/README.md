
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meinformationprotection` Documentation

The `meinformationprotection` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meinformationprotection"
```


### Client Initialization

```go
client := meinformationprotection.NewMeInformationProtectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeInformationProtectionClient.CreateMeInformationProtectionDecryptBuffer`

```go
ctx := context.TODO()

payload := meinformationprotection.CreateMeInformationProtectionDecryptBufferRequest{
	// ...
}


read, err := client.CreateMeInformationProtectionDecryptBuffer(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionClient.CreateMeInformationProtectionEncryptBuffer`

```go
ctx := context.TODO()

payload := meinformationprotection.CreateMeInformationProtectionEncryptBufferRequest{
	// ...
}


read, err := client.CreateMeInformationProtectionEncryptBuffer(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionClient.CreateMeInformationProtectionSignDigest`

```go
ctx := context.TODO()

payload := meinformationprotection.CreateMeInformationProtectionSignDigestRequest{
	// ...
}


read, err := client.CreateMeInformationProtectionSignDigest(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionClient.CreateMeInformationProtectionVerifySignature`

```go
ctx := context.TODO()

payload := meinformationprotection.CreateMeInformationProtectionVerifySignatureRequest{
	// ...
}


read, err := client.CreateMeInformationProtectionVerifySignature(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionClient.DeleteMeInformationProtection`

```go
ctx := context.TODO()


read, err := client.DeleteMeInformationProtection(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionClient.GetMeInformationProtection`

```go
ctx := context.TODO()


read, err := client.GetMeInformationProtection(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionClient.UpdateMeInformationProtection`

```go
ctx := context.TODO()

payload := meinformationprotection.InformationProtection{
	// ...
}


read, err := client.UpdateMeInformationProtection(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
