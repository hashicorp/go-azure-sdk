
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupeventattachment` Documentation

The `groupeventattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupeventattachment"
```


### Client Initialization

```go
client := groupeventattachment.NewGroupEventAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupEventAttachmentClient.CreateGroupByIdEventByIdAttachment`

```go
ctx := context.TODO()
id := groupeventattachment.NewGroupEventID("groupIdValue", "eventIdValue")

payload := groupeventattachment.Attachment{
	// ...
}


read, err := client.CreateGroupByIdEventByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventAttachmentClient.CreateGroupByIdEventByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := groupeventattachment.NewGroupEventID("groupIdValue", "eventIdValue")

payload := groupeventattachment.CreateGroupByIdEventByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventAttachmentClient.DeleteGroupByIdEventByIdAttachmentById`

```go
ctx := context.TODO()
id := groupeventattachment.NewGroupEventAttachmentID("groupIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.DeleteGroupByIdEventByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventAttachmentClient.GetGroupByIdEventByIdAttachmentById`

```go
ctx := context.TODO()
id := groupeventattachment.NewGroupEventAttachmentID("groupIdValue", "eventIdValue", "attachmentIdValue")

read, err := client.GetGroupByIdEventByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventAttachmentClient.GetGroupByIdEventByIdAttachmentCount`

```go
ctx := context.TODO()
id := groupeventattachment.NewGroupEventID("groupIdValue", "eventIdValue")

read, err := client.GetGroupByIdEventByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventAttachmentClient.ListGroupByIdEventByIdAttachments`

```go
ctx := context.TODO()
id := groupeventattachment.NewGroupEventID("groupIdValue", "eventIdValue")

// alternatively `client.ListGroupByIdEventByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdEventByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
