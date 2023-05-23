
## `github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2022-08-15/restorables` Documentation

The `restorables` SDK allows for interaction with the Azure Resource Manager Service `cosmosdb` (API Version `2022-08-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2022-08-15/restorables"
```


### Client Initialization

```go
client := restorables.NewRestorablesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RestorablesClient.MongoDBResourcesRetrieveContinuousBackupInformation`

```go
ctx := context.TODO()
id := restorables.NewMongodbDatabaseCollectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "databaseAccountValue", "mongodbDatabaseValue", "collectionValue")

payload := restorables.ContinuousBackupRestoreLocation{
	// ...
}


if err := client.MongoDBResourcesRetrieveContinuousBackupInformationThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `RestorablesClient.RestorableDatabaseAccountsGetByLocation`

```go
ctx := context.TODO()
id := restorables.NewRestorableDatabaseAccountID("12345678-1234-9876-4563-123456789012", "locationValue", "instanceIdValue")

read, err := client.RestorableDatabaseAccountsGetByLocation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RestorablesClient.RestorableDatabaseAccountsList`

```go
ctx := context.TODO()
id := restorables.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.RestorableDatabaseAccountsList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RestorablesClient.RestorableDatabaseAccountsListByLocation`

```go
ctx := context.TODO()
id := restorables.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

read, err := client.RestorableDatabaseAccountsListByLocation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RestorablesClient.RestorableMongodbCollectionsList`

```go
ctx := context.TODO()
id := restorables.NewRestorableDatabaseAccountID("12345678-1234-9876-4563-123456789012", "locationValue", "instanceIdValue")

read, err := client.RestorableMongodbCollectionsList(ctx, id, restorables.DefaultRestorableMongodbCollectionsListOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RestorablesClient.RestorableMongodbDatabasesList`

```go
ctx := context.TODO()
id := restorables.NewRestorableDatabaseAccountID("12345678-1234-9876-4563-123456789012", "locationValue", "instanceIdValue")

read, err := client.RestorableMongodbDatabasesList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RestorablesClient.RestorableMongodbResourcesList`

```go
ctx := context.TODO()
id := restorables.NewRestorableDatabaseAccountID("12345678-1234-9876-4563-123456789012", "locationValue", "instanceIdValue")

read, err := client.RestorableMongodbResourcesList(ctx, id, restorables.DefaultRestorableMongodbResourcesListOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RestorablesClient.RestorableSqlContainersList`

```go
ctx := context.TODO()
id := restorables.NewRestorableDatabaseAccountID("12345678-1234-9876-4563-123456789012", "locationValue", "instanceIdValue")

read, err := client.RestorableSqlContainersList(ctx, id, restorables.DefaultRestorableSqlContainersListOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RestorablesClient.RestorableSqlDatabasesList`

```go
ctx := context.TODO()
id := restorables.NewRestorableDatabaseAccountID("12345678-1234-9876-4563-123456789012", "locationValue", "instanceIdValue")

read, err := client.RestorableSqlDatabasesList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RestorablesClient.RestorableSqlResourcesList`

```go
ctx := context.TODO()
id := restorables.NewRestorableDatabaseAccountID("12345678-1234-9876-4563-123456789012", "locationValue", "instanceIdValue")

read, err := client.RestorableSqlResourcesList(ctx, id, restorables.DefaultRestorableSqlResourcesListOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RestorablesClient.SqlResourcesRetrieveContinuousBackupInformation`

```go
ctx := context.TODO()
id := restorables.NewContainerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "databaseAccountValue", "sqlDatabaseValue", "containerValue")

payload := restorables.ContinuousBackupRestoreLocation{
	// ...
}


if err := client.SqlResourcesRetrieveContinuousBackupInformationThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
