
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolstransparentdataencryption` Documentation

The `sqlpoolstransparentdataencryption` SDK allows for interaction with the Azure Resource Manager Service `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolstransparentdataencryption"
```


### Client Initialization

```go
client := sqlpoolstransparentdataencryption.NewSqlPoolsTransparentDataEncryptionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SqlPoolsTransparentDataEncryptionClient.SqlPoolTransparentDataEncryptionsCreateOrUpdate`

```go
ctx := context.TODO()
id := sqlpoolstransparentdataencryption.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue")

payload := sqlpoolstransparentdataencryption.TransparentDataEncryption{
	// ...
}


read, err := client.SqlPoolTransparentDataEncryptionsCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlPoolsTransparentDataEncryptionClient.SqlPoolTransparentDataEncryptionsGet`

```go
ctx := context.TODO()
id := sqlpoolstransparentdataencryption.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue")

read, err := client.SqlPoolTransparentDataEncryptionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SqlPoolsTransparentDataEncryptionClient.SqlPoolTransparentDataEncryptionsList`

```go
ctx := context.TODO()
id := sqlpoolstransparentdataencryption.NewSqlPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sqlPoolValue")

// alternatively `client.SqlPoolTransparentDataEncryptionsList(ctx, id)` can be used to do batched pagination
items, err := client.SqlPoolTransparentDataEncryptionsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
