
## `github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/artifacts` Documentation

The `artifacts` SDK allows for interaction with the Azure Resource Manager Service `devtestlab` (API Version `2018-09-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/artifacts"
```


### Client Initialization

```go
client := artifacts.NewArtifactsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ArtifactsClient.GenerateArmTemplate`

```go
ctx := context.TODO()
id := artifacts.NewArtifactID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "artifactSourceValue", "artifactValue")

payload := artifacts.GenerateArmTemplateRequest{
	// ...
}


read, err := client.GenerateArmTemplate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ArtifactsClient.Get`

```go
ctx := context.TODO()
id := artifacts.NewArtifactID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "artifactSourceValue", "artifactValue")

read, err := client.Get(ctx, id, artifacts.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ArtifactsClient.List`

```go
ctx := context.TODO()
id := artifacts.NewArtifactSourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "artifactSourceValue")

// alternatively `client.List(ctx, id, artifacts.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, artifacts.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
