
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicyuploadeddefinitionfile` Documentation

The `grouppolicyuploadeddefinitionfile` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicyuploadeddefinitionfile"
```


### Client Initialization

```go
client := grouppolicyuploadeddefinitionfile.NewGroupPolicyUploadedDefinitionFileClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupPolicyUploadedDefinitionFileClient.AddDeviceManagementGroupPolicyUploadedDefinitionFileLanguageFile`

```go
ctx := context.TODO()
id := grouppolicyuploadeddefinitionfile.NewDeviceManagementGroupPolicyUploadedDefinitionFileID("groupPolicyUploadedDefinitionFileIdValue")

payload := grouppolicyuploadeddefinitionfile.AddDeviceManagementGroupPolicyUploadedDefinitionFileLanguageFileRequest{
	// ...
}


read, err := client.AddDeviceManagementGroupPolicyUploadedDefinitionFileLanguageFile(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyUploadedDefinitionFileClient.CreateGroupPolicyUploadedDefinitionFile`

```go
ctx := context.TODO()

payload := grouppolicyuploadeddefinitionfile.GroupPolicyUploadedDefinitionFile{
	// ...
}


read, err := client.CreateGroupPolicyUploadedDefinitionFile(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyUploadedDefinitionFileClient.CreateGroupPolicyUploadedDefinitionFileUploadNewVersion`

```go
ctx := context.TODO()
id := grouppolicyuploadeddefinitionfile.NewDeviceManagementGroupPolicyUploadedDefinitionFileID("groupPolicyUploadedDefinitionFileIdValue")

payload := grouppolicyuploadeddefinitionfile.CreateGroupPolicyUploadedDefinitionFileUploadNewVersionRequest{
	// ...
}


read, err := client.CreateGroupPolicyUploadedDefinitionFileUploadNewVersion(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyUploadedDefinitionFileClient.DeleteGroupPolicyUploadedDefinitionFile`

```go
ctx := context.TODO()
id := grouppolicyuploadeddefinitionfile.NewDeviceManagementGroupPolicyUploadedDefinitionFileID("groupPolicyUploadedDefinitionFileIdValue")

read, err := client.DeleteGroupPolicyUploadedDefinitionFile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyUploadedDefinitionFileClient.GetGroupPolicyUploadedDefinitionFile`

```go
ctx := context.TODO()
id := grouppolicyuploadeddefinitionfile.NewDeviceManagementGroupPolicyUploadedDefinitionFileID("groupPolicyUploadedDefinitionFileIdValue")

read, err := client.GetGroupPolicyUploadedDefinitionFile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyUploadedDefinitionFileClient.GetGroupPolicyUploadedDefinitionFileCount`

```go
ctx := context.TODO()


read, err := client.GetGroupPolicyUploadedDefinitionFileCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyUploadedDefinitionFileClient.ListGroupPolicyUploadedDefinitionFiles`

```go
ctx := context.TODO()


// alternatively `client.ListGroupPolicyUploadedDefinitionFiles(ctx)` can be used to do batched pagination
items, err := client.ListGroupPolicyUploadedDefinitionFilesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupPolicyUploadedDefinitionFileClient.RemoveDeviceManagementGroupPolicyUploadedDefinitionFile`

```go
ctx := context.TODO()
id := grouppolicyuploadeddefinitionfile.NewDeviceManagementGroupPolicyUploadedDefinitionFileID("groupPolicyUploadedDefinitionFileIdValue")

read, err := client.RemoveDeviceManagementGroupPolicyUploadedDefinitionFile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyUploadedDefinitionFileClient.RemoveDeviceManagementGroupPolicyUploadedDefinitionFileLanguageFile`

```go
ctx := context.TODO()
id := grouppolicyuploadeddefinitionfile.NewDeviceManagementGroupPolicyUploadedDefinitionFileID("groupPolicyUploadedDefinitionFileIdValue")

payload := grouppolicyuploadeddefinitionfile.RemoveDeviceManagementGroupPolicyUploadedDefinitionFileLanguageFileRequest{
	// ...
}


read, err := client.RemoveDeviceManagementGroupPolicyUploadedDefinitionFileLanguageFile(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyUploadedDefinitionFileClient.UpdateDeviceManagementGroupPolicyUploadedDefinitionFileLanguageFile`

```go
ctx := context.TODO()
id := grouppolicyuploadeddefinitionfile.NewDeviceManagementGroupPolicyUploadedDefinitionFileID("groupPolicyUploadedDefinitionFileIdValue")

payload := grouppolicyuploadeddefinitionfile.UpdateDeviceManagementGroupPolicyUploadedDefinitionFileLanguageFileRequest{
	// ...
}


read, err := client.UpdateDeviceManagementGroupPolicyUploadedDefinitionFileLanguageFile(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyUploadedDefinitionFileClient.UpdateGroupPolicyUploadedDefinitionFile`

```go
ctx := context.TODO()
id := grouppolicyuploadeddefinitionfile.NewDeviceManagementGroupPolicyUploadedDefinitionFileID("groupPolicyUploadedDefinitionFileIdValue")

payload := grouppolicyuploadeddefinitionfile.GroupPolicyUploadedDefinitionFile{
	// ...
}


read, err := client.UpdateGroupPolicyUploadedDefinitionFile(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
