
## `github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/scriptpackages` Documentation

The `scriptpackages` SDK allows for interaction with Azure Resource Manager `vmware` (API Version `2024-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/scriptpackages"
```


### Client Initialization

```go
client := scriptpackages.NewScriptPackagesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ScriptPackagesClient.Get`

```go
ctx := context.TODO()
id := scriptpackages.NewScriptPackageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "scriptPackageName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ScriptPackagesClient.List`

```go
ctx := context.TODO()
id := scriptpackages.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
