
## `github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-09-01/privatelinkscopesapis` Documentation

The `privatelinkscopesapis` SDK allows for interaction with Azure Resource Manager `insights` (API Version `2021-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/insights/2021-09-01/privatelinkscopesapis"
```


### Client Initialization

```go
client := privatelinkscopesapis.NewPrivateLinkScopesAPIsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateLinkScopesAPIsClient.PrivateLinkScopesCreateOrUpdate`

```go
ctx := context.TODO()
id := privatelinkscopesapis.NewPrivateLinkScopeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateLinkScopeName")

payload := privatelinkscopesapis.AzureMonitorPrivateLinkScope{
	// ...
}


read, err := client.PrivateLinkScopesCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateLinkScopesAPIsClient.PrivateLinkScopesDelete`

```go
ctx := context.TODO()
id := privatelinkscopesapis.NewPrivateLinkScopeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateLinkScopeName")

if err := client.PrivateLinkScopesDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateLinkScopesAPIsClient.PrivateLinkScopesGet`

```go
ctx := context.TODO()
id := privatelinkscopesapis.NewPrivateLinkScopeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateLinkScopeName")

read, err := client.PrivateLinkScopesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateLinkScopesAPIsClient.PrivateLinkScopesList`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.PrivateLinkScopesList(ctx, id)` can be used to do batched pagination
items, err := client.PrivateLinkScopesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PrivateLinkScopesAPIsClient.PrivateLinkScopesListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.PrivateLinkScopesListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.PrivateLinkScopesListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PrivateLinkScopesAPIsClient.PrivateLinkScopesUpdateTags`

```go
ctx := context.TODO()
id := privatelinkscopesapis.NewPrivateLinkScopeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateLinkScopeName")

payload := privatelinkscopesapis.TagsResource{
	// ...
}


read, err := client.PrivateLinkScopesUpdateTags(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
