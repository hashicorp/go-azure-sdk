
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/deleteditem` Documentation

The `deleteditem` SDK allows for interaction with the Azure Resource Manager Service `directory` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/stable/deleteditem"
```


### Client Initialization

```go
client := deleteditem.NewDeletedItemClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeletedItemClient.CheckDirectoryDeletedItemMemberGroup`

```go
ctx := context.TODO()
id := deleteditem.NewDirectoryDeletedItemID("directoryObjectIdValue")

payload := deleteditem.CheckDirectoryDeletedItemMemberGroupRequest{
	// ...
}


read, err := client.CheckDirectoryDeletedItemMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedItemClient.CheckDirectoryDeletedItemMemberObject`

```go
ctx := context.TODO()
id := deleteditem.NewDirectoryDeletedItemID("directoryObjectIdValue")

payload := deleteditem.CheckDirectoryDeletedItemMemberObjectRequest{
	// ...
}


read, err := client.CheckDirectoryDeletedItemMemberObject(ctx, id, payload)
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

read, err := client.GetDeletedItem(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedItemClient.GetDeletedItemCount`

```go
ctx := context.TODO()


read, err := client.GetDeletedItemCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedItemClient.GetDirectoryDeletedItemMemberGroup`

```go
ctx := context.TODO()
id := deleteditem.NewDirectoryDeletedItemID("directoryObjectIdValue")

payload := deleteditem.GetDirectoryDeletedItemMemberGroupRequest{
	// ...
}


read, err := client.GetDirectoryDeletedItemMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedItemClient.GetDirectoryDeletedItemMemberObject`

```go
ctx := context.TODO()
id := deleteditem.NewDirectoryDeletedItemID("directoryObjectIdValue")

payload := deleteditem.GetDirectoryDeletedItemMemberObjectRequest{
	// ...
}


read, err := client.GetDirectoryDeletedItemMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedItemClient.GetDirectoryDeletedItemsAvailableExtensionProperties`

```go
ctx := context.TODO()

payload := deleteditem.GetDirectoryDeletedItemsAvailableExtensionPropertiesRequest{
	// ...
}


// alternatively `client.GetDirectoryDeletedItemsAvailableExtensionProperties(ctx, payload)` can be used to do batched pagination
items, err := client.GetDirectoryDeletedItemsAvailableExtensionPropertiesComplete(ctx, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeletedItemClient.GetDirectoryDeletedItemsByIds`

```go
ctx := context.TODO()

payload := deleteditem.GetDirectoryDeletedItemsByIdsRequest{
	// ...
}


// alternatively `client.GetDirectoryDeletedItemsByIds(ctx, payload)` can be used to do batched pagination
items, err := client.GetDirectoryDeletedItemsByIdsComplete(ctx, payload)
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


// alternatively `client.ListDeletedItems(ctx)` can be used to do batched pagination
items, err := client.ListDeletedItemsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DeletedItemClient.RestoreDirectoryDeletedItem`

```go
ctx := context.TODO()
id := deleteditem.NewDirectoryDeletedItemID("directoryObjectIdValue")

read, err := client.RestoreDirectoryDeletedItem(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeletedItemClient.ValidateDirectoryDeletedItemsProperty`

```go
ctx := context.TODO()

payload := deleteditem.ValidateDirectoryDeletedItemsPropertyRequest{
	// ...
}


read, err := client.ValidateDirectoryDeletedItemsProperty(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
