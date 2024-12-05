
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/huntrelations` Documentation

The `huntrelations` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2023-12-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2023-12-01-preview/huntrelations"
```


### Client Initialization

```go
client := huntrelations.NewHuntRelationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `HuntRelationsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := huntrelations.NewRelationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "huntId", "huntRelationId")

payload := huntrelations.HuntRelation{
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


### Example Usage: `HuntRelationsClient.Delete`

```go
ctx := context.TODO()
id := huntrelations.NewRelationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "huntId", "huntRelationId")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HuntRelationsClient.Get`

```go
ctx := context.TODO()
id := huntrelations.NewRelationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "huntId", "huntRelationId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HuntRelationsClient.List`

```go
ctx := context.TODO()
id := huntrelations.NewHuntID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "huntId")

// alternatively `client.List(ctx, id, huntrelations.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, huntrelations.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
