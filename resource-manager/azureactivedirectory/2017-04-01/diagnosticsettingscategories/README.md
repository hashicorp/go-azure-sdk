
## `github.com/hashicorp/go-azure-sdk/resource-manager/azureactivedirectory/2017-04-01/diagnosticsettingscategories` Documentation

The `diagnosticsettingscategories` SDK allows for interaction with the Azure Resource Manager Service `azureactivedirectory` (API Version `2017-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/azureactivedirectory/2017-04-01/diagnosticsettingscategories"
```


### Client Initialization

```go
client := diagnosticsettingscategories.NewDiagnosticSettingsCategoriesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DiagnosticSettingsCategoriesClient.DiagnosticSettingsCategoryList`

```go
ctx := context.TODO()


read, err := client.DiagnosticSettingsCategoryList(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
