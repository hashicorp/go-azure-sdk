
## `github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2022-10-31/privateendpoints` Documentation

The `privateendpoints` SDK allows for interaction with the Azure Resource Manager Service `digitaltwins` (API Version `2022-10-31`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/digitaltwins/2022-10-31/privateendpoints"
```


### Client Initialization

```go
client := privateendpoints.NewPrivateEndpointsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateEndpointsClient.PrivateEndpointConnectionsCreateOrUpdate`

```go
ctx := context.TODO()
id := privateendpoints.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "digitalTwinsInstanceValue", "privateEndpointConnectionValue")

payload := privateendpoints.PrivateEndpointConnection{
	// ...
}


if err := client.PrivateEndpointConnectionsCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointsClient.PrivateEndpointConnectionsDelete`

```go
ctx := context.TODO()
id := privateendpoints.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "digitalTwinsInstanceValue", "privateEndpointConnectionValue")

if err := client.PrivateEndpointConnectionsDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointsClient.PrivateEndpointConnectionsGet`

```go
ctx := context.TODO()
id := privateendpoints.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "digitalTwinsInstanceValue", "privateEndpointConnectionValue")

read, err := client.PrivateEndpointConnectionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateEndpointsClient.PrivateEndpointConnectionsList`

```go
ctx := context.TODO()
id := privateendpoints.NewDigitalTwinsInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "digitalTwinsInstanceValue")

read, err := client.PrivateEndpointConnectionsList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateEndpointsClient.PrivateLinkResourcesGet`

```go
ctx := context.TODO()
id := privateendpoints.NewResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "digitalTwinsInstanceValue", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.PrivateLinkResourcesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateEndpointsClient.PrivateLinkResourcesList`

```go
ctx := context.TODO()
id := privateendpoints.NewDigitalTwinsInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "digitalTwinsInstanceValue")

read, err := client.PrivateLinkResourcesList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
