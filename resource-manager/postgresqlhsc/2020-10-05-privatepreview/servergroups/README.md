
## `github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2020-10-05-privatepreview/servergroups` Documentation

The `servergroups` SDK allows for interaction with the Azure Resource Manager Service `postgresqlhsc` (API Version `2020-10-05-privatepreview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2020-10-05-privatepreview/servergroups"
```


### Client Initialization

```go
client := servergroups.NewServerGroupsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServerGroupsClient.CheckNameAvailability`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := servergroups.NameAvailabilityRequest{
	// ...
}


read, err := client.CheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServerGroupsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := servergroups.NewServerGroupsv2ID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverGroupsv2Value")

payload := servergroups.ServerGroup{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ServerGroupsClient.Delete`

```go
ctx := context.TODO()
id := servergroups.NewServerGroupsv2ID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverGroupsv2Value")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ServerGroupsClient.Get`

```go
ctx := context.TODO()
id := servergroups.NewServerGroupsv2ID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverGroupsv2Value")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServerGroupsClient.List`

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


### Example Usage: `ServerGroupsClient.ListByResourceGroup`

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


### Example Usage: `ServerGroupsClient.Update`

```go
ctx := context.TODO()
id := servergroups.NewServerGroupsv2ID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverGroupsv2Value")

payload := servergroups.ServerGroupForUpdate{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
