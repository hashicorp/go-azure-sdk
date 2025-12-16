
## `github.com/hashicorp/go-azure-sdk/resource-manager/managementgroups/2023-04-01/managements` Documentation

The `managements` SDK allows for interaction with Azure Resource Manager `managementgroups` (API Version `2023-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/managementgroups/2023-04-01/managements"
```


### Client Initialization

```go
client := managements.NewManagementsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagementsClient.CheckNameAvailability`

```go
ctx := context.TODO()

payload := managements.CheckNameAvailabilityRequest{
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


### Example Usage: `ManagementsClient.EntitiesList`

```go
ctx := context.TODO()


// alternatively `client.EntitiesList(ctx, managements.DefaultEntitiesListOperationOptions())` can be used to do batched pagination
items, err := client.EntitiesListComplete(ctx, managements.DefaultEntitiesListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagementsClient.ManagementGroupsList`

```go
ctx := context.TODO()


// alternatively `client.ManagementGroupsList(ctx, managements.DefaultManagementGroupsListOperationOptions())` can be used to do batched pagination
items, err := client.ManagementGroupsListComplete(ctx, managements.DefaultManagementGroupsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagementsClient.StartTenantBackfill`

```go
ctx := context.TODO()


read, err := client.StartTenantBackfill(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagementsClient.TenantBackfillStatus`

```go
ctx := context.TODO()


read, err := client.TenantBackfillStatus(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
