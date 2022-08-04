
## `github.com/hashicorp/go-azure-sdk/resource-manager/deviceprovisioningservices/2022-02-05/post` Documentation

The `post` SDK allows for interaction with the Azure Resource Manager Service `deviceprovisioningservices` (API Version `2022-02-05`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/deviceprovisioningservices/2022-02-05/post"
```


### Client Initialization

```go
client := post.NewPOSTClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `POSTClient.DpsCertificateGenerateVerificationCode`

```go
ctx := context.TODO()
id := post.NewCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "provisioningServiceValue", "certificateValue")

read, err := client.DpsCertificateGenerateVerificationCode(ctx, id, post.DefaultDpsCertificateGenerateVerificationCodeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `POSTClient.DpsCertificateVerifyCertificate`

```go
ctx := context.TODO()
id := post.NewCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "provisioningServiceValue", "certificateValue")

payload := post.VerificationCodeRequest{
	// ...
}


read, err := client.DpsCertificateVerifyCertificate(ctx, id, payload, post.DefaultDpsCertificateVerifyCertificateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `POSTClient.IotDpsResourceCheckProvisioningServiceNameAvailability`

```go
ctx := context.TODO()
id := post.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := post.OperationInputs{
	// ...
}


read, err := client.IotDpsResourceCheckProvisioningServiceNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `POSTClient.IotDpsResourceListKeys`

```go
ctx := context.TODO()
id := post.NewProvisioningServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "provisioningServiceValue")

// alternatively `client.IotDpsResourceListKeys(ctx, id)` can be used to do batched pagination
items, err := client.IotDpsResourceListKeysComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `POSTClient.IotDpsResourceListKeysForKeyName`

```go
ctx := context.TODO()
id := post.NewKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "provisioningServiceValue", "keyValue")

read, err := client.IotDpsResourceListKeysForKeyName(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
