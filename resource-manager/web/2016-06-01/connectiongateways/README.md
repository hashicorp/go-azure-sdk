
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2016-06-01/connectiongateways` Documentation

The `connectiongateways` SDK allows for interaction with Azure Resource Manager `web` (API Version `2016-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2016-06-01/connectiongateways"
```


### Client Initialization

```go
client := connectiongateways.NewConnectionGatewaysClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConnectionGatewaysClient.ConnectionGatewayInstallationsGet`

```go
ctx := context.TODO()
id := connectiongateways.NewConnectionGatewayInstallationID("12345678-1234-9876-4563-123456789012", "locationValue", "gatewayIdValue")

read, err := client.ConnectionGatewayInstallationsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConnectionGatewaysClient.ConnectionGatewayInstallationsList`

```go
ctx := context.TODO()
id := connectiongateways.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

read, err := client.ConnectionGatewayInstallationsList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConnectionGatewaysClient.ConnectionGatewaysCreateOrUpdate`

```go
ctx := context.TODO()
id := connectiongateways.NewConnectionGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectionGatewayValue")

payload := connectiongateways.ConnectionGatewayDefinition{
	// ...
}


read, err := client.ConnectionGatewaysCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConnectionGatewaysClient.ConnectionGatewaysDelete`

```go
ctx := context.TODO()
id := connectiongateways.NewConnectionGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectionGatewayValue")

read, err := client.ConnectionGatewaysDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConnectionGatewaysClient.ConnectionGatewaysGet`

```go
ctx := context.TODO()
id := connectiongateways.NewConnectionGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectionGatewayValue")

read, err := client.ConnectionGatewaysGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConnectionGatewaysClient.ConnectionGatewaysList`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.ConnectionGatewaysList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConnectionGatewaysClient.ConnectionGatewaysListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

read, err := client.ConnectionGatewaysListByResourceGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConnectionGatewaysClient.ConnectionGatewaysUpdate`

```go
ctx := context.TODO()
id := connectiongateways.NewConnectionGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectionGatewayValue")

payload := connectiongateways.ConnectionGatewayDefinition{
	// ...
}


read, err := client.ConnectionGatewaysUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
