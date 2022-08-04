
## `github.com/hashicorp/go-azure-sdk/resource-manager/deviceprovisioningservices/2022-02-05/delete` Documentation

The `delete` SDK allows for interaction with the Azure Resource Manager Service `deviceprovisioningservices` (API Version `2022-02-05`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/deviceprovisioningservices/2022-02-05/delete"
```


### Client Initialization

```go
client := delete.NewDELETEClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DELETEClient.DpsCertificateDelete`

```go
ctx := context.TODO()
id := delete.NewCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "provisioningServiceValue", "certificateValue")

read, err := client.DpsCertificateDelete(ctx, id, delete.DefaultDpsCertificateDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DELETEClient.IotDpsResourceDelete`

```go
ctx := context.TODO()
id := delete.NewProvisioningServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "provisioningServiceValue")

if err := client.IotDpsResourceDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DELETEClient.IotDpsResourceDeletePrivateEndpointConnection`

```go
ctx := context.TODO()
id := delete.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue", "privateEndpointConnectionValue")

if err := client.IotDpsResourceDeletePrivateEndpointConnectionThenPoll(ctx, id); err != nil {
	// handle the error
}
```
