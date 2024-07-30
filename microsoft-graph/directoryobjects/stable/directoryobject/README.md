
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryobjects/stable/directoryobject` Documentation

The `directoryobject` SDK allows for interaction with the Azure Resource Manager Service `directoryobjects` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryobjects/stable/directoryobject"
```


### Client Initialization

```go
client := directoryobject.NewDirectoryObjectClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DirectoryObjectClient.CheckDirectoryObjectMemberGroup`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectIdValue")

payload := directoryobject.CheckDirectoryObjectMemberGroupRequest{
	// ...
}


read, err := client.CheckDirectoryObjectMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryObjectClient.CheckDirectoryObjectMemberObject`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectIdValue")

payload := directoryobject.CheckDirectoryObjectMemberObjectRequest{
	// ...
}


read, err := client.CheckDirectoryObjectMemberObject(ctx, id, payload)
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


### Example Usage: `DirectoryObjectClient.DeleteDirectoryObject`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectIdValue")

read, err := client.DeleteDirectoryObject(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryObjectClient.GetCount`

```go
ctx := context.TODO()


read, err := client.GetCount(ctx)
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

read, err := client.GetDirectoryObject(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryObjectClient.GetDirectoryObjectMemberGroup`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectIdValue")

payload := directoryobject.GetDirectoryObjectMemberGroupRequest{
	// ...
}


read, err := client.GetDirectoryObjectMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryObjectClient.GetDirectoryObjectMemberObject`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectIdValue")

payload := directoryobject.GetDirectoryObjectMemberObjectRequest{
	// ...
}


read, err := client.GetDirectoryObjectMemberObject(ctx, id, payload)
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

payload := directoryobject.GetDirectoryObjectsAvailableExtensionPropertiesRequest{
	// ...
}


// alternatively `client.GetDirectoryObjectsAvailableExtensionProperties(ctx, payload)` can be used to do batched pagination
items, err := client.GetDirectoryObjectsAvailableExtensionPropertiesComplete(ctx, payload)
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

payload := directoryobject.GetDirectoryObjectsByIdsRequest{
	// ...
}


// alternatively `client.GetDirectoryObjectsByIds(ctx, payload)` can be used to do batched pagination
items, err := client.GetDirectoryObjectsByIdsComplete(ctx, payload)
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


### Example Usage: `DirectoryObjectClient.RestoreDirectoryObject`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectIdValue")

read, err := client.RestoreDirectoryObject(ctx, id)
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
