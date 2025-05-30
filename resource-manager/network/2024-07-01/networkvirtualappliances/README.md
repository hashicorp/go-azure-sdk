
## `github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/networkvirtualappliances` Documentation

The `networkvirtualappliances` SDK allows for interaction with Azure Resource Manager `network` (API Version `2024-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/network/2024-07-01/networkvirtualappliances"
```


### Client Initialization

```go
client := networkvirtualappliances.NewNetworkVirtualAppliancesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NetworkVirtualAppliancesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := networkvirtualappliances.NewNetworkVirtualApplianceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "networkVirtualApplianceName")

payload := networkvirtualappliances.NetworkVirtualAppliance{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkVirtualAppliancesClient.Delete`

```go
ctx := context.TODO()
id := networkvirtualappliances.NewNetworkVirtualApplianceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "networkVirtualApplianceName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkVirtualAppliancesClient.Get`

```go
ctx := context.TODO()
id := networkvirtualappliances.NewNetworkVirtualApplianceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "networkVirtualApplianceName")

read, err := client.Get(ctx, id, networkvirtualappliances.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkVirtualAppliancesClient.GetBootDiagnosticLogs`

```go
ctx := context.TODO()
id := networkvirtualappliances.NewNetworkVirtualApplianceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "networkVirtualApplianceName")

payload := networkvirtualappliances.NetworkVirtualApplianceBootDiagnosticParameters{
	// ...
}


if err := client.GetBootDiagnosticLogsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkVirtualAppliancesClient.InboundSecurityRuleCreateOrUpdate`

```go
ctx := context.TODO()
id := networkvirtualappliances.NewInboundSecurityRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "networkVirtualApplianceName", "inboundSecurityRuleName")

payload := networkvirtualappliances.InboundSecurityRule{
	// ...
}


if err := client.InboundSecurityRuleCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkVirtualAppliancesClient.InboundSecurityRuleGet`

```go
ctx := context.TODO()
id := networkvirtualappliances.NewInboundSecurityRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "networkVirtualApplianceName", "inboundSecurityRuleName")

read, err := client.InboundSecurityRuleGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkVirtualAppliancesClient.List`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkVirtualAppliancesClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkVirtualAppliancesClient.Reimage`

```go
ctx := context.TODO()
id := networkvirtualappliances.NewNetworkVirtualApplianceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "networkVirtualApplianceName")

payload := networkvirtualappliances.NetworkVirtualApplianceInstanceIds{
	// ...
}


if err := client.ReimageThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkVirtualAppliancesClient.Restart`

```go
ctx := context.TODO()
id := networkvirtualappliances.NewNetworkVirtualApplianceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "networkVirtualApplianceName")

payload := networkvirtualappliances.NetworkVirtualApplianceInstanceIds{
	// ...
}


if err := client.RestartThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkVirtualAppliancesClient.UpdateTags`

```go
ctx := context.TODO()
id := networkvirtualappliances.NewNetworkVirtualApplianceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "networkVirtualApplianceName")

payload := networkvirtualappliances.TagsObject{
	// ...
}


read, err := client.UpdateTags(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
