
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2021-08-01/openidconnectprovider` Documentation

The `openidconnectprovider` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2021-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2021-08-01/openidconnectprovider"
```


### Client Initialization

```go
client := openidconnectprovider.NewOpenidConnectProviderClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OpenidConnectProviderClient.OpenIdConnectProviderCreateOrUpdate`

```go
ctx := context.TODO()
id := openidconnectprovider.NewOpenidConnectProviderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "opidValue")

payload := openidconnectprovider.OpenidConnectProviderContract{
	// ...
}


read, err := client.OpenIdConnectProviderCreateOrUpdate(ctx, id, payload, openidconnectprovider.DefaultOpenIdConnectProviderCreateOrUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenidConnectProviderClient.OpenIdConnectProviderDelete`

```go
ctx := context.TODO()
id := openidconnectprovider.NewOpenidConnectProviderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "opidValue")

read, err := client.OpenIdConnectProviderDelete(ctx, id, openidconnectprovider.DefaultOpenIdConnectProviderDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenidConnectProviderClient.OpenIdConnectProviderGet`

```go
ctx := context.TODO()
id := openidconnectprovider.NewOpenidConnectProviderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "opidValue")

read, err := client.OpenIdConnectProviderGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenidConnectProviderClient.OpenIdConnectProviderGetEntityTag`

```go
ctx := context.TODO()
id := openidconnectprovider.NewOpenidConnectProviderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "opidValue")

read, err := client.OpenIdConnectProviderGetEntityTag(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenidConnectProviderClient.OpenIdConnectProviderListByService`

```go
ctx := context.TODO()
id := openidconnectprovider.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue")

// alternatively `client.OpenIdConnectProviderListByService(ctx, id, openidconnectprovider.DefaultOpenIdConnectProviderListByServiceOperationOptions())` can be used to do batched pagination
items, err := client.OpenIdConnectProviderListByServiceComplete(ctx, id, openidconnectprovider.DefaultOpenIdConnectProviderListByServiceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenidConnectProviderClient.OpenIdConnectProviderListSecrets`

```go
ctx := context.TODO()
id := openidconnectprovider.NewOpenidConnectProviderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "opidValue")

read, err := client.OpenIdConnectProviderListSecrets(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenidConnectProviderClient.OpenIdConnectProviderUpdate`

```go
ctx := context.TODO()
id := openidconnectprovider.NewOpenidConnectProviderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "opidValue")

payload := openidconnectprovider.OpenidConnectProviderUpdateContract{
	// ...
}


read, err := client.OpenIdConnectProviderUpdate(ctx, id, payload, openidconnectprovider.DefaultOpenIdConnectProviderUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
