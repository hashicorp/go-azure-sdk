
## `github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/managednetwork` Documentation

The `managednetwork` SDK allows for interaction with Azure Resource Manager `cognitive` (API Version `2025-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-12-01/managednetwork"
```


### Client Initialization

```go
client := managednetwork.NewManagedNetworkClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedNetworkClient.OutboundRuleCreateOrUpdate`

```go
ctx := context.TODO()
id := managednetwork.NewOutboundRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "managedNetworkName", "outboundRuleName")

payload := managednetwork.OutboundRuleBasicResource{
	// ...
}


if err := client.OutboundRuleCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ManagedNetworkClient.OutboundRuleDelete`

```go
ctx := context.TODO()
id := managednetwork.NewOutboundRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "managedNetworkName", "outboundRuleName")

if err := client.OutboundRuleDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ManagedNetworkClient.OutboundRuleGet`

```go
ctx := context.TODO()
id := managednetwork.NewOutboundRuleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "managedNetworkName", "outboundRuleName")

read, err := client.OutboundRuleGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedNetworkClient.OutboundRuleList`

```go
ctx := context.TODO()
id := managednetwork.NewManagedNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "managedNetworkName")

// alternatively `client.OutboundRuleList(ctx, id)` can be used to do batched pagination
items, err := client.OutboundRuleListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedNetworkClient.OutboundRulesPost`

```go
ctx := context.TODO()
id := managednetwork.NewManagedNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "managedNetworkName")

payload := managednetwork.ManagedNetworkSettingsBasicResource{
	// ...
}


// alternatively `client.OutboundRulesPost(ctx, id, payload)` can be used to do batched pagination
items, err := client.OutboundRulesPostComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedNetworkClient.ProvisionsProvisionManagedNetwork`

```go
ctx := context.TODO()
id := managednetwork.NewManagedNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "managedNetworkName")
var payload interface{}

if err := client.ProvisionsProvisionManagedNetworkThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ManagedNetworkClient.SettingsGet`

```go
ctx := context.TODO()
id := managednetwork.NewManagedNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "managedNetworkName")

read, err := client.SettingsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedNetworkClient.SettingsList`

```go
ctx := context.TODO()
id := managednetwork.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName")

// alternatively `client.SettingsList(ctx, id)` can be used to do batched pagination
items, err := client.SettingsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedNetworkClient.SettingsPatch`

```go
ctx := context.TODO()
id := managednetwork.NewManagedNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "managedNetworkName")

payload := managednetwork.ManagedNetworkSettingsPropertiesBasicResource{
	// ...
}


if err := client.SettingsPatchThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ManagedNetworkClient.SettingsPut`

```go
ctx := context.TODO()
id := managednetwork.NewManagedNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "managedNetworkName")

payload := managednetwork.ManagedNetworkSettingsPropertiesBasicResource{
	// ...
}


if err := client.SettingsPutThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
