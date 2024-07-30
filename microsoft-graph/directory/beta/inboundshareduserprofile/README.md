
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/inboundshareduserprofile` Documentation

The `inboundshareduserprofile` SDK allows for interaction with the Azure Resource Manager Service `directory` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/directory/beta/inboundshareduserprofile"
```


### Client Initialization

```go
client := inboundshareduserprofile.NewInboundSharedUserProfileClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InboundSharedUserProfileClient.CreateInboundSharedUserProfile`

```go
ctx := context.TODO()

payload := inboundshareduserprofile.InboundSharedUserProfile{
	// ...
}


read, err := client.CreateInboundSharedUserProfile(ctx, payload)
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
id := inboundshareduserprofile.NewDirectoryInboundSharedUserProfileID("inboundSharedUserProfileUserIdValue")

payload := inboundshareduserprofile.CreateInboundSharedUserProfileExportPersonalDataRequest{
	// ...
}


read, err := client.CreateInboundSharedUserProfileExportPersonalData(ctx, id, payload)
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
id := inboundshareduserprofile.NewDirectoryInboundSharedUserProfileID("inboundSharedUserProfileUserIdValue")

read, err := client.DeleteInboundSharedUserProfile(ctx, id)
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
id := inboundshareduserprofile.NewDirectoryInboundSharedUserProfileID("inboundSharedUserProfileUserIdValue")

read, err := client.GetInboundSharedUserProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InboundSharedUserProfileClient.GetInboundSharedUserProfileCount`

```go
ctx := context.TODO()


read, err := client.GetInboundSharedUserProfileCount(ctx)
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


// alternatively `client.ListInboundSharedUserProfiles(ctx)` can be used to do batched pagination
items, err := client.ListInboundSharedUserProfilesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `InboundSharedUserProfileClient.RemoveDirectoryInboundSharedUserProfilePersonalData`

```go
ctx := context.TODO()
id := inboundshareduserprofile.NewDirectoryInboundSharedUserProfileID("inboundSharedUserProfileUserIdValue")

read, err := client.RemoveDirectoryInboundSharedUserProfilePersonalData(ctx, id)
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
id := inboundshareduserprofile.NewDirectoryInboundSharedUserProfileID("inboundSharedUserProfileUserIdValue")

payload := inboundshareduserprofile.InboundSharedUserProfile{
	// ...
}


read, err := client.UpdateInboundSharedUserProfile(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
