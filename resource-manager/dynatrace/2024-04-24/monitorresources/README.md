
## `github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2024-04-24/monitorresources` Documentation

The `monitorresources` SDK allows for interaction with Azure Resource Manager `dynatrace` (API Version `2024-04-24`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2024-04-24/monitorresources"
```


### Client Initialization

```go
client := monitorresources.NewMonitorResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MonitorResourcesClient.MonitorsCreateOrUpdate`

```go
ctx := context.TODO()
id := monitorresources.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

payload := monitorresources.MonitorResource{
	// ...
}


if err := client.MonitorsCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `MonitorResourcesClient.MonitorsDelete`

```go
ctx := context.TODO()
id := monitorresources.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

if err := client.MonitorsDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `MonitorResourcesClient.MonitorsGet`

```go
ctx := context.TODO()
id := monitorresources.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

read, err := client.MonitorsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MonitorResourcesClient.MonitorsGetMetricStatus`

```go
ctx := context.TODO()
id := monitorresources.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

payload := monitorresources.MetricStatusRequest{
	// ...
}


read, err := client.MonitorsGetMetricStatus(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MonitorResourcesClient.MonitorsGetSSODetails`

```go
ctx := context.TODO()
id := monitorresources.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

payload := monitorresources.SSODetailsRequest{
	// ...
}


read, err := client.MonitorsGetSSODetails(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MonitorResourcesClient.MonitorsGetVMHostPayload`

```go
ctx := context.TODO()
id := monitorresources.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

read, err := client.MonitorsGetVMHostPayload(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MonitorResourcesClient.MonitorsListAppServices`

```go
ctx := context.TODO()
id := monitorresources.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

// alternatively `client.MonitorsListAppServices(ctx, id)` can be used to do batched pagination
items, err := client.MonitorsListAppServicesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MonitorResourcesClient.MonitorsListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.MonitorsListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.MonitorsListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MonitorResourcesClient.MonitorsListBySubscriptionId`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.MonitorsListBySubscriptionId(ctx, id)` can be used to do batched pagination
items, err := client.MonitorsListBySubscriptionIdComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MonitorResourcesClient.MonitorsListHosts`

```go
ctx := context.TODO()
id := monitorresources.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

// alternatively `client.MonitorsListHosts(ctx, id)` can be used to do batched pagination
items, err := client.MonitorsListHostsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MonitorResourcesClient.MonitorsListLinkableEnvironments`

```go
ctx := context.TODO()
id := monitorresources.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

payload := monitorresources.LinkableEnvironmentRequest{
	// ...
}


// alternatively `client.MonitorsListLinkableEnvironments(ctx, id, payload)` can be used to do batched pagination
items, err := client.MonitorsListLinkableEnvironmentsComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MonitorResourcesClient.MonitorsListMonitoredResources`

```go
ctx := context.TODO()
id := monitorresources.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

payload := monitorresources.LogStatusRequest{
	// ...
}


// alternatively `client.MonitorsListMonitoredResources(ctx, id, payload)` can be used to do batched pagination
items, err := client.MonitorsListMonitoredResourcesComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MonitorResourcesClient.MonitorsManageAgentInstallation`

```go
ctx := context.TODO()
id := monitorresources.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

payload := monitorresources.ManageAgentInstallationRequest{
	// ...
}


read, err := client.MonitorsManageAgentInstallation(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MonitorResourcesClient.MonitorsUpdate`

```go
ctx := context.TODO()
id := monitorresources.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

payload := monitorresources.MonitorResourceUpdate{
	// ...
}


read, err := client.MonitorsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MonitorResourcesClient.MonitorsUpgradePlan`

```go
ctx := context.TODO()
id := monitorresources.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

payload := monitorresources.UpgradePlanRequest{
	// ...
}


if err := client.MonitorsUpgradePlanThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
