
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


### Example Usage: `MicrosoftTunnelSiteMicrosoftTunnelServerClient.CreateMicrosoftTunnelSiteServer`

```go
ctx := context.TODO()
id := microsofttunnelsitemicrosofttunnelserver.NewDeviceManagementMicrosoftTunnelSiteID("microsoftTunnelSiteIdValue")

payload := microsofttunnelsitemicrosofttunnelserver.MicrosoftTunnelServer{
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


### Example Usage: `MicrosoftTunnelSiteMicrosoftTunnelServerClient.CreateMicrosoftTunnelSiteServerGenerateServerLogCollectionRequest`

```go
ctx := context.TODO()
id := microsofttunnelsitemicrosofttunnelserver.NewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID("microsoftTunnelSiteIdValue", "microsoftTunnelServerIdValue")

payload := microsofttunnelsitemicrosofttunnelserver.CreateMicrosoftTunnelSiteServerGenerateServerLogCollectionRequestRequest{
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


### Example Usage: `MicrosoftTunnelSiteMicrosoftTunnelServerClient.CreateMicrosoftTunnelSiteServerLogCollectionRequest`

```go
ctx := context.TODO()
id := microsofttunnelsitemicrosofttunnelserver.NewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID("microsoftTunnelSiteIdValue", "microsoftTunnelServerIdValue")

payload := microsofttunnelsitemicrosofttunnelserver.CreateMicrosoftTunnelSiteServerLogCollectionRequestRequest{
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


### Example Usage: `MicrosoftTunnelSiteMicrosoftTunnelServerClient.DeleteMicrosoftTunnelSiteServer`

```go
ctx := context.TODO()
id := microsofttunnelsitemicrosofttunnelserver.NewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID("microsoftTunnelSiteIdValue", "microsoftTunnelServerIdValue")

read, err := client.DeleteMicrosoftTunnelSiteServer(ctx, id, microsofttunnelsitemicrosofttunnelserver.DefaultDeleteMicrosoftTunnelSiteServerOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MicrosoftTunnelSiteMicrosoftTunnelServerClient.GetMicrosoftTunnelSiteServer`

```go
ctx := context.TODO()
id := microsofttunnelsitemicrosofttunnelserver.NewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID("microsoftTunnelSiteIdValue", "microsoftTunnelServerIdValue")

read, err := client.GetMicrosoftTunnelSiteServer(ctx, id, microsofttunnelsitemicrosofttunnelserver.DefaultGetMicrosoftTunnelSiteServerOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MicrosoftTunnelSiteMicrosoftTunnelServerClient.GetMicrosoftTunnelSiteServerHealthMetricTimeSeries`

```go
ctx := context.TODO()
id := microsofttunnelsitemicrosofttunnelserver.NewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID("microsoftTunnelSiteIdValue", "microsoftTunnelServerIdValue")

payload := microsofttunnelsitemicrosofttunnelserver.GetMicrosoftTunnelSiteServerHealthMetricTimeSeriesRequest{
	// ...
}


// alternatively `client.GetMicrosoftTunnelSiteServerHealthMetricTimeSeries(ctx, id, payload, microsofttunnelsitemicrosofttunnelserver.DefaultGetMicrosoftTunnelSiteServerHealthMetricTimeSeriesOperationOptions())` can be used to do batched pagination
items, err := client.GetMicrosoftTunnelSiteServerHealthMetricTimeSeriesComplete(ctx, id, payload, microsofttunnelsitemicrosofttunnelserver.DefaultGetMicrosoftTunnelSiteServerHealthMetricTimeSeriesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MicrosoftTunnelSiteMicrosoftTunnelServerClient.GetMicrosoftTunnelSiteServerHealthMetrics`

```go
ctx := context.TODO()
id := microsofttunnelsitemicrosofttunnelserver.NewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID("microsoftTunnelSiteIdValue", "microsoftTunnelServerIdValue")

payload := microsofttunnelsitemicrosofttunnelserver.GetMicrosoftTunnelSiteServerHealthMetricsRequest{
	// ...
}


// alternatively `client.GetMicrosoftTunnelSiteServerHealthMetrics(ctx, id, payload, microsofttunnelsitemicrosofttunnelserver.DefaultGetMicrosoftTunnelSiteServerHealthMetricsOperationOptions())` can be used to do batched pagination
items, err := client.GetMicrosoftTunnelSiteServerHealthMetricsComplete(ctx, id, payload, microsofttunnelsitemicrosofttunnelserver.DefaultGetMicrosoftTunnelSiteServerHealthMetricsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MicrosoftTunnelSiteMicrosoftTunnelServerClient.GetMicrosoftTunnelSiteServersCount`

```go
ctx := context.TODO()
id := microsofttunnelsitemicrosofttunnelserver.NewDeviceManagementMicrosoftTunnelSiteID("microsoftTunnelSiteIdValue")

read, err := client.GetMicrosoftTunnelSiteServersCount(ctx, id, microsofttunnelsitemicrosofttunnelserver.DefaultGetMicrosoftTunnelSiteServersCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MicrosoftTunnelSiteMicrosoftTunnelServerClient.ListMicrosoftTunnelSiteServers`

```go
ctx := context.TODO()
id := microsofttunnelsitemicrosofttunnelserver.NewDeviceManagementMicrosoftTunnelSiteID("microsoftTunnelSiteIdValue")

// alternatively `client.ListMicrosoftTunnelSiteServers(ctx, id, microsofttunnelsitemicrosofttunnelserver.DefaultListMicrosoftTunnelSiteServersOperationOptions())` can be used to do batched pagination
items, err := client.ListMicrosoftTunnelSiteServersComplete(ctx, id, microsofttunnelsitemicrosofttunnelserver.DefaultListMicrosoftTunnelSiteServersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MicrosoftTunnelSiteMicrosoftTunnelServerClient.UpdateMicrosoftTunnelSiteServer`

```go
ctx := context.TODO()
id := microsofttunnelsitemicrosofttunnelserver.NewDeviceManagementMicrosoftTunnelSiteIdMicrosoftTunnelServerID("microsoftTunnelSiteIdValue", "microsoftTunnelServerIdValue")

payload := microsofttunnelsitemicrosofttunnelserver.MicrosoftTunnelServer{
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
