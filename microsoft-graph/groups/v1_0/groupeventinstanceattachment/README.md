
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupeventinstanceattachment` Documentation

The `groupeventinstanceattachment` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/groupeventinstanceattachment"
```


### Client Initialization

```go
client := groupeventinstanceattachment.NewGroupEventInstanceAttachmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupEventInstanceAttachmentClient.CreateGroupByIdEventByIdInstanceByIdAttachment`

```go
ctx := context.TODO()
id := groupeventinstanceattachment.NewGroupEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupeventinstanceattachment.Attachment{
	// ...
}


read, err := client.CreateGroupByIdEventByIdInstanceByIdAttachment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceAttachmentClient.CreateGroupByIdEventByIdInstanceByIdAttachmentCreateUploadSession`

```go
ctx := context.TODO()
id := groupeventinstanceattachment.NewGroupEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

payload := groupeventinstanceattachment.CreateGroupByIdEventByIdInstanceByIdAttachmentCreateUploadSessionRequest{
	// ...
}


read, err := client.CreateGroupByIdEventByIdInstanceByIdAttachmentCreateUploadSession(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceAttachmentClient.DeleteGroupByIdEventByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupeventinstanceattachment.NewGroupEventInstanceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.DeleteGroupByIdEventByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceAttachmentClient.GetGroupByIdEventByIdInstanceByIdAttachmentById`

```go
ctx := context.TODO()
id := groupeventinstanceattachment.NewGroupEventInstanceAttachmentID("groupIdValue", "eventIdValue", "eventId1Value", "attachmentIdValue")

read, err := client.GetGroupByIdEventByIdInstanceByIdAttachmentById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceAttachmentClient.GetGroupByIdEventByIdInstanceByIdAttachmentCount`

```go
ctx := context.TODO()
id := groupeventinstanceattachment.NewGroupEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

read, err := client.GetGroupByIdEventByIdInstanceByIdAttachmentCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupEventInstanceAttachmentClient.ListGroupByIdEventByIdInstanceByIdAttachments`

```go
ctx := context.TODO()
id := groupeventinstanceattachment.NewGroupEventInstanceID("groupIdValue", "eventIdValue", "eventId1Value")

// alternatively `client.ListGroupByIdEventByIdInstanceByIdAttachments(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdEventByIdInstanceByIdAttachmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
