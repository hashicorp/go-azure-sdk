
## `github.com/hashicorp/go-azure-sdk/resource-manager/directory/beta/directoryinboundshareduserprofile` Documentation

The `directoryinboundshareduserprofile` SDK allows for interaction with the Azure Resource Manager Service `directory` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/directory/beta/directoryinboundshareduserprofile"
```


### Client Initialization

```go
client := directoryinboundshareduserprofile.NewDirectoryInboundSharedUserProfileClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DirectoryInboundSharedUserProfileClient.CreateDirectoryInboundSharedUserProfile`

```go
ctx := context.TODO()

payload := directoryinboundshareduserprofile.InboundSharedUserProfile{
	// ...
}


read, err := client.CreateDirectoryInboundSharedUserProfile(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryInboundSharedUserProfileClient.CreateDirectoryInboundSharedUserProfileByIdExportPersonalData`

```go
ctx := context.TODO()
id := directoryinboundshareduserprofile.NewDirectoryInboundSharedUserProfileID("inboundSharedUserProfileUserIdValue")

payload := directoryinboundshareduserprofile.CreateDirectoryInboundSharedUserProfileByIdExportPersonalDataRequest{
	// ...
}


read, err := client.CreateDirectoryInboundSharedUserProfileByIdExportPersonalData(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryInboundSharedUserProfileClient.DeleteDirectoryInboundSharedUserProfileById`

```go
ctx := context.TODO()
id := directoryinboundshareduserprofile.NewDirectoryInboundSharedUserProfileID("inboundSharedUserProfileUserIdValue")

read, err := client.DeleteDirectoryInboundSharedUserProfileById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryInboundSharedUserProfileClient.GetDirectoryInboundSharedUserProfileById`

```go
ctx := context.TODO()
id := directoryinboundshareduserprofile.NewDirectoryInboundSharedUserProfileID("inboundSharedUserProfileUserIdValue")

read, err := client.GetDirectoryInboundSharedUserProfileById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryInboundSharedUserProfileClient.GetDirectoryInboundSharedUserProfileCount`

```go
ctx := context.TODO()


read, err := client.GetDirectoryInboundSharedUserProfileCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryInboundSharedUserProfileClient.ListDirectoryInboundSharedUserProfiles`

```go
ctx := context.TODO()


// alternatively `client.ListDirectoryInboundSharedUserProfiles(ctx)` can be used to do batched pagination
items, err := client.ListDirectoryInboundSharedUserProfilesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryInboundSharedUserProfileClient.RemoveDirectoryInboundSharedUserProfileByIdPersonalData`

```go
ctx := context.TODO()
id := directoryinboundshareduserprofile.NewDirectoryInboundSharedUserProfileID("inboundSharedUserProfileUserIdValue")

read, err := client.RemoveDirectoryInboundSharedUserProfileByIdPersonalData(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryInboundSharedUserProfileClient.UpdateDirectoryInboundSharedUserProfileById`

```go
ctx := context.TODO()
id := directoryinboundshareduserprofile.NewDirectoryInboundSharedUserProfileID("inboundSharedUserProfileUserIdValue")

payload := directoryinboundshareduserprofile.InboundSharedUserProfile{
	// ...
}


read, err := client.UpdateDirectoryInboundSharedUserProfileById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
