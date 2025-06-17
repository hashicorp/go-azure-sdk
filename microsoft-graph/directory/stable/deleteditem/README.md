
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/deleteditem` Documentation

The `deleteditem` SDK allows for interaction with Microsoft Graph `directory` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/deleteditem"
```


### Client Initialization

```go
client := deleteditem.NewDeletedItemClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeletedItemClient.CheckDeletedItemMemberGroups`

```go
ctx := context.TODO()
id := deleteditem.NewDirectoryDeletedItemID("directoryObjectId")

payload := deleteditem.CheckDeletedItemMemberGroupsRequest{
	// ...
}


// alternatively `client.CheckDeletedItemMemberGroups(ctx, id, payload, deleteditem.DefaultCheckDeletedItemMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.CheckDeletedItemMemberGroupsComplete(ctx, id, payload, deleteditem.DefaultCheckDeletedItemMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeletedItemClient.CheckDeletedItemMemberObjects`

```go
ctx := context.TODO()
id := deleteditem.NewDirectoryDeletedItemID("directoryObjectId")

payload := deleteditem.CheckDeletedItemMemberObjectsRequest{
	// ...
}


// alternatively `client.CheckDeletedItemMemberObjects(ctx, id, payload, deleteditem.DefaultCheckDeletedItemMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.CheckDeletedItemMemberObjectsComplete(ctx, id, payload, deleteditem.DefaultCheckDeletedItemMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeletedItemClient.DeleteDeletedItem`

```go
ctx := context.TODO()
id := deleteditem.NewDirectoryDeletedItemID("directoryObjectId")

read, err := client.DeleteDeletedItem(ctx, id, deleteditem.DefaultDeleteDeletedItemOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedItemClient.GetDeletedItem`

```go
ctx := context.TODO()
id := deleteditem.NewDirectoryDeletedItemID("directoryObjectId")

read, err := client.GetDeletedItem(ctx, id, deleteditem.DefaultGetDeletedItemOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedItemClient.GetDeletedItemAdministrativeUnit`

```go
ctx := context.TODO()
id := deleteditem.NewDirectoryDeletedItemID("directoryObjectId")

read, err := client.GetDeletedItemAdministrativeUnit(ctx, id, deleteditem.DefaultGetDeletedItemAdministrativeUnitOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedItemClient.GetDeletedItemAdministrativeUnitCount`

```go
ctx := context.TODO()


read, err := client.GetDeletedItemAdministrativeUnitCount(ctx, deleteditem.DefaultGetDeletedItemAdministrativeUnitCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedItemClient.GetDeletedItemApplication`

```go
ctx := context.TODO()
id := deleteditem.NewDirectoryDeletedItemID("directoryObjectId")

read, err := client.GetDeletedItemApplication(ctx, id, deleteditem.DefaultGetDeletedItemApplicationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedItemClient.GetDeletedItemApplicationCount`

```go
ctx := context.TODO()


read, err := client.GetDeletedItemApplicationCount(ctx, deleteditem.DefaultGetDeletedItemApplicationCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedItemClient.GetDeletedItemDevice`

```go
ctx := context.TODO()
id := deleteditem.NewDirectoryDeletedItemID("directoryObjectId")

read, err := client.GetDeletedItemDevice(ctx, id, deleteditem.DefaultGetDeletedItemDeviceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedItemClient.GetDeletedItemDeviceCount`

```go
ctx := context.TODO()


read, err := client.GetDeletedItemDeviceCount(ctx, deleteditem.DefaultGetDeletedItemDeviceCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedItemClient.GetDeletedItemGroup`

```go
ctx := context.TODO()
id := deleteditem.NewDirectoryDeletedItemID("directoryObjectId")

read, err := client.GetDeletedItemGroup(ctx, id, deleteditem.DefaultGetDeletedItemGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedItemClient.GetDeletedItemGroupCount`

```go
ctx := context.TODO()


read, err := client.GetDeletedItemGroupCount(ctx, deleteditem.DefaultGetDeletedItemGroupCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedItemClient.GetDeletedItemMemberGroups`

```go
ctx := context.TODO()
id := deleteditem.NewDirectoryDeletedItemID("directoryObjectId")

payload := deleteditem.GetDeletedItemMemberGroupsRequest{
	// ...
}


// alternatively `client.GetDeletedItemMemberGroups(ctx, id, payload, deleteditem.DefaultGetDeletedItemMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.GetDeletedItemMemberGroupsComplete(ctx, id, payload, deleteditem.DefaultGetDeletedItemMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeletedItemClient.GetDeletedItemMemberObjects`

```go
ctx := context.TODO()
id := deleteditem.NewDirectoryDeletedItemID("directoryObjectId")

payload := deleteditem.GetDeletedItemMemberObjectsRequest{
	// ...
}


// alternatively `client.GetDeletedItemMemberObjects(ctx, id, payload, deleteditem.DefaultGetDeletedItemMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.GetDeletedItemMemberObjectsComplete(ctx, id, payload, deleteditem.DefaultGetDeletedItemMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeletedItemClient.GetDeletedItemServicePrincipal`

```go
ctx := context.TODO()
id := deleteditem.NewDirectoryDeletedItemID("directoryObjectId")

read, err := client.GetDeletedItemServicePrincipal(ctx, id, deleteditem.DefaultGetDeletedItemServicePrincipalOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedItemClient.GetDeletedItemServicePrincipalCount`

```go
ctx := context.TODO()


read, err := client.GetDeletedItemServicePrincipalCount(ctx, deleteditem.DefaultGetDeletedItemServicePrincipalCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedItemClient.GetDeletedItemUser`

```go
ctx := context.TODO()
id := deleteditem.NewDirectoryDeletedItemID("directoryObjectId")

read, err := client.GetDeletedItemUser(ctx, id, deleteditem.DefaultGetDeletedItemUserOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedItemClient.GetDeletedItemUserCount`

```go
ctx := context.TODO()


read, err := client.GetDeletedItemUserCount(ctx, deleteditem.DefaultGetDeletedItemUserCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedItemClient.GetDeletedItemsAvailableExtensionProperties`

```go
ctx := context.TODO()

payload := deleteditem.GetDeletedItemsAvailableExtensionPropertiesRequest{
	// ...
}


// alternatively `client.GetDeletedItemsAvailableExtensionProperties(ctx, payload, deleteditem.DefaultGetDeletedItemsAvailableExtensionPropertiesOperationOptions())` can be used to do batched pagination
items, err := client.GetDeletedItemsAvailableExtensionPropertiesComplete(ctx, payload, deleteditem.DefaultGetDeletedItemsAvailableExtensionPropertiesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeletedItemClient.GetDeletedItemsByIds`

```go
ctx := context.TODO()

payload := deleteditem.GetDeletedItemsByIdsRequest{
	// ...
}


// alternatively `client.GetDeletedItemsByIds(ctx, payload, deleteditem.DefaultGetDeletedItemsByIdsOperationOptions())` can be used to do batched pagination
items, err := client.GetDeletedItemsByIdsComplete(ctx, payload, deleteditem.DefaultGetDeletedItemsByIdsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeletedItemClient.GetDeletedItemsCount`

```go
ctx := context.TODO()


read, err := client.GetDeletedItemsCount(ctx, deleteditem.DefaultGetDeletedItemsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedItemClient.ListDeletedItemAdministrativeUnits`

```go
ctx := context.TODO()


// alternatively `client.ListDeletedItemAdministrativeUnits(ctx, deleteditem.DefaultListDeletedItemAdministrativeUnitsOperationOptions())` can be used to do batched pagination
items, err := client.ListDeletedItemAdministrativeUnitsComplete(ctx, deleteditem.DefaultListDeletedItemAdministrativeUnitsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeletedItemClient.ListDeletedItemApplications`

```go
ctx := context.TODO()


// alternatively `client.ListDeletedItemApplications(ctx, deleteditem.DefaultListDeletedItemApplicationsOperationOptions())` can be used to do batched pagination
items, err := client.ListDeletedItemApplicationsComplete(ctx, deleteditem.DefaultListDeletedItemApplicationsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeletedItemClient.ListDeletedItemDevices`

```go
ctx := context.TODO()


// alternatively `client.ListDeletedItemDevices(ctx, deleteditem.DefaultListDeletedItemDevicesOperationOptions())` can be used to do batched pagination
items, err := client.ListDeletedItemDevicesComplete(ctx, deleteditem.DefaultListDeletedItemDevicesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeletedItemClient.ListDeletedItemGroups`

```go
ctx := context.TODO()


// alternatively `client.ListDeletedItemGroups(ctx, deleteditem.DefaultListDeletedItemGroupsOperationOptions())` can be used to do batched pagination
items, err := client.ListDeletedItemGroupsComplete(ctx, deleteditem.DefaultListDeletedItemGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeletedItemClient.ListDeletedItemServicePrincipals`

```go
ctx := context.TODO()


// alternatively `client.ListDeletedItemServicePrincipals(ctx, deleteditem.DefaultListDeletedItemServicePrincipalsOperationOptions())` can be used to do batched pagination
items, err := client.ListDeletedItemServicePrincipalsComplete(ctx, deleteditem.DefaultListDeletedItemServicePrincipalsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeletedItemClient.ListDeletedItemUsers`

```go
ctx := context.TODO()


// alternatively `client.ListDeletedItemUsers(ctx, deleteditem.DefaultListDeletedItemUsersOperationOptions())` can be used to do batched pagination
items, err := client.ListDeletedItemUsersComplete(ctx, deleteditem.DefaultListDeletedItemUsersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeletedItemClient.ListDeletedItems`

```go
ctx := context.TODO()


// alternatively `client.ListDeletedItems(ctx, deleteditem.DefaultListDeletedItemsOperationOptions())` can be used to do batched pagination
items, err := client.ListDeletedItemsComplete(ctx, deleteditem.DefaultListDeletedItemsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeletedItemClient.RestoreDeletedItem`

```go
ctx := context.TODO()
id := deleteditem.NewDirectoryDeletedItemID("directoryObjectId")

read, err := client.RestoreDeletedItem(ctx, id, deleteditem.DefaultRestoreDeletedItemOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedItemClient.ValidateDeletedItemsProperties`

```go
ctx := context.TODO()

payload := deleteditem.ValidateDeletedItemsPropertiesRequest{
	// ...
}


read, err := client.ValidateDeletedItemsProperties(ctx, payload, deleteditem.DefaultValidateDeletedItemsPropertiesOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
