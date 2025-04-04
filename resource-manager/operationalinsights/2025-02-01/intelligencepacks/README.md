
## `github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2025-02-01/intelligencepacks` Documentation

The `intelligencepacks` SDK allows for interaction with Azure Resource Manager `operationalinsights` (API Version `2025-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2025-02-01/intelligencepacks"
```


### Client Initialization

```go
client := intelligencepacks.NewIntelligencePacksClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IntelligencePacksClient.Disable`

```go
ctx := context.TODO()
id := intelligencepacks.NewIntelligencePackID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "intelligencePackName")

read, err := client.Disable(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IntelligencePacksClient.Enable`

```go
ctx := context.TODO()
id := intelligencepacks.NewIntelligencePackID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "intelligencePackName")

read, err := client.Enable(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IntelligencePacksClient.List`

```go
ctx := context.TODO()
id := intelligencepacks.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
