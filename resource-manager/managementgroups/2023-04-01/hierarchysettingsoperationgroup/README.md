
## `github.com/hashicorp/go-azure-sdk/resource-manager/managementgroups/2023-04-01/hierarchysettingsoperationgroup` Documentation

The `hierarchysettingsoperationgroup` SDK allows for interaction with Azure Resource Manager `managementgroups` (API Version `2023-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/managementgroups/2023-04-01/hierarchysettingsoperationgroup"
```


### Client Initialization

```go
client := hierarchysettingsoperationgroup.NewHierarchySettingsOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `HierarchySettingsOperationGroupClient.HierarchySettingsCreateOrUpdate`

```go
ctx := context.TODO()
id := commonids.NewManagementGroupID("groupId")

payload := hierarchysettingsoperationgroup.CreateOrUpdateSettingsRequest{
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


### Example Usage: `HierarchySettingsOperationGroupClient.HierarchySettingsDelete`

```go
ctx := context.TODO()
id := commonids.NewManagementGroupID("groupId")

read, err := client.HierarchySettingsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HierarchySettingsOperationGroupClient.HierarchySettingsGet`

```go
ctx := context.TODO()
id := commonids.NewManagementGroupID("groupId")

read, err := client.HierarchySettingsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HierarchySettingsOperationGroupClient.HierarchySettingsUpdate`

```go
ctx := context.TODO()
id := commonids.NewManagementGroupID("groupId")

payload := hierarchysettingsoperationgroup.CreateOrUpdateSettingsRequest{
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
