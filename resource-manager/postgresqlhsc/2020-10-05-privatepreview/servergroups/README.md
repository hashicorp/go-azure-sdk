
## `github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2020-10-05-privatepreview/servergroups` Documentation

The `servergroups` SDK allows for interaction with the Azure Resource Manager Service `postgresqlhsc` (API Version `2020-10-05-privatepreview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2020-10-05-privatepreview/servergroups"
```


### Client Initialization

```go
client := servergroups.NewServerGroupsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `ServerGroupsClient.CheckNameAvailability`

```go
ctx := context.TODO()
id := servergroups.NewSubscriptionID()

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
id := servergroups.NewServerGroupsv2ID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverGroupValue")

payload := servergroups.ServerGroup{
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


### Example Usage: `ServerGroupsClient.Delete`

```go
ctx := context.TODO()
id := servergroups.NewServerGroupsv2ID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverGroupValue")
future, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `ServerGroupsClient.Get`

```go
ctx := context.TODO()
id := servergroups.NewServerGroupsv2ID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverGroupValue")
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
id := servergroups.NewSubscriptionID()
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
id := servergroups.NewResourceGroupID()
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
id := servergroups.NewServerGroupsv2ID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverGroupValue")

payload := servergroups.ServerGroupForUpdate{
	// ...
}

future, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```
