
## `github.com/hashicorp/go-azure-sdk/resource-manager/deviceprovisioningservices/2022-02-05/put` Documentation

The `put` SDK allows for interaction with the Azure Resource Manager Service `deviceprovisioningservices` (API Version `2022-02-05`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/deviceprovisioningservices/2022-02-05/put"
```


### Client Initialization

```go
client := put.NewPUTClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PUTClient.DpsCertificateCreateOrUpdate`

```go
ctx := context.TODO()
id := put.NewCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "provisioningServiceValue", "certificateValue")

payload := put.CertificateResponse{
	// ...
}


read, err := client.DpsCertificateCreateOrUpdate(ctx, id, payload, put.DefaultDpsCertificateCreateOrUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PUTClient.IotDpsResourceCreateOrUpdate`

```go
ctx := context.TODO()
id := put.NewProvisioningServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "provisioningServiceValue")

payload := put.ProvisioningServiceDescription{
	// ...
}


if err := client.IotDpsResourceCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `PUTClient.IotDpsResourceCreateOrUpdatePrivateEndpointConnection`

```go
ctx := context.TODO()
id := put.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue", "privateEndpointConnectionValue")

payload := put.PrivateEndpointConnection{
	// ...
}


if err := client.IotDpsResourceCreateOrUpdatePrivateEndpointConnectionThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
