
## `github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2023-09-15/services` Documentation

The `services` SDK allows for interaction with Azure Resource Manager `cosmosdb` (API Version `2023-09-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2023-09-15/services"
```


### Client Initialization

```go
client := services.NewServicesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServicesClient.ServiceList`

```go
ctx := context.TODO()
id := services.NewDatabaseAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "databaseAccountName")

read, err := client.ServiceList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
