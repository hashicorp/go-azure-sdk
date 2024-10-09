
## `github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/armtemplates` Documentation

The `armtemplates` SDK allows for interaction with Azure Resource Manager `devtestlab` (API Version `2018-09-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/armtemplates"
```


### Client Initialization

```go
client := armtemplates.NewArmTemplatesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ArmTemplatesClient.Get`

```go
ctx := context.TODO()
id := armtemplates.NewArmTemplateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "artifactSourceName", "armTemplateName")

read, err := client.Get(ctx, id, armtemplates.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ArmTemplatesClient.List`

```go
ctx := context.TODO()
id := armtemplates.NewArtifactSourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "artifactSourceName")

// alternatively `client.List(ctx, id, armtemplates.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, armtemplates.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
