
## `github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2024-08-01/customoperation` Documentation

The `customoperation` SDK allows for interaction with Azure Resource Manager `postgresql` (API Version `2024-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2024-08-01/customoperation"
```


### Client Initialization

```go
client := customoperation.NewCustomOperationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CustomOperationClient.CheckMigrationNameAvailability`

```go
ctx := context.TODO()
id := customoperation.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName")

payload := customoperation.MigrationNameAvailabilityResource{
	// ...
}


read, err := client.CheckMigrationNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
