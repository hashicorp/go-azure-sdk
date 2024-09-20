
## `github.com/hashicorp/go-azure-sdk/resource-manager/alertsmanagement/2019-05-05-preview/smartgroups` Documentation

The `smartgroups` SDK allows for interaction with Azure Resource Manager `alertsmanagement` (API Version `2019-05-05-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/alertsmanagement/2019-05-05-preview/smartgroups"
```


### Client Initialization

```go
client := smartgroups.NewSmartGroupsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SmartGroupsClient.ChangeState`

```go
ctx := context.TODO()
id := smartgroups.NewSmartGroupID("12345678-1234-9876-4563-123456789012", "smartGroupId")

read, err := client.ChangeState(ctx, id, smartgroups.DefaultChangeStateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SmartGroupsClient.GetAll`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.GetAll(ctx, id, smartgroups.DefaultGetAllOperationOptions())` can be used to do batched pagination
items, err := client.GetAllComplete(ctx, id, smartgroups.DefaultGetAllOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SmartGroupsClient.GetById`

```go
ctx := context.TODO()
id := smartgroups.NewSmartGroupID("12345678-1234-9876-4563-123456789012", "smartGroupId")

read, err := client.GetById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SmartGroupsClient.GetHistory`

```go
ctx := context.TODO()
id := smartgroups.NewSmartGroupID("12345678-1234-9876-4563-123456789012", "smartGroupId")

read, err := client.GetHistory(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
