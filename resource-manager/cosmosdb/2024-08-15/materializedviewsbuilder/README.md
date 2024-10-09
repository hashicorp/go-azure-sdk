
## `github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2024-08-15/materializedviewsbuilder` Documentation

The `materializedviewsbuilder` SDK allows for interaction with Azure Resource Manager `cosmosdb` (API Version `2024-08-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2024-08-15/materializedviewsbuilder"
```


### Client Initialization

```go
client := materializedviewsbuilder.NewMaterializedViewsBuilderClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MaterializedViewsBuilderClient.ServiceCreate`

```go
ctx := context.TODO()
id := materializedviewsbuilder.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "databaseAccountName", "serviceName")

payload := materializedviewsbuilder.ServiceResourceCreateUpdateParameters{
	// ...
}


if err := client.ServiceCreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `MaterializedViewsBuilderClient.ServiceDelete`

```go
ctx := context.TODO()
id := materializedviewsbuilder.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "databaseAccountName", "serviceName")

if err := client.ServiceDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `MaterializedViewsBuilderClient.ServiceGet`

```go
ctx := context.TODO()
id := materializedviewsbuilder.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "databaseAccountName", "serviceName")

read, err := client.ServiceGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
