
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2025-08-01/deletedvaults` Documentation

The `deletedvaults` SDK allows for interaction with Azure Resource Manager `recoveryservices` (API Version `2025-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2025-08-01/deletedvaults"
```


### Client Initialization

```go
client := deletedvaults.NewDeletedVaultsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeletedVaultsClient.Get`

```go
ctx := context.TODO()
id := deletedvaults.NewDeletedVaultID("12345678-1234-9876-4563-123456789012", "locationName", "deletedVaultName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedVaultsClient.GetOperationStatus`

```go
ctx := context.TODO()
id := deletedvaults.NewOperationID("12345678-1234-9876-4563-123456789012", "locationName", "deletedVaultName", "operationId")

read, err := client.GetOperationStatus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedVaultsClient.ListBySubscriptionId`

```go
ctx := context.TODO()
id := deletedvaults.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.ListBySubscriptionId(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionIdComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeletedVaultsClient.Undelete`

```go
ctx := context.TODO()
id := deletedvaults.NewDeletedVaultID("12345678-1234-9876-4563-123456789012", "locationName", "deletedVaultName")

payload := deletedvaults.DeletedVaultUndeleteInput{
	// ...
}


if err := client.UndeleteThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
