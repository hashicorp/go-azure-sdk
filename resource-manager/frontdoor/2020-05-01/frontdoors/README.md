
## `github.com/hashicorp/go-azure-sdk/resource-manager/frontdoor/2020-05-01/frontdoors` Documentation

The `frontdoors` SDK allows for interaction with the Azure Resource Manager Service `frontdoor` (API Version `2020-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/frontdoor/2020-05-01/frontdoors"
```


### Client Initialization

```go
client := frontdoors.NewFrontDoorsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `FrontDoorsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := frontdoors.NewFrontDoorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "frontDoorValue")

payload := frontdoors.FrontDoor{
	// ...
}

future, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `FrontDoorsClient.Delete`

```go
ctx := context.TODO()
id := frontdoors.NewFrontDoorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "frontDoorValue")
future, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `FrontDoorsClient.EndpointsPurgeContent`

```go
ctx := context.TODO()
id := frontdoors.NewFrontDoorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "frontDoorValue")

payload := frontdoors.PurgeParameters{
	// ...
}

future, err := client.EndpointsPurgeContent(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `FrontDoorsClient.FrontendEndpointsDisableHttps`

```go
ctx := context.TODO()
id := frontdoors.NewFrontendEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "frontDoorValue", "frontendEndpointValue")
future, err := client.FrontendEndpointsDisableHttps(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `FrontDoorsClient.FrontendEndpointsEnableHttps`

```go
ctx := context.TODO()
id := frontdoors.NewFrontendEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "frontDoorValue", "frontendEndpointValue")

payload := frontdoors.CustomHttpsConfiguration{
	// ...
}

future, err := client.FrontendEndpointsEnableHttps(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `FrontDoorsClient.FrontendEndpointsGet`

```go
ctx := context.TODO()
id := frontdoors.NewFrontendEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "frontDoorValue", "frontendEndpointValue")
read, err := client.FrontendEndpointsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FrontDoorsClient.FrontendEndpointsListByFrontDoor`

```go
ctx := context.TODO()
id := frontdoors.NewFrontDoorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "frontDoorValue")
// alternatively `client.FrontendEndpointsListByFrontDoor(ctx, id)` can be used to do batched pagination
items, err := client.FrontendEndpointsListByFrontDoorComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `FrontDoorsClient.Get`

```go
ctx := context.TODO()
id := frontdoors.NewFrontDoorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "frontDoorValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FrontDoorsClient.List`

```go
ctx := context.TODO()
id := frontdoors.NewSubscriptionID()
// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `FrontDoorsClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := frontdoors.NewResourceGroupID()
// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `FrontDoorsClient.RulesEnginesCreateOrUpdate`

```go
ctx := context.TODO()
id := frontdoors.NewRulesEngineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "frontDoorValue", "rulesEngineValue")

payload := frontdoors.RulesEngine{
	// ...
}

future, err := client.RulesEnginesCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `FrontDoorsClient.RulesEnginesDelete`

```go
ctx := context.TODO()
id := frontdoors.NewRulesEngineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "frontDoorValue", "rulesEngineValue")
future, err := client.RulesEnginesDelete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `FrontDoorsClient.RulesEnginesGet`

```go
ctx := context.TODO()
id := frontdoors.NewRulesEngineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "frontDoorValue", "rulesEngineValue")
read, err := client.RulesEnginesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `FrontDoorsClient.RulesEnginesListByFrontDoor`

```go
ctx := context.TODO()
id := frontdoors.NewFrontDoorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "frontDoorValue")
// alternatively `client.RulesEnginesListByFrontDoor(ctx, id)` can be used to do batched pagination
items, err := client.RulesEnginesListByFrontDoorComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `FrontDoorsClient.ValidateCustomDomain`

```go
ctx := context.TODO()
id := frontdoors.NewFrontDoorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "frontDoorValue")

payload := frontdoors.ValidateCustomDomainInput{
	// ...
}

read, err := client.ValidateCustomDomain(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
