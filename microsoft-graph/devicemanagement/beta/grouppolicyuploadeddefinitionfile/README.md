
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


### Example Usage: `GroupPolicyUploadedDefinitionFileClient.AddGroupPolicyUploadedDefinitionFileLanguageFile`

```go
ctx := context.TODO()
id := grouppolicyuploadeddefinitionfile.NewDeviceManagementGroupPolicyUploadedDefinitionFileID("groupPolicyUploadedDefinitionFileIdValue")

payload := grouppolicyuploadeddefinitionfile.AddGroupPolicyUploadedDefinitionFileLanguageFileRequest{
	// ...
}


read, err := client.AddGroupPolicyUploadedDefinitionFileLanguageFile(ctx, id, payload)
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
id := grouppolicyuploadeddefinitionfile.NewDeviceManagementGroupPolicyUploadedDefinitionFileID("groupPolicyUploadedDefinitionFileIdValue")

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
id := grouppolicyuploadeddefinitionfile.NewDeviceManagementGroupPolicyUploadedDefinitionFileID("groupPolicyUploadedDefinitionFileIdValue")

read, err := client.RemoveGroupPolicyUploadedDefinitionFile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyUploadedDefinitionFileClient.RemoveGroupPolicyUploadedDefinitionFileLanguageFile`

```go
ctx := context.TODO()
id := grouppolicyuploadeddefinitionfile.NewDeviceManagementGroupPolicyUploadedDefinitionFileID("groupPolicyUploadedDefinitionFileIdValue")

payload := grouppolicyuploadeddefinitionfile.RemoveGroupPolicyUploadedDefinitionFileLanguageFileRequest{
	// ...
}


read, err := client.RemoveGroupPolicyUploadedDefinitionFileLanguageFile(ctx, id, payload)
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


### Example Usage: `GroupPolicyUploadedDefinitionFileClient.UpdateGroupPolicyUploadedDefinitionFileLanguageFile`

```go
ctx := context.TODO()
id := grouppolicyuploadeddefinitionfile.NewDeviceManagementGroupPolicyUploadedDefinitionFileID("groupPolicyUploadedDefinitionFileIdValue")

payload := grouppolicyuploadeddefinitionfile.UpdateGroupPolicyUploadedDefinitionFileLanguageFileRequest{
	// ...
}


read, err := client.UpdateGroupPolicyUploadedDefinitionFileLanguageFile(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
