
## `github.com/hashicorp/go-azure-sdk/resource-manager/directoryobjects/v1.0/directoryobject` Documentation

The `directoryobject` SDK allows for interaction with the Azure Resource Manager Service `directoryobjects` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/directoryobjects/v1.0/directoryobject"
```


### Client Initialization

```go
client := directoryobject.NewDirectoryObjectClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DirectoryObjectClient.CheckDirectoryObjectByIdMemberGroup`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectIdValue")

payload := directoryobject.CheckDirectoryObjectByIdMemberGroupRequest{
	// ...
}


read, err := client.CheckDirectoryObjectByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryObjectClient.CheckDirectoryObjectByIdMemberObject`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectIdValue")

payload := directoryobject.CheckDirectoryObjectByIdMemberObjectRequest{
	// ...
}


read, err := client.CheckDirectoryObjectByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
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


### Example Usage: `DirectoryObjectClient.DeleteDirectoryObjectById`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectIdValue")

read, err := client.DeleteDirectoryObjectById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryObjectClient.GetDirectoryObjectById`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectIdValue")

read, err := client.GetDirectoryObjectById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryObjectClient.GetDirectoryObjectByIdMemberGroup`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectIdValue")

payload := directoryobject.GetDirectoryObjectByIdMemberGroupRequest{
	// ...
}


read, err := client.GetDirectoryObjectByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryObjectClient.GetDirectoryObjectByIdMemberObject`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectIdValue")

payload := directoryobject.GetDirectoryObjectByIdMemberObjectRequest{
	// ...
}


read, err := client.GetDirectoryObjectByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryObjectClient.GetDirectoryObjectCount`

```go
ctx := context.TODO()


read, err := client.GetDirectoryObjectCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryObjectClient.GetDirectoryObjectsAvailableExtensionProperties`

```go
ctx := context.TODO()


// alternatively `client.GetDirectoryObjectsAvailableExtensionProperties(ctx)` can be used to do batched pagination
items, err := client.GetDirectoryObjectsAvailableExtensionPropertiesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryObjectClient.GetDirectoryObjectsByIds`

```go
ctx := context.TODO()


// alternatively `client.GetDirectoryObjectsByIds(ctx)` can be used to do batched pagination
items, err := client.GetDirectoryObjectsByIdsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryObjectClient.ListDirectoryObjects`

```go
ctx := context.TODO()


// alternatively `client.ListDirectoryObjects(ctx)` can be used to do batched pagination
items, err := client.ListDirectoryObjectsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryObjectClient.RestoreDirectoryObjectById`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectIdValue")

read, err := client.RestoreDirectoryObjectById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryObjectClient.UpdateDirectoryObjectById`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectIdValue")

payload := directoryobject.DirectoryObject{
	// ...
}


read, err := client.UpdateDirectoryObjectById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryObjectClient.ValidateDirectoryObjectsProperty`

```go
ctx := context.TODO()

payload := directoryobject.ValidateDirectoryObjectsPropertyRequest{
	// ...
}


read, err := client.ValidateDirectoryObjectsProperty(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
