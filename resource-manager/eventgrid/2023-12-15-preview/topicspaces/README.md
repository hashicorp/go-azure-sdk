
## `github.com/hashicorp/go-azure-sdk/resource-manager/eventgrid/2023-12-15-preview/topicspaces` Documentation

The `topicspaces` SDK allows for interaction with the Azure Resource Manager Service `eventgrid` (API Version `2023-12-15-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/eventgrid/2023-12-15-preview/topicspaces"
```


### Client Initialization

```go
client := topicspaces.NewTopicSpacesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TopicSpacesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := topicspaces.NewTopicSpaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue", "topicSpaceValue")

payload := topicspaces.TopicSpace{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `TopicSpacesClient.Delete`

```go
ctx := context.TODO()
id := topicspaces.NewTopicSpaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue", "topicSpaceValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `TopicSpacesClient.Get`

```go
ctx := context.TODO()
id := topicspaces.NewTopicSpaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue", "topicSpaceValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TopicSpacesClient.ListByNamespace`

```go
ctx := context.TODO()
id := topicspaces.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue")

// alternatively `client.ListByNamespace(ctx, id, topicspaces.DefaultListByNamespaceOperationOptions())` can be used to do batched pagination
items, err := client.ListByNamespaceComplete(ctx, id, topicspaces.DefaultListByNamespaceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
