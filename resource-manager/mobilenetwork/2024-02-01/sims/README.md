
## `github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2024-02-01/sims` Documentation

The `sims` SDK allows for interaction with Azure Resource Manager `mobilenetwork` (API Version `2024-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2024-02-01/sims"
```


### Client Initialization

```go
client := sims.NewSIMsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SIMsClient.BulkDelete`

```go
ctx := context.TODO()
id := sims.NewSimGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "simGroupName")

payload := sims.SimDeleteList{
	// ...
}


if err := client.BulkDeleteThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SIMsClient.BulkUpload`

```go
ctx := context.TODO()
id := sims.NewSimGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "simGroupName")

payload := sims.SimUploadList{
	// ...
}


if err := client.BulkUploadThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SIMsClient.BulkUploadEncrypted`

```go
ctx := context.TODO()
id := sims.NewSimGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "simGroupName")

payload := sims.EncryptedSimUploadList{
	// ...
}


if err := client.BulkUploadEncryptedThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SIMsClient.ListByGroup`

```go
ctx := context.TODO()
id := sims.NewSimGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "simGroupName")

// alternatively `client.ListByGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
