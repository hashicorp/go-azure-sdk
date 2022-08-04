
## `github.com/hashicorp/go-azure-sdk/resource-manager/deviceprovisioningservices/2022-02-05/get` Documentation

The `get` SDK allows for interaction with the Azure Resource Manager Service `deviceprovisioningservices` (API Version `2022-02-05`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/deviceprovisioningservices/2022-02-05/get"
```


### Client Initialization

```go
client := get.NewGETClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GETClient.DpsCertificateGet`

```go
ctx := context.TODO()
id := get.NewCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "provisioningServiceValue", "certificateValue")

read, err := client.DpsCertificateGet(ctx, id, get.DefaultDpsCertificateGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GETClient.DpsCertificateList`

```go
ctx := context.TODO()
id := get.NewProvisioningServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "provisioningServiceValue")

read, err := client.DpsCertificateList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GETClient.IotDpsResourceGet`

```go
ctx := context.TODO()
id := get.NewProvisioningServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "provisioningServiceValue")

read, err := client.IotDpsResourceGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GETClient.IotDpsResourceGetPrivateEndpointConnection`

```go
ctx := context.TODO()
id := get.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue", "privateEndpointConnectionValue")

read, err := client.IotDpsResourceGetPrivateEndpointConnection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GETClient.IotDpsResourceGetPrivateLinkResources`

```go
ctx := context.TODO()
id := get.NewPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceValue", "groupIdValue")

read, err := client.IotDpsResourceGetPrivateLinkResources(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GETClient.IotDpsResourceListByResourceGroup`

```go
ctx := context.TODO()
id := get.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.IotDpsResourceListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.IotDpsResourceListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GETClient.IotDpsResourceListBySubscription`

```go
ctx := context.TODO()
id := get.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.IotDpsResourceListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.IotDpsResourceListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GETClient.IotDpsResourceListPrivateEndpointConnections`

```go
ctx := context.TODO()
id := get.NewProvisioningServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "provisioningServiceValue")

read, err := client.IotDpsResourceListPrivateEndpointConnections(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GETClient.IotDpsResourceListPrivateLinkResources`

```go
ctx := context.TODO()
id := get.NewProvisioningServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "provisioningServiceValue")

read, err := client.IotDpsResourceListPrivateLinkResources(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GETClient.IotDpsResourcelistValidSkus`

```go
ctx := context.TODO()
id := get.NewProvisioningServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "provisioningServiceValue")

// alternatively `client.IotDpsResourcelistValidSkus(ctx, id)` can be used to do batched pagination
items, err := client.IotDpsResourcelistValidSkusComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
