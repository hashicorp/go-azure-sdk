
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-11-01/sourcecontrols` Documentation

The `sourcecontrols` SDK allows for interaction with the Azure Resource Manager Service `securityinsights` (API Version `2023-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-11-01/sourcecontrols"
```


### Client Initialization

```go
client := sourcecontrols.NewSourceControlsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SourceControlsClient.Create`

```go
ctx := context.TODO()
id := sourcecontrols.NewSourceControlID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sourceControlIdValue")

payload := sourcecontrols.SourceControl{
	// ...
}


read, err := client.Create(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SourceControlsClient.Delete`

```go
ctx := context.TODO()
id := sourcecontrols.NewSourceControlID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sourceControlIdValue")

payload := sourcecontrols.RepositoryAccessProperties{
	// ...
}


read, err := client.Delete(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SourceControlsClient.Get`

```go
ctx := context.TODO()
id := sourcecontrols.NewSourceControlID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "sourceControlIdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SourceControlsClient.List`

```go
ctx := context.TODO()
id := sourcecontrols.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
