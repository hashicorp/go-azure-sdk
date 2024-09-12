
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryobjects/beta/directoryobject` Documentation

The `directoryobject` SDK allows for interaction with the Azure Resource Manager Service `directoryobjects` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryobjects/beta/directoryobject"
```


### Client Initialization

```go
client := directoryobject.NewDirectoryObjectClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DirectoryObjectClient.CheckMemberGroups`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectIdValue")

payload := directoryobject.CheckMemberGroupsRequest{
	// ...
}


// alternatively `client.CheckMemberGroups(ctx, id, payload, directoryobject.DefaultCheckMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberGroupsComplete(ctx, id, payload, directoryobject.DefaultCheckMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryObjectClient.CheckMemberObjects`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectIdValue")

payload := directoryobject.CheckMemberObjectsRequest{
	// ...
}


// alternatively `client.CheckMemberObjects(ctx, id, payload, directoryobject.DefaultCheckMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberObjectsComplete(ctx, id, payload, directoryobject.DefaultCheckMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryObjectClient.CreateDirectoryObject`

```go
ctx := context.TODO()

payload := directoryobject.DirectoryObject{
	// ...
}


read, err := client.CreateDirectoryObject(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryObjectClient.DeleteDirectoryObject`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectIdValue")

read, err := client.DeleteDirectoryObject(ctx, id, directoryobject.DefaultDeleteDirectoryObjectOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryObjectClient.GetByIds`

```go
ctx := context.TODO()

payload := directoryobject.GetByIdsRequest{
	// ...
}


// alternatively `client.GetByIds(ctx, payload, directoryobject.DefaultGetByIdsOperationOptions())` can be used to do batched pagination
items, err := client.GetByIdsComplete(ctx, payload, directoryobject.DefaultGetByIdsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryObjectClient.GetCount`

```go
ctx := context.TODO()


read, err := client.GetCount(ctx, directoryobject.DefaultGetCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryObjectClient.GetDirectoryObject`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectIdValue")

read, err := client.GetDirectoryObject(ctx, id, directoryobject.DefaultGetDirectoryObjectOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryObjectClient.GetMemberGroups`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectIdValue")

payload := directoryobject.GetMemberGroupsRequest{
	// ...
}


// alternatively `client.GetMemberGroups(ctx, id, payload, directoryobject.DefaultGetMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberGroupsComplete(ctx, id, payload, directoryobject.DefaultGetMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryObjectClient.GetMemberObjects`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectIdValue")

payload := directoryobject.GetMemberObjectsRequest{
	// ...
}


// alternatively `client.GetMemberObjects(ctx, id, payload, directoryobject.DefaultGetMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberObjectsComplete(ctx, id, payload, directoryobject.DefaultGetMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryObjectClient.GetUserOwnedObject`

```go
ctx := context.TODO()

payload := directoryobject.GetUserOwnedObjectRequest{
	// ...
}


read, err := client.GetUserOwnedObject(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryObjectClient.ListDirectoryObjects`

```go
ctx := context.TODO()


// alternatively `client.ListDirectoryObjects(ctx, directoryobject.DefaultListDirectoryObjectsOperationOptions())` can be used to do batched pagination
items, err := client.ListDirectoryObjectsComplete(ctx, directoryobject.DefaultListDirectoryObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryObjectClient.Restore`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectIdValue")

payload := directoryobject.RestoreRequest{
	// ...
}


read, err := client.Restore(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryObjectClient.UpdateDirectoryObject`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectIdValue")

payload := directoryobject.DirectoryObject{
	// ...
}


read, err := client.UpdateDirectoryObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryObjectClient.ValidateProperty`

```go
ctx := context.TODO()

payload := directoryobject.ValidatePropertyRequest{
	// ...
}


read, err := client.ValidateProperty(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
