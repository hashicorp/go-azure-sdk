
## `github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/openapis` Documentation

The `openapis` SDK allows for interaction with Azure Resource Manager `cdn` (API Version `2025-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/openapis"
```


### Client Initialization

```go
client := openapis.NewOpenapisClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OpenapisClient.CheckEndpointNameAvailability`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

payload := openapis.CheckEndpointNameAvailabilityInput{
	// ...
}


read, err := client.CheckEndpointNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.CheckNameAvailability`

```go
ctx := context.TODO()

payload := openapis.CheckNameAvailabilityInput{
	// ...
}


read, err := client.CheckNameAvailability(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.CheckNameAvailabilityWithSubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := openapis.CheckNameAvailabilityInput{
	// ...
}


read, err := client.CheckNameAvailabilityWithSubscription(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.EdgeNodesList`

```go
ctx := context.TODO()


// alternatively `client.EdgeNodesList(ctx)` can be used to do batched pagination
items, err := client.EdgeNodesListComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.ManagedRuleSetsList`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ManagedRuleSetsList(ctx, id)` can be used to do batched pagination
items, err := client.ManagedRuleSetsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.ProfilesCanMigrate`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

payload := openapis.CanMigrateParameters{
	// ...
}


if err := client.ProfilesCanMigrateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `OpenapisClient.ProfilesMigrate`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

payload := openapis.MigrationParameters{
	// ...
}


if err := client.ProfilesMigrateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `OpenapisClient.ResourceUsageList`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ResourceUsageList(ctx, id)` can be used to do batched pagination
items, err := client.ResourceUsageListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.ValidateProbe`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := openapis.ValidateProbeInput{
	// ...
}


read, err := client.ValidateProbe(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
