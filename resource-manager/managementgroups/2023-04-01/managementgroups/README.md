
## `github.com/hashicorp/go-azure-sdk/resource-manager/managementgroups/2023-04-01/managementgroups` Documentation

The `managementgroups` SDK allows for interaction with the Azure Resource Manager Service `managementgroups` (API Version `2023-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/managementgroups/2023-04-01/managementgroups"
```


### Client Initialization

```go
client := managementgroups.NewManagementGroupsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagementGroupsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := managementgroups.NewManagementGroupID("groupIdValue")

payload := managementgroups.CreateManagementGroupRequest{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload, managementgroups.DefaultCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `ManagementGroupsClient.Delete`

```go
ctx := context.TODO()
id := managementgroups.NewManagementGroupID("groupIdValue")

if err := client.DeleteThenPoll(ctx, id, managementgroups.DefaultDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `ManagementGroupsClient.Get`

```go
ctx := context.TODO()
id := managementgroups.NewManagementGroupID("groupIdValue")

read, err := client.Get(ctx, id, managementgroups.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagementGroupsClient.GetDescendants`

```go
ctx := context.TODO()
id := managementgroups.NewManagementGroupID("groupIdValue")

// alternatively `client.GetDescendants(ctx, id, managementgroups.DefaultGetDescendantsOperationOptions())` can be used to do batched pagination
items, err := client.GetDescendantsComplete(ctx, id, managementgroups.DefaultGetDescendantsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagementGroupsClient.HierarchySettingsCreateOrUpdate`

```go
ctx := context.TODO()
id := managementgroups.NewManagementGroupID("groupIdValue")

payload := managementgroups.CreateOrUpdateSettingsRequest{
	// ...
}


read, err := client.HierarchySettingsCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagementGroupsClient.HierarchySettingsDelete`

```go
ctx := context.TODO()
id := managementgroups.NewManagementGroupID("groupIdValue")

read, err := client.HierarchySettingsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagementGroupsClient.HierarchySettingsGet`

```go
ctx := context.TODO()
id := managementgroups.NewManagementGroupID("groupIdValue")

read, err := client.HierarchySettingsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagementGroupsClient.HierarchySettingsList`

```go
ctx := context.TODO()
id := managementgroups.NewManagementGroupID("groupIdValue")

read, err := client.HierarchySettingsList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagementGroupsClient.HierarchySettingsUpdate`

```go
ctx := context.TODO()
id := managementgroups.NewManagementGroupID("groupIdValue")

payload := managementgroups.CreateOrUpdateSettingsRequest{
	// ...
}


read, err := client.HierarchySettingsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagementGroupsClient.List`

```go
ctx := context.TODO()


// alternatively `client.List(ctx, managementgroups.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, managementgroups.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagementGroupsClient.UbscriptionsCreate`

```go
ctx := context.TODO()
id := managementgroups.NewSubscriptionID("groupIdValue", "12345678-1234-9876-4563-123456789012")

read, err := client.UbscriptionsCreate(ctx, id, managementgroups.DefaultUbscriptionsCreateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagementGroupsClient.UbscriptionsDelete`

```go
ctx := context.TODO()
id := managementgroups.NewSubscriptionID("groupIdValue", "12345678-1234-9876-4563-123456789012")

read, err := client.UbscriptionsDelete(ctx, id, managementgroups.DefaultUbscriptionsDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagementGroupsClient.UbscriptionsGetSubscription`

```go
ctx := context.TODO()
id := managementgroups.NewSubscriptionID("groupIdValue", "12345678-1234-9876-4563-123456789012")

read, err := client.UbscriptionsGetSubscription(ctx, id, managementgroups.DefaultUbscriptionsGetSubscriptionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagementGroupsClient.UbscriptionsGetSubscriptionsUnderManagementGroup`

```go
ctx := context.TODO()
id := managementgroups.NewManagementGroupID("groupIdValue")

// alternatively `client.UbscriptionsGetSubscriptionsUnderManagementGroup(ctx, id)` can be used to do batched pagination
items, err := client.UbscriptionsGetSubscriptionsUnderManagementGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagementGroupsClient.Update`

```go
ctx := context.TODO()
id := managementgroups.NewManagementGroupID("groupIdValue")

payload := managementgroups.PatchManagementGroupRequest{
	// ...
}


read, err := client.Update(ctx, id, payload, managementgroups.DefaultUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
