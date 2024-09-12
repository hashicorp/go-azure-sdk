
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/deleteditem` Documentation

The `deleteditem` SDK allows for interaction with the Azure Resource Manager Service `directory` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/deleteditem"
```


### Client Initialization

```go
client := deleteditem.NewDeletedItemClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeletedItemClient.CheckDeletedItemMemberGroups`

```go
ctx := context.TODO()
id := deleteditem.NewDirectoryDeletedItemID("directoryObjectIdValue")

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
id := deleteditem.NewDirectoryDeletedItemID("directoryObjectIdValue")

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
id := deleteditem.NewDirectoryDeletedItemID("directoryObjectIdValue")

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
id := deleteditem.NewDirectoryDeletedItemID("directoryObjectIdValue")

read, err := client.GetDeletedItem(ctx, id, deleteditem.DefaultGetDeletedItemOperationOptions())
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
id := deleteditem.NewDirectoryDeletedItemID("directoryObjectIdValue")

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
id := deleteditem.NewDirectoryDeletedItemID("directoryObjectIdValue")

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
id := deleteditem.NewDirectoryDeletedItemID("directoryObjectIdValue")

payload := deleteditem.RestoreDeletedItemRequest{
	// ...
}


read, err := client.RestoreDeletedItem(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedItemClient.ValidateDeletedItemsProperty`

```go
ctx := context.TODO()

payload := deleteditem.ValidateDeletedItemsPropertyRequest{
	// ...
}


read, err := client.ValidateDeletedItemsProperty(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
