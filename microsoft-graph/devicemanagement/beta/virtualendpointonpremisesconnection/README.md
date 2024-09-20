
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointonpremisesconnection` Documentation

The `virtualendpointonpremisesconnection` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/virtualendpointonpremisesconnection"
```


### Client Initialization

```go
client := virtualendpointonpremisesconnection.NewVirtualEndpointOnPremisesConnectionClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VirtualEndpointOnPremisesConnectionClient.CreateVirtualEndpointOnPremisesConnection`

```go
ctx := context.TODO()

payload := virtualendpointonpremisesconnection.CloudPCOnPremisesConnection{
	// ...
}


read, err := client.CreateVirtualEndpointOnPremisesConnection(ctx, payload, virtualendpointonpremisesconnection.DefaultCreateVirtualEndpointOnPremisesConnectionOperationOptions())
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
id := virtualendpointonpremisesconnection.NewDeviceManagementVirtualEndpointOnPremisesConnectionID("cloudPCOnPremisesConnectionId")

read, err := client.CreateVirtualEndpointOnPremisesConnectionRunHealthCheck(ctx, id, virtualendpointonpremisesconnection.DefaultCreateVirtualEndpointOnPremisesConnectionRunHealthCheckOperationOptions())
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
id := virtualendpointonpremisesconnection.NewDeviceManagementVirtualEndpointOnPremisesConnectionID("cloudPCOnPremisesConnectionId")

read, err := client.DeleteVirtualEndpointOnPremisesConnection(ctx, id, virtualendpointonpremisesconnection.DefaultDeleteVirtualEndpointOnPremisesConnectionOperationOptions())
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
id := virtualendpointonpremisesconnection.NewDeviceManagementVirtualEndpointOnPremisesConnectionID("cloudPCOnPremisesConnectionId")

read, err := client.GetVirtualEndpointOnPremisesConnection(ctx, id, virtualendpointonpremisesconnection.DefaultGetVirtualEndpointOnPremisesConnectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointOnPremisesConnectionClient.GetVirtualEndpointOnPremisesConnectionsCount`

```go
ctx := context.TODO()


read, err := client.GetVirtualEndpointOnPremisesConnectionsCount(ctx, virtualendpointonpremisesconnection.DefaultGetVirtualEndpointOnPremisesConnectionsCountOperationOptions())
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


// alternatively `client.ListVirtualEndpointOnPremisesConnections(ctx, virtualendpointonpremisesconnection.DefaultListVirtualEndpointOnPremisesConnectionsOperationOptions())` can be used to do batched pagination
items, err := client.ListVirtualEndpointOnPremisesConnectionsComplete(ctx, virtualendpointonpremisesconnection.DefaultListVirtualEndpointOnPremisesConnectionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VirtualEndpointOnPremisesConnectionClient.UpdateVirtualEndpointOnPremisesConnection`

```go
ctx := context.TODO()
id := virtualendpointonpremisesconnection.NewDeviceManagementVirtualEndpointOnPremisesConnectionID("cloudPCOnPremisesConnectionId")

payload := virtualendpointonpremisesconnection.CloudPCOnPremisesConnection{
	// ...
}


read, err := client.UpdateVirtualEndpointOnPremisesConnection(ctx, id, payload, virtualendpointonpremisesconnection.DefaultUpdateVirtualEndpointOnPremisesConnectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualEndpointOnPremisesConnectionClient.UpdateVirtualEndpointOnPremisesConnectionAdDomainPassword`

```go
ctx := context.TODO()
id := virtualendpointonpremisesconnection.NewDeviceManagementVirtualEndpointOnPremisesConnectionID("cloudPCOnPremisesConnectionId")

payload := virtualendpointonpremisesconnection.UpdateVirtualEndpointOnPremisesConnectionAdDomainPasswordRequest{
	// ...
}


read, err := client.UpdateVirtualEndpointOnPremisesConnectionAdDomainPassword(ctx, id, payload, virtualendpointonpremisesconnection.DefaultUpdateVirtualEndpointOnPremisesConnectionAdDomainPasswordOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
