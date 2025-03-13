
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/hunts` Documentation

The `hunts` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2023-12-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/hunts"
```


### Client Initialization

```go
client := hunts.NewHuntsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `HuntsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := hunts.NewHuntID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "huntId")

payload := hunts.Hunt{
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


### Example Usage: `HuntsClient.Delete`

```go
ctx := context.TODO()
id := hunts.NewHuntID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "huntId")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HuntsClient.Get`

```go
ctx := context.TODO()
id := hunts.NewHuntID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "huntId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HuntsClient.List`

```go
ctx := context.TODO()
id := hunts.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.List(ctx, id, hunts.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, hunts.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
