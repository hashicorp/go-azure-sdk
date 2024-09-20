
## `github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/artifactsources` Documentation

The `artifactsources` SDK allows for interaction with Azure Resource Manager `devtestlab` (API Version `2018-09-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/artifactsources"
```


### Client Initialization

```go
client := artifactsources.NewArtifactSourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ArtifactSourcesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := artifactsources.NewArtifactSourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "artifactSourceName")

payload := artifactsources.ArtifactSource{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ArtifactSourcesClient.Delete`

```go
ctx := context.TODO()
id := artifactsources.NewArtifactSourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "artifactSourceName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ArtifactSourcesClient.Get`

```go
ctx := context.TODO()
id := artifactsources.NewArtifactSourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "artifactSourceName")

read, err := client.Get(ctx, id, artifactsources.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ArtifactSourcesClient.List`

```go
ctx := context.TODO()
id := artifactsources.NewLabID("12345678-1234-9876-4563-123456789012", "example-resource-group", "name")

// alternatively `client.List(ctx, id, artifactsources.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, artifactsources.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ArtifactSourcesClient.Update`

```go
ctx := context.TODO()
id := artifactsources.NewArtifactSourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "artifactSourceName")

payload := artifactsources.UpdateResource{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
