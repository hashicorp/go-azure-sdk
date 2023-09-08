
## `github.com/hashicorp/go-azure-sdk/resource-manager/directory/beta/directorydeleteditem` Documentation

The `directorydeleteditem` SDK allows for interaction with the Azure Resource Manager Service `directory` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/directory/beta/directorydeleteditem"
```


### Client Initialization

```go
client := directorydeleteditem.NewDirectoryDeletedItemClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DirectoryDeletedItemClient.CheckDirectoryDeletedItemByIdMemberGroup`

```go
ctx := context.TODO()
id := directorydeleteditem.NewDirectoryDeletedItemID("directoryObjectIdValue")

payload := directorydeleteditem.CheckDirectoryDeletedItemByIdMemberGroupRequest{
	// ...
}


read, err := client.CheckDirectoryDeletedItemByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryDeletedItemClient.CheckDirectoryDeletedItemByIdMemberObject`

```go
ctx := context.TODO()
id := directorydeleteditem.NewDirectoryDeletedItemID("directoryObjectIdValue")

payload := directorydeleteditem.CheckDirectoryDeletedItemByIdMemberObjectRequest{
	// ...
}


read, err := client.CheckDirectoryDeletedItemByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryDeletedItemClient.CreateDirectoryDeletedItem`

```go
ctx := context.TODO()

payload := directorydeleteditem.DirectoryObject{
	// ...
}


read, err := client.CreateDirectoryDeletedItem(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryDeletedItemClient.DeleteDirectoryDeletedItemById`

```go
ctx := context.TODO()
id := directorydeleteditem.NewDirectoryDeletedItemID("directoryObjectIdValue")

read, err := client.DeleteDirectoryDeletedItemById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryDeletedItemClient.GetDirectoryDeletedItemById`

```go
ctx := context.TODO()
id := directorydeleteditem.NewDirectoryDeletedItemID("directoryObjectIdValue")

read, err := client.GetDirectoryDeletedItemById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryDeletedItemClient.GetDirectoryDeletedItemByIdMemberGroup`

```go
ctx := context.TODO()
id := directorydeleteditem.NewDirectoryDeletedItemID("directoryObjectIdValue")

payload := directorydeleteditem.GetDirectoryDeletedItemByIdMemberGroupRequest{
	// ...
}


read, err := client.GetDirectoryDeletedItemByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryDeletedItemClient.GetDirectoryDeletedItemByIdMemberObject`

```go
ctx := context.TODO()
id := directorydeleteditem.NewDirectoryDeletedItemID("directoryObjectIdValue")

payload := directorydeleteditem.GetDirectoryDeletedItemByIdMemberObjectRequest{
	// ...
}


read, err := client.GetDirectoryDeletedItemByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryDeletedItemClient.GetDirectoryDeletedItemCount`

```go
ctx := context.TODO()


read, err := client.GetDirectoryDeletedItemCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryDeletedItemClient.GetDirectoryDeletedItemsByIds`

```go
ctx := context.TODO()


// alternatively `client.GetDirectoryDeletedItemsByIds(ctx)` can be used to do batched pagination
items, err := client.GetDirectoryDeletedItemsByIdsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryDeletedItemClient.GetDirectoryDeletedItemsUserOwnedObject`

```go
ctx := context.TODO()

payload := directorydeleteditem.GetDirectoryDeletedItemsUserOwnedObjectRequest{
	// ...
}


read, err := client.GetDirectoryDeletedItemsUserOwnedObject(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryDeletedItemClient.ListDirectoryDeletedItems`

```go
ctx := context.TODO()


// alternatively `client.ListDirectoryDeletedItems(ctx)` can be used to do batched pagination
items, err := client.ListDirectoryDeletedItemsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryDeletedItemClient.RestoreDirectoryDeletedItemById`

```go
ctx := context.TODO()
id := directorydeleteditem.NewDirectoryDeletedItemID("directoryObjectIdValue")

read, err := client.RestoreDirectoryDeletedItemById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryDeletedItemClient.UpdateDirectoryDeletedItemById`

```go
ctx := context.TODO()
id := directorydeleteditem.NewDirectoryDeletedItemID("directoryObjectIdValue")

payload := directorydeleteditem.DirectoryObject{
	// ...
}


read, err := client.UpdateDirectoryDeletedItemById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryDeletedItemClient.ValidateDirectoryDeletedItemsProperty`

```go
ctx := context.TODO()

payload := directorydeleteditem.ValidateDirectoryDeletedItemsPropertyRequest{
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
