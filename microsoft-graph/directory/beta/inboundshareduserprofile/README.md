
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/inboundshareduserprofile` Documentation

The `inboundshareduserprofile` SDK allows for interaction with Microsoft Graph `directory` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/inboundshareduserprofile"
```


### Client Initialization

```go
client := inboundshareduserprofile.NewInboundSharedUserProfileClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InboundSharedUserProfileClient.CreateInboundSharedUserProfile`

```go
ctx := context.TODO()

payload := inboundshareduserprofile.InboundSharedUserProfile{
	// ...
}


read, err := client.CreateInboundSharedUserProfile(ctx, payload, inboundshareduserprofile.DefaultCreateInboundSharedUserProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InboundSharedUserProfileClient.CreateInboundSharedUserProfileExportPersonalData`

```go
ctx := context.TODO()
id := inboundshareduserprofile.NewDirectoryInboundSharedUserProfileID("inboundSharedUserProfileUserId")

payload := inboundshareduserprofile.CreateInboundSharedUserProfileExportPersonalDataRequest{
	// ...
}


read, err := client.CreateInboundSharedUserProfileExportPersonalData(ctx, id, payload, inboundshareduserprofile.DefaultCreateInboundSharedUserProfileExportPersonalDataOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InboundSharedUserProfileClient.DeleteInboundSharedUserProfile`

```go
ctx := context.TODO()
id := inboundshareduserprofile.NewDirectoryInboundSharedUserProfileID("inboundSharedUserProfileUserId")

read, err := client.DeleteInboundSharedUserProfile(ctx, id, inboundshareduserprofile.DefaultDeleteInboundSharedUserProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InboundSharedUserProfileClient.GetInboundSharedUserProfile`

```go
ctx := context.TODO()
id := inboundshareduserprofile.NewDirectoryInboundSharedUserProfileID("inboundSharedUserProfileUserId")

read, err := client.GetInboundSharedUserProfile(ctx, id, inboundshareduserprofile.DefaultGetInboundSharedUserProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InboundSharedUserProfileClient.GetInboundSharedUserProfilesCount`

```go
ctx := context.TODO()


read, err := client.GetInboundSharedUserProfilesCount(ctx, inboundshareduserprofile.DefaultGetInboundSharedUserProfilesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InboundSharedUserProfileClient.ListInboundSharedUserProfiles`

```go
ctx := context.TODO()


// alternatively `client.ListInboundSharedUserProfiles(ctx, inboundshareduserprofile.DefaultListInboundSharedUserProfilesOperationOptions())` can be used to do batched pagination
items, err := client.ListInboundSharedUserProfilesComplete(ctx, inboundshareduserprofile.DefaultListInboundSharedUserProfilesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `InboundSharedUserProfileClient.RemoveInboundSharedUserProfilePersonalData`

```go
ctx := context.TODO()
id := inboundshareduserprofile.NewDirectoryInboundSharedUserProfileID("inboundSharedUserProfileUserId")

read, err := client.RemoveInboundSharedUserProfilePersonalData(ctx, id, inboundshareduserprofile.DefaultRemoveInboundSharedUserProfilePersonalDataOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InboundSharedUserProfileClient.UpdateInboundSharedUserProfile`

```go
ctx := context.TODO()
id := inboundshareduserprofile.NewDirectoryInboundSharedUserProfileID("inboundSharedUserProfileUserId")

payload := inboundshareduserprofile.InboundSharedUserProfile{
	// ...
}


read, err := client.UpdateInboundSharedUserProfile(ctx, id, payload, inboundshareduserprofile.DefaultUpdateInboundSharedUserProfileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
