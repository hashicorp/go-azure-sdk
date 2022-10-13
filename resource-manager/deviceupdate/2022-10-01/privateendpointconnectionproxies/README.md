
## `github.com/hashicorp/go-azure-sdk/resource-manager/deviceupdate/2022-10-01/privateendpointconnectionproxies` Documentation

The `privateendpointconnectionproxies` SDK allows for interaction with the Azure Resource Manager Service `deviceupdate` (API Version `2022-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/deviceupdate/2022-10-01/privateendpointconnectionproxies"
```


### Client Initialization

```go
client := privateendpointconnectionproxies.NewPrivateEndpointConnectionProxiesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateEndpointConnectionProxiesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := privateendpointconnectionproxies.NewPrivateEndpointConnectionProxyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "privateEndpointConnectionProxyIdValue")

payload := privateendpointconnectionproxies.PrivateEndpointConnectionProxy{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionProxiesClient.Delete`

```go
ctx := context.TODO()
id := privateendpointconnectionproxies.NewPrivateEndpointConnectionProxyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "privateEndpointConnectionProxyIdValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionProxiesClient.ListByAccount`

```go
ctx := context.TODO()
id := privateendpointconnectionproxies.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

read, err := client.ListByAccount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
