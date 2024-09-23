
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicyuploadeddefinitionfile` Documentation

The `grouppolicyuploadeddefinitionfile` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicyuploadeddefinitionfile"
```


### Client Initialization

```go
client := grouppolicyuploadeddefinitionfile.NewGroupPolicyUploadedDefinitionFileClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupPolicyUploadedDefinitionFileClient.AddGroupPolicyUploadedDefinitionFileLanguageFiles`

```go
ctx := context.TODO()
id := grouppolicyuploadeddefinitionfile.NewDeviceManagementGroupPolicyUploadedDefinitionFileID("groupPolicyUploadedDefinitionFileId")

payload := grouppolicyuploadeddefinitionfile.AddGroupPolicyUploadedDefinitionFileLanguageFilesRequest{
	// ...
}


read, err := client.AddGroupPolicyUploadedDefinitionFileLanguageFiles(ctx, id, payload, grouppolicyuploadeddefinitionfile.DefaultAddGroupPolicyUploadedDefinitionFileLanguageFilesOperationOptions())
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


read, err := client.CreateGroupPolicyUploadedDefinitionFile(ctx, payload, grouppolicyuploadeddefinitionfile.DefaultCreateGroupPolicyUploadedDefinitionFileOperationOptions())
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
id := grouppolicyuploadeddefinitionfile.NewDeviceManagementGroupPolicyUploadedDefinitionFileID("groupPolicyUploadedDefinitionFileId")

payload := grouppolicyuploadeddefinitionfile.CreateGroupPolicyUploadedDefinitionFileUploadNewVersionRequest{
	// ...
}


read, err := client.CreateGroupPolicyUploadedDefinitionFileUploadNewVersion(ctx, id, payload, grouppolicyuploadeddefinitionfile.DefaultCreateGroupPolicyUploadedDefinitionFileUploadNewVersionOperationOptions())
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
id := grouppolicyuploadeddefinitionfile.NewDeviceManagementGroupPolicyUploadedDefinitionFileID("groupPolicyUploadedDefinitionFileId")

read, err := client.DeleteGroupPolicyUploadedDefinitionFile(ctx, id, grouppolicyuploadeddefinitionfile.DefaultDeleteGroupPolicyUploadedDefinitionFileOperationOptions())
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
id := grouppolicyuploadeddefinitionfile.NewDeviceManagementGroupPolicyUploadedDefinitionFileID("groupPolicyUploadedDefinitionFileId")

read, err := client.GetGroupPolicyUploadedDefinitionFile(ctx, id, grouppolicyuploadeddefinitionfile.DefaultGetGroupPolicyUploadedDefinitionFileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyUploadedDefinitionFileClient.GetGroupPolicyUploadedDefinitionFilesCount`

```go
ctx := context.TODO()


read, err := client.GetGroupPolicyUploadedDefinitionFilesCount(ctx, grouppolicyuploadeddefinitionfile.DefaultGetGroupPolicyUploadedDefinitionFilesCountOperationOptions())
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


// alternatively `client.ListGroupPolicyUploadedDefinitionFiles(ctx, grouppolicyuploadeddefinitionfile.DefaultListGroupPolicyUploadedDefinitionFilesOperationOptions())` can be used to do batched pagination
items, err := client.ListGroupPolicyUploadedDefinitionFilesComplete(ctx, grouppolicyuploadeddefinitionfile.DefaultListGroupPolicyUploadedDefinitionFilesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupPolicyUploadedDefinitionFileClient.RemoveGroupPolicyUploadedDefinitionFile`

```go
ctx := context.TODO()
id := grouppolicyuploadeddefinitionfile.NewDeviceManagementGroupPolicyUploadedDefinitionFileID("groupPolicyUploadedDefinitionFileId")

read, err := client.RemoveGroupPolicyUploadedDefinitionFile(ctx, id, grouppolicyuploadeddefinitionfile.DefaultRemoveGroupPolicyUploadedDefinitionFileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyUploadedDefinitionFileClient.RemoveGroupPolicyUploadedDefinitionFileLanguageFiles`

```go
ctx := context.TODO()
id := grouppolicyuploadeddefinitionfile.NewDeviceManagementGroupPolicyUploadedDefinitionFileID("groupPolicyUploadedDefinitionFileId")

payload := grouppolicyuploadeddefinitionfile.RemoveGroupPolicyUploadedDefinitionFileLanguageFilesRequest{
	// ...
}


read, err := client.RemoveGroupPolicyUploadedDefinitionFileLanguageFiles(ctx, id, payload, grouppolicyuploadeddefinitionfile.DefaultRemoveGroupPolicyUploadedDefinitionFileLanguageFilesOperationOptions())
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
id := grouppolicyuploadeddefinitionfile.NewDeviceManagementGroupPolicyUploadedDefinitionFileID("groupPolicyUploadedDefinitionFileId")

payload := grouppolicyuploadeddefinitionfile.GroupPolicyUploadedDefinitionFile{
	// ...
}


read, err := client.UpdateGroupPolicyUploadedDefinitionFile(ctx, id, payload, grouppolicyuploadeddefinitionfile.DefaultUpdateGroupPolicyUploadedDefinitionFileOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyUploadedDefinitionFileClient.UpdateGroupPolicyUploadedDefinitionFileLanguageFiles`

```go
ctx := context.TODO()
id := grouppolicyuploadeddefinitionfile.NewDeviceManagementGroupPolicyUploadedDefinitionFileID("groupPolicyUploadedDefinitionFileId")

payload := grouppolicyuploadeddefinitionfile.UpdateGroupPolicyUploadedDefinitionFileLanguageFilesRequest{
	// ...
}


read, err := client.UpdateGroupPolicyUploadedDefinitionFileLanguageFiles(ctx, id, payload, grouppolicyuploadeddefinitionfile.DefaultUpdateGroupPolicyUploadedDefinitionFileLanguageFilesOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
