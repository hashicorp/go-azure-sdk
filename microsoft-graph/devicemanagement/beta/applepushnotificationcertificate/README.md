
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/applepushnotificationcertificate` Documentation

The `applepushnotificationcertificate` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/applepushnotificationcertificate"
```


### Client Initialization

```go
client := applepushnotificationcertificate.NewApplePushNotificationCertificateClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApplePushNotificationCertificateClient.CreateApplePushNotificationCertificateGenerateSigningRequest`

```go
ctx := context.TODO()


read, err := client.CreateApplePushNotificationCertificateGenerateSigningRequest(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplePushNotificationCertificateClient.DeleteApplePushNotificationCertificate`

```go
ctx := context.TODO()


read, err := client.DeleteApplePushNotificationCertificate(ctx, applepushnotificationcertificate.DefaultDeleteApplePushNotificationCertificateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplePushNotificationCertificateClient.GetApplePushNotificationCertificate`

```go
ctx := context.TODO()


read, err := client.GetApplePushNotificationCertificate(ctx, applepushnotificationcertificate.DefaultGetApplePushNotificationCertificateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplePushNotificationCertificateClient.UpdateApplePushNotificationCertificate`

```go
ctx := context.TODO()

payload := applepushnotificationcertificate.ApplePushNotificationCertificate{
	// ...
}


read, err := client.UpdateApplePushNotificationCertificate(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
