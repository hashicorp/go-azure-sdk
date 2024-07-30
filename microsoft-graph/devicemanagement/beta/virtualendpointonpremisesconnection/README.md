
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointonpremisesconnection` Documentation

The `virtualendpointonpremisesconnection` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointonpremisesconnection"
```


### Client Initialization

```go
client := virtualendpointonpremisesconnection.NewVirtualEndpointOnPremisesConnectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VirtualEndpointOnPremisesConnectionClient.CreateVirtualEndpointOnPremisesConnection`

```go
ctx := context.TODO()

payload := virtualendpointonpremisesconnection.CloudPCOnPremisesConnection{
	// ...
}


read, err := client.CreateVirtualEndpointOnPremisesConnection(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointOnPremisesConnectionClient.CreateVirtualEndpointOnPremisesConnectionRunHealthCheck`

```go
ctx := context.TODO()
id := virtualendpointonpremisesconnection.NewDeviceManagementVirtualEndpointOnPremisesConnectionID("cloudPCOnPremisesConnectionIdValue")

read, err := client.CreateVirtualEndpointOnPremisesConnectionRunHealthCheck(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointOnPremisesConnectionClient.DeleteVirtualEndpointOnPremisesConnection`

```go
ctx := context.TODO()
id := virtualendpointonpremisesconnection.NewDeviceManagementVirtualEndpointOnPremisesConnectionID("cloudPCOnPremisesConnectionIdValue")

read, err := client.DeleteVirtualEndpointOnPremisesConnection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointOnPremisesConnectionClient.GetVirtualEndpointOnPremisesConnection`

```go
ctx := context.TODO()
id := virtualendpointonpremisesconnection.NewDeviceManagementVirtualEndpointOnPremisesConnectionID("cloudPCOnPremisesConnectionIdValue")

read, err := client.GetVirtualEndpointOnPremisesConnection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointOnPremisesConnectionClient.GetVirtualEndpointOnPremisesConnectionCount`

```go
ctx := context.TODO()


read, err := client.GetVirtualEndpointOnPremisesConnectionCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointOnPremisesConnectionClient.ListVirtualEndpointOnPremisesConnections`

```go
ctx := context.TODO()


// alternatively `client.ListVirtualEndpointOnPremisesConnections(ctx)` can be used to do batched pagination
items, err := client.ListVirtualEndpointOnPremisesConnectionsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VirtualEndpointOnPremisesConnectionClient.UpdateDeviceManagementVirtualEndpointOnPremisesConnectionAdDomainPassword`

```go
ctx := context.TODO()
id := virtualendpointonpremisesconnection.NewDeviceManagementVirtualEndpointOnPremisesConnectionID("cloudPCOnPremisesConnectionIdValue")

payload := virtualendpointonpremisesconnection.UpdateDeviceManagementVirtualEndpointOnPremisesConnectionAdDomainPasswordRequest{
	// ...
}


read, err := client.UpdateDeviceManagementVirtualEndpointOnPremisesConnectionAdDomainPassword(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointOnPremisesConnectionClient.UpdateVirtualEndpointOnPremisesConnection`

```go
ctx := context.TODO()
id := virtualendpointonpremisesconnection.NewDeviceManagementVirtualEndpointOnPremisesConnectionID("cloudPCOnPremisesConnectionIdValue")

payload := virtualendpointonpremisesconnection.CloudPCOnPremisesConnection{
	// ...
}


read, err := client.UpdateVirtualEndpointOnPremisesConnection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
