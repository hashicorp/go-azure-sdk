
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/distributedavailabilitygroups` Documentation

The `distributedavailabilitygroups` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/distributedavailabilitygroups"
```


### Client Initialization

```go
client := distributedavailabilitygroups.NewDistributedAvailabilityGroupsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DistributedAvailabilityGroupsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := distributedavailabilitygroups.NewDistributedAvailabilityGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "distributedAvailabilityGroupName")

payload := distributedavailabilitygroups.DistributedAvailabilityGroup{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DistributedAvailabilityGroupsClient.Delete`

```go
ctx := context.TODO()
id := distributedavailabilitygroups.NewDistributedAvailabilityGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "distributedAvailabilityGroupName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DistributedAvailabilityGroupsClient.Failover`

```go
ctx := context.TODO()
id := distributedavailabilitygroups.NewDistributedAvailabilityGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "distributedAvailabilityGroupName")

payload := distributedavailabilitygroups.DistributedAvailabilityGroupsFailoverRequest{
	// ...
}


if err := client.FailoverThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DistributedAvailabilityGroupsClient.Get`

```go
ctx := context.TODO()
id := distributedavailabilitygroups.NewDistributedAvailabilityGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "distributedAvailabilityGroupName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DistributedAvailabilityGroupsClient.ListByInstance`

```go
ctx := context.TODO()
id := commonids.NewSqlManagedInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName")

// alternatively `client.ListByInstance(ctx, id)` can be used to do batched pagination
items, err := client.ListByInstanceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DistributedAvailabilityGroupsClient.SetRole`

```go
ctx := context.TODO()
id := distributedavailabilitygroups.NewDistributedAvailabilityGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "distributedAvailabilityGroupName")

payload := distributedavailabilitygroups.DistributedAvailabilityGroupSetRole{
	// ...
}


if err := client.SetRoleThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DistributedAvailabilityGroupsClient.Update`

```go
ctx := context.TODO()
id := distributedavailabilitygroups.NewDistributedAvailabilityGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceName", "distributedAvailabilityGroupName")

payload := distributedavailabilitygroups.DistributedAvailabilityGroup{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
