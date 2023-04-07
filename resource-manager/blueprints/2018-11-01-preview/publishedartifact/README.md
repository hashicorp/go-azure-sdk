
## `github.com/hashicorp/go-azure-sdk/resource-manager/blueprints/2018-11-01-preview/publishedartifact` Documentation

The `publishedartifact` SDK allows for interaction with the Azure Resource Manager Service `blueprints` (API Version `2018-11-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/blueprints/2018-11-01-preview/publishedartifact"
```


### Client Initialization

```go
client := publishedartifact.NewPublishedArtifactClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PublishedArtifactClient.Get`

```go
ctx := context.TODO()
id := publishedartifact.NewVersionArtifactScopedID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "blueprintValue", "versionIdValue", "artifactValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PublishedArtifactClient.List`

```go
ctx := context.TODO()
id := publishedartifact.NewScopedVersionID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "blueprintValue", "versionIdValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
