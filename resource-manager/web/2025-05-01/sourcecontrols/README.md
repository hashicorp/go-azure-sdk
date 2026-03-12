
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/sourcecontrols` Documentation

The `sourcecontrols` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/sourcecontrols"
```


### Client Initialization

```go
client := sourcecontrols.NewSourceControlsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SourceControlsClient.GetSourceControl`

```go
ctx := context.TODO()
id := sourcecontrols.NewSourceControlID("sourceControlName")

read, err := client.GetSourceControl(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SourceControlsClient.ListSourceControls`

```go
ctx := context.TODO()


// alternatively `client.ListSourceControls(ctx)` can be used to do batched pagination
items, err := client.ListSourceControlsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SourceControlsClient.UpdateSourceControl`

```go
ctx := context.TODO()
id := sourcecontrols.NewSourceControlID("sourceControlName")

payload := sourcecontrols.SourceControl{
	// ...
}


read, err := client.UpdateSourceControl(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
