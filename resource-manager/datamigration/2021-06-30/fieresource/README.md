
## `github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/fieresource` Documentation

The `fieresource` SDK allows for interaction with the Azure Resource Manager Service `datamigration` (API Version `2021-06-30`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/fieresource"
```


### Client Initialization

```go
client := fieresource.NewFieResourceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FieResourceClient.FilesCreateOrUpdate`

```go
ctx := context.TODO()
id := fieresource.NewFileID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue", "fileValue")

payload := fieresource.ProjectFile{
	// ...
}


read, err := client.FilesCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
