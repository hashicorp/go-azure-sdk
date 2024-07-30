
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/microsofttunnelsitemicrosofttunnelserver` Documentation

The `microsofttunnelsitemicrosofttunnelserver` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/microsofttunnelsitemicrosofttunnelserver"
```


### Client Initialization

```go
client := microsofttunnelsitemicrosofttunnelserver.NewMicrosoftTunnelSiteMicrosoftTunnelServerClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MicrosoftTunnelSiteMicrosoftTunnelServerClient.CreateMicrosoftTunnelSiteMicrosoftTunnelServer`

```go
ctx := context.TODO()
id := microsofttunnelsitemicrosofttunnelserver.NewDeviceManagementMicrosoftTunnelSiteID("microsoftTunnelSiteIdValue")

payload := microsofttunnelsitemicrosofttunnelserver.MicrosoftTunnelServer{
	// ...
}


read, err := client.CreateMicrosoftTunnelSiteMicrosoftTunnelServer(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MicrosoftTunnelSiteMicrosoftTunnelServerClient.CreateMicrosoftTunnelSiteMicrosoftTunnelServerCreateServerLogCollectionRequest`

```go
ctx := context.TODO()
id := microsofttunnelsitemicrosofttunnelserver.NewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID("microsoftTunnelSiteIdValue", "microsoftTunnelServerIdValue")

payload := microsofttunnelsitemicrosofttunnelserver.CreateMicrosoftTunnelSiteMicrosoftTunnelServerCreateServerLogCollectionRequestRequest{
	// ...
}


read, err := client.CreateMicrosoftTunnelSiteMicrosoftTunnelServerCreateServerLogCollectionRequest(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MicrosoftTunnelSiteMicrosoftTunnelServerClient.CreateMicrosoftTunnelSiteMicrosoftTunnelServerGenerateServerLogCollectionRequest`

```go
ctx := context.TODO()
id := microsofttunnelsitemicrosofttunnelserver.NewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID("microsoftTunnelSiteIdValue", "microsoftTunnelServerIdValue")

payload := microsofttunnelsitemicrosofttunnelserver.CreateMicrosoftTunnelSiteMicrosoftTunnelServerGenerateServerLogCollectionRequestRequest{
	// ...
}


read, err := client.CreateMicrosoftTunnelSiteMicrosoftTunnelServerGenerateServerLogCollectionRequest(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MicrosoftTunnelSiteMicrosoftTunnelServerClient.DeleteMicrosoftTunnelSiteMicrosoftTunnelServer`

```go
ctx := context.TODO()
id := microsofttunnelsitemicrosofttunnelserver.NewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID("microsoftTunnelSiteIdValue", "microsoftTunnelServerIdValue")

read, err := client.DeleteMicrosoftTunnelSiteMicrosoftTunnelServer(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MicrosoftTunnelSiteMicrosoftTunnelServerClient.GetDeviceManagementMicrosoftTunnelSiteMicrosoftTunnelServerHealthMetric`

```go
ctx := context.TODO()
id := microsofttunnelsitemicrosofttunnelserver.NewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID("microsoftTunnelSiteIdValue", "microsoftTunnelServerIdValue")

payload := microsofttunnelsitemicrosofttunnelserver.GetDeviceManagementMicrosoftTunnelSiteMicrosoftTunnelServerHealthMetricRequest{
	// ...
}


read, err := client.GetDeviceManagementMicrosoftTunnelSiteMicrosoftTunnelServerHealthMetric(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MicrosoftTunnelSiteMicrosoftTunnelServerClient.GetDeviceManagementMicrosoftTunnelSiteMicrosoftTunnelServerHealthMetricTimeSery`

```go
ctx := context.TODO()
id := microsofttunnelsitemicrosofttunnelserver.NewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID("microsoftTunnelSiteIdValue", "microsoftTunnelServerIdValue")

payload := microsofttunnelsitemicrosofttunnelserver.GetDeviceManagementMicrosoftTunnelSiteMicrosoftTunnelServerHealthMetricTimeSeryRequest{
	// ...
}


read, err := client.GetDeviceManagementMicrosoftTunnelSiteMicrosoftTunnelServerHealthMetricTimeSery(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MicrosoftTunnelSiteMicrosoftTunnelServerClient.GetMicrosoftTunnelSiteMicrosoftTunnelServer`

```go
ctx := context.TODO()
id := microsofttunnelsitemicrosofttunnelserver.NewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID("microsoftTunnelSiteIdValue", "microsoftTunnelServerIdValue")

read, err := client.GetMicrosoftTunnelSiteMicrosoftTunnelServer(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MicrosoftTunnelSiteMicrosoftTunnelServerClient.GetMicrosoftTunnelSiteMicrosoftTunnelServerCount`

```go
ctx := context.TODO()
id := microsofttunnelsitemicrosofttunnelserver.NewDeviceManagementMicrosoftTunnelSiteID("microsoftTunnelSiteIdValue")

read, err := client.GetMicrosoftTunnelSiteMicrosoftTunnelServerCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MicrosoftTunnelSiteMicrosoftTunnelServerClient.ListMicrosoftTunnelSiteMicrosoftTunnelServers`

```go
ctx := context.TODO()
id := microsofttunnelsitemicrosofttunnelserver.NewDeviceManagementMicrosoftTunnelSiteID("microsoftTunnelSiteIdValue")

// alternatively `client.ListMicrosoftTunnelSiteMicrosoftTunnelServers(ctx, id)` can be used to do batched pagination
items, err := client.ListMicrosoftTunnelSiteMicrosoftTunnelServersComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MicrosoftTunnelSiteMicrosoftTunnelServerClient.UpdateMicrosoftTunnelSiteMicrosoftTunnelServer`

```go
ctx := context.TODO()
id := microsofttunnelsitemicrosofttunnelserver.NewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID("microsoftTunnelSiteIdValue", "microsoftTunnelServerIdValue")

payload := microsofttunnelsitemicrosofttunnelserver.MicrosoftTunnelServer{
	// ...
}


read, err := client.UpdateMicrosoftTunnelSiteMicrosoftTunnelServer(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
