
## `github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-09-01/scriptcmdlets` Documentation

The `scriptcmdlets` SDK allows for interaction with Azure Resource Manager `vmware` (API Version `2023-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-09-01/scriptcmdlets"
```


### Client Initialization

```go
client := scriptcmdlets.NewScriptCmdletsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ScriptCmdletsClient.Get`

```go
ctx := context.TODO()
id := scriptcmdlets.NewScriptCmdletID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "scriptPackageValue", "scriptCmdletValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ScriptCmdletsClient.List`

```go
ctx := context.TODO()
id := scriptcmdlets.NewScriptPackageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "scriptPackageValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
