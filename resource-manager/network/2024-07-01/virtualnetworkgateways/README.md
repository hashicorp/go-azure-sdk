
## `github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/virtualnetworkgateways` Documentation

The `virtualnetworkgateways` SDK allows for interaction with Azure Resource Manager `network` (API Version `2024-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/virtualnetworkgateways"
```


### Client Initialization

```go
client := virtualnetworkgateways.NewVirtualNetworkGatewaysClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VirtualNetworkGatewaysClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

payload := virtualnetworkgateways.VirtualNetworkGateway{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.Delete`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.DisconnectVirtualNetworkGatewayVpnConnections`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

payload := virtualnetworkgateways.P2SVpnConnectionRequest{
	// ...
}


if err := client.DisconnectVirtualNetworkGatewayVpnConnectionsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.GenerateVpnProfile`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

payload := virtualnetworkgateways.VpnClientParameters{
	// ...
}


if err := client.GenerateVpnProfileThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.Generatevpnclientpackage`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

payload := virtualnetworkgateways.VpnClientParameters{
	// ...
}


if err := client.GeneratevpnclientpackageThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.Get`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualNetworkGatewaysClient.GetAdvertisedRoutes`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

if err := client.GetAdvertisedRoutesThenPoll(ctx, id, virtualnetworkgateways.DefaultGetAdvertisedRoutesOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.GetBgpPeerStatus`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

if err := client.GetBgpPeerStatusThenPoll(ctx, id, virtualnetworkgateways.DefaultGetBgpPeerStatusOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.GetFailoverAllTestDetails`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

if err := client.GetFailoverAllTestDetailsThenPoll(ctx, id, virtualnetworkgateways.DefaultGetFailoverAllTestDetailsOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.GetFailoverSingleTestDetails`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

if err := client.GetFailoverSingleTestDetailsThenPoll(ctx, id, virtualnetworkgateways.DefaultGetFailoverSingleTestDetailsOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.GetLearnedRoutes`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

if err := client.GetLearnedRoutesThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.GetResiliencyInformation`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

if err := client.GetResiliencyInformationThenPoll(ctx, id, virtualnetworkgateways.DefaultGetResiliencyInformationOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.GetRoutesInformation`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

if err := client.GetRoutesInformationThenPoll(ctx, id, virtualnetworkgateways.DefaultGetRoutesInformationOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.GetVpnProfilePackageURL`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

if err := client.GetVpnProfilePackageURLThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.GetVpnclientConnectionHealth`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

if err := client.GetVpnclientConnectionHealthThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.GetVpnclientIPsecParameters`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

if err := client.GetVpnclientIPsecParametersThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.List`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VirtualNetworkGatewaysClient.ListConnections`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

// alternatively `client.ListConnections(ctx, id)` can be used to do batched pagination
items, err := client.ListConnectionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VirtualNetworkGatewaysClient.Reset`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

if err := client.ResetThenPoll(ctx, id, virtualnetworkgateways.DefaultResetOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.ResetVpnClientSharedKey`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

if err := client.ResetVpnClientSharedKeyThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.SetVpnclientIPsecParameters`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

payload := virtualnetworkgateways.VpnClientIPsecParameters{
	// ...
}


if err := client.SetVpnclientIPsecParametersThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.StartExpressRouteSiteFailoverSimulation`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

if err := client.StartExpressRouteSiteFailoverSimulationThenPoll(ctx, id, virtualnetworkgateways.DefaultStartExpressRouteSiteFailoverSimulationOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.StartPacketCapture`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

payload := virtualnetworkgateways.VpnPacketCaptureStartParameters{
	// ...
}


if err := client.StartPacketCaptureThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.StopExpressRouteSiteFailoverSimulation`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

payload := virtualnetworkgateways.ExpressRouteFailoverStopApiParameters{
	// ...
}


if err := client.StopExpressRouteSiteFailoverSimulationThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.StopPacketCapture`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

payload := virtualnetworkgateways.VpnPacketCaptureStopParameters{
	// ...
}


if err := client.StopPacketCaptureThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.SupportedVpnDevices`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

read, err := client.SupportedVpnDevices(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualNetworkGatewaysClient.UpdateTags`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

payload := virtualnetworkgateways.TagsObject{
	// ...
}


if err := client.UpdateTagsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.VirtualNetworkGatewayNatRulesCreateOrUpdate`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayNatRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName", "natRuleName")

payload := virtualnetworkgateways.VirtualNetworkGatewayNatRule{
	// ...
}


if err := client.VirtualNetworkGatewayNatRulesCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.VirtualNetworkGatewayNatRulesDelete`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayNatRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName", "natRuleName")

if err := client.VirtualNetworkGatewayNatRulesDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.VirtualNetworkGatewayNatRulesGet`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayNatRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName", "natRuleName")

read, err := client.VirtualNetworkGatewayNatRulesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualNetworkGatewaysClient.VirtualNetworkGatewayNatRulesListByVirtualNetworkGateway`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

// alternatively `client.VirtualNetworkGatewayNatRulesListByVirtualNetworkGateway(ctx, id)` can be used to do batched pagination
items, err := client.VirtualNetworkGatewayNatRulesListByVirtualNetworkGatewayComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VirtualNetworkGatewaysClient.VirtualNetworkGatewaysInvokeAbortMigration`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

if err := client.VirtualNetworkGatewaysInvokeAbortMigrationThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.VirtualNetworkGatewaysInvokeCommitMigration`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

if err := client.VirtualNetworkGatewaysInvokeCommitMigrationThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.VirtualNetworkGatewaysInvokeExecuteMigration`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

if err := client.VirtualNetworkGatewaysInvokeExecuteMigrationThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.VirtualNetworkGatewaysInvokePrepareMigration`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewVirtualNetworkGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualNetworkGatewayName")

payload := virtualnetworkgateways.VirtualNetworkGatewayMigrationParameters{
	// ...
}


if err := client.VirtualNetworkGatewaysInvokePrepareMigrationThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `VirtualNetworkGatewaysClient.VpnDeviceConfigurationScript`

```go
ctx := context.TODO()
id := virtualnetworkgateways.NewConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectionName")

payload := virtualnetworkgateways.VpnDeviceScriptParameters{
	// ...
}


read, err := client.VpnDeviceConfigurationScript(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
