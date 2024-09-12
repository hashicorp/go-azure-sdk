
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/microsofttunnelsiteserver` Documentation

The `microsofttunnelsiteserver` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/microsofttunnelsiteserver"
```


### Client Initialization

```go
client := microsofttunnelsiteserver.NewMicrosoftTunnelSiteServerClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MicrosoftTunnelSiteServerClient.CreateMicrosoftTunnelSiteServer`

```go
ctx := context.TODO()
id := microsofttunnelsiteserver.NewDeviceManagementMicrosoftTunnelSiteID("microsoftTunnelSiteIdValue")

payload := microsofttunnelsiteserver.MicrosoftTunnelServer{
	// ...
}


read, err := client.CreateMicrosoftTunnelSiteServer(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MicrosoftTunnelSiteServerClient.CreateMicrosoftTunnelSiteServerGenerateServerLogCollectionRequest`

```go
ctx := context.TODO()
id := microsofttunnelsiteserver.NewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID("microsoftTunnelSiteIdValue", "microsoftTunnelServerIdValue")

payload := microsofttunnelsiteserver.CreateMicrosoftTunnelSiteServerGenerateServerLogCollectionRequestRequest{
	// ...
}


read, err := client.CreateMicrosoftTunnelSiteServerGenerateServerLogCollectionRequest(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MicrosoftTunnelSiteServerClient.CreateMicrosoftTunnelSiteServerLogCollectionRequest`

```go
ctx := context.TODO()
id := microsofttunnelsiteserver.NewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID("microsoftTunnelSiteIdValue", "microsoftTunnelServerIdValue")

payload := microsofttunnelsiteserver.CreateMicrosoftTunnelSiteServerLogCollectionRequestRequest{
	// ...
}


read, err := client.CreateMicrosoftTunnelSiteServerLogCollectionRequest(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MicrosoftTunnelSiteServerClient.DeleteMicrosoftTunnelSiteServer`

```go
ctx := context.TODO()
id := microsofttunnelsiteserver.NewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID("microsoftTunnelSiteIdValue", "microsoftTunnelServerIdValue")

read, err := client.DeleteMicrosoftTunnelSiteServer(ctx, id, microsofttunnelsiteserver.DefaultDeleteMicrosoftTunnelSiteServerOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MicrosoftTunnelSiteServerClient.GetMicrosoftTunnelSiteServer`

```go
ctx := context.TODO()
id := microsofttunnelsiteserver.NewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID("microsoftTunnelSiteIdValue", "microsoftTunnelServerIdValue")

read, err := client.GetMicrosoftTunnelSiteServer(ctx, id, microsofttunnelsiteserver.DefaultGetMicrosoftTunnelSiteServerOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MicrosoftTunnelSiteServerClient.GetMicrosoftTunnelSiteServerHealthMetricTimeSeries`

```go
ctx := context.TODO()
id := microsofttunnelsiteserver.NewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID("microsoftTunnelSiteIdValue", "microsoftTunnelServerIdValue")

payload := microsofttunnelsiteserver.GetMicrosoftTunnelSiteServerHealthMetricTimeSeriesRequest{
	// ...
}


// alternatively `client.GetMicrosoftTunnelSiteServerHealthMetricTimeSeries(ctx, id, payload, microsofttunnelsiteserver.DefaultGetMicrosoftTunnelSiteServerHealthMetricTimeSeriesOperationOptions())` can be used to do batched pagination
items, err := client.GetMicrosoftTunnelSiteServerHealthMetricTimeSeriesComplete(ctx, id, payload, microsofttunnelsiteserver.DefaultGetMicrosoftTunnelSiteServerHealthMetricTimeSeriesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MicrosoftTunnelSiteServerClient.GetMicrosoftTunnelSiteServerHealthMetrics`

```go
ctx := context.TODO()
id := microsofttunnelsiteserver.NewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID("microsoftTunnelSiteIdValue", "microsoftTunnelServerIdValue")

payload := microsofttunnelsiteserver.GetMicrosoftTunnelSiteServerHealthMetricsRequest{
	// ...
}


// alternatively `client.GetMicrosoftTunnelSiteServerHealthMetrics(ctx, id, payload, microsofttunnelsiteserver.DefaultGetMicrosoftTunnelSiteServerHealthMetricsOperationOptions())` can be used to do batched pagination
items, err := client.GetMicrosoftTunnelSiteServerHealthMetricsComplete(ctx, id, payload, microsofttunnelsiteserver.DefaultGetMicrosoftTunnelSiteServerHealthMetricsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MicrosoftTunnelSiteServerClient.GetMicrosoftTunnelSiteServersCount`

```go
ctx := context.TODO()
id := microsofttunnelsiteserver.NewDeviceManagementMicrosoftTunnelSiteID("microsoftTunnelSiteIdValue")

read, err := client.GetMicrosoftTunnelSiteServersCount(ctx, id, microsofttunnelsiteserver.DefaultGetMicrosoftTunnelSiteServersCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MicrosoftTunnelSiteServerClient.ListMicrosoftTunnelSiteServers`

```go
ctx := context.TODO()
id := microsofttunnelsiteserver.NewDeviceManagementMicrosoftTunnelSiteID("microsoftTunnelSiteIdValue")

// alternatively `client.ListMicrosoftTunnelSiteServers(ctx, id, microsofttunnelsiteserver.DefaultListMicrosoftTunnelSiteServersOperationOptions())` can be used to do batched pagination
items, err := client.ListMicrosoftTunnelSiteServersComplete(ctx, id, microsofttunnelsiteserver.DefaultListMicrosoftTunnelSiteServersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MicrosoftTunnelSiteServerClient.UpdateMicrosoftTunnelSiteServer`

```go
ctx := context.TODO()
id := microsofttunnelsiteserver.NewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID("microsoftTunnelSiteIdValue", "microsoftTunnelServerIdValue")

payload := microsofttunnelsiteserver.MicrosoftTunnelServer{
	// ...
}


read, err := client.UpdateMicrosoftTunnelSiteServer(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
