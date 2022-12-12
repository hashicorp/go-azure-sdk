
## `github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/sims` Documentation

The `sims` SDK allows for interaction with the Azure Resource Manager Service `mobilenetwork` (API Version `2022-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mobilenetwork/2022-11-01/sims"
```


### Client Initialization

```go
client := sims.NewSIMsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SIMsClient.ListByGroup`

```go
ctx := context.TODO()
id := sims.NewSimGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "simGroupValue")

// alternatively `client.ListByGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SIMsClient.SimBulkDelete`

```go
ctx := context.TODO()
id := sims.NewSimGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "simGroupValue")

payload := sims.SimDeleteList{
	// ...
}


if err := client.SimBulkDeleteThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SIMsClient.SimBulkUpload`

```go
ctx := context.TODO()
id := sims.NewSimGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "simGroupValue")

payload := sims.SimUploadList{
	// ...
}


if err := client.SimBulkUploadThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SIMsClient.SimBulkUploadEncrypted`

```go
ctx := context.TODO()
id := sims.NewSimGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "simGroupValue")

payload := sims.EncryptedSimUploadList{
	// ...
}


if err := client.SimBulkUploadEncryptedThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
