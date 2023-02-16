
## `github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/logfiles` Documentation

The `logfiles` SDK allows for interaction with the Azure Resource Manager Service `mariadb` (API Version `2018-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/logfiles"
```


### Client Initialization

```go
client := logfiles.NewLogFilesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LogFilesClient.ListByServer`

```go
ctx := context.TODO()
id := logfiles.NewServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue")

read, err := client.ListByServer(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
