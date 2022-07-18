
## `github.com/hashicorp/go-azure-sdk/resource-manager/media/2021-05-01/media` Documentation

The `media` SDK allows for interaction with the Azure Resource Manager Service `media` (API Version `2021-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/media/2021-05-01/media"
```


### Client Initialization

```go
client := media.NewMediaClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MediaClient.LocationsCheckNameAvailability`

```go
ctx := context.TODO()
id := media.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := media.CheckNameAvailabilityInput{
	// ...
}


read, err := client.LocationsCheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.MediaservicesCreateOrUpdate`

```go
ctx := context.TODO()
id := media.NewMediaServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

payload := media.MediaService{
	// ...
}


read, err := client.MediaservicesCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.MediaservicesDelete`

```go
ctx := context.TODO()
id := media.NewMediaServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

read, err := client.MediaservicesDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.MediaservicesGet`

```go
ctx := context.TODO()
id := media.NewMediaServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

read, err := client.MediaservicesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.MediaservicesList`

```go
ctx := context.TODO()
id := media.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.MediaservicesList(ctx, id)` can be used to do batched pagination
items, err := client.MediaservicesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MediaClient.MediaservicesListBySubscription`

```go
ctx := context.TODO()
id := media.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.MediaservicesListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.MediaservicesListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MediaClient.MediaservicesListEdgePolicies`

```go
ctx := context.TODO()
id := media.NewMediaServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

payload := media.ListEdgePoliciesInput{
	// ...
}


read, err := client.MediaservicesListEdgePolicies(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.MediaservicesSyncStorageKeys`

```go
ctx := context.TODO()
id := media.NewMediaServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

payload := media.SyncStorageKeysInput{
	// ...
}


read, err := client.MediaservicesSyncStorageKeys(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.MediaservicesUpdate`

```go
ctx := context.TODO()
id := media.NewMediaServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

payload := media.MediaServiceUpdate{
	// ...
}


read, err := client.MediaservicesUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.PrivateEndpointConnectionsCreateOrUpdate`

```go
ctx := context.TODO()
id := media.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "nameValue")

payload := media.PrivateEndpointConnection{
	// ...
}


read, err := client.PrivateEndpointConnectionsCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.PrivateEndpointConnectionsDelete`

```go
ctx := context.TODO()
id := media.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "nameValue")

read, err := client.PrivateEndpointConnectionsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.PrivateEndpointConnectionsGet`

```go
ctx := context.TODO()
id := media.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "nameValue")

read, err := client.PrivateEndpointConnectionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.PrivateEndpointConnectionsList`

```go
ctx := context.TODO()
id := media.NewMediaServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

read, err := client.PrivateEndpointConnectionsList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.PrivateLinkResourcesGet`

```go
ctx := context.TODO()
id := media.NewPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "nameValue")

read, err := client.PrivateLinkResourcesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MediaClient.PrivateLinkResourcesList`

```go
ctx := context.TODO()
id := media.NewMediaServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

read, err := client.PrivateLinkResourcesList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
