
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/grouponenotesection` Documentation

The `grouponenotesection` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/grouponenotesection"
```


### Client Initialization

```go
client := grouponenotesection.NewGroupOnenoteSectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupOnenoteSectionClient.CreateGroupByIdOnenoteSection`

```go
ctx := context.TODO()
id := grouponenotesection.NewGroupID("groupIdValue")

payload := grouponenotesection.OnenoteSection{
	// ...
}


read, err := client.CreateGroupByIdOnenoteSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteSectionClient.CreateGroupByIdOnenoteSectionByIdCopyToNotebook`

```go
ctx := context.TODO()
id := grouponenotesection.NewGroupOnenoteSectionID("groupIdValue", "onenoteSectionIdValue")

payload := grouponenotesection.CreateGroupByIdOnenoteSectionByIdCopyToNotebookRequest{
	// ...
}


read, err := client.CreateGroupByIdOnenoteSectionByIdCopyToNotebook(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteSectionClient.CreateGroupByIdOnenoteSectionByIdCopyToSectionGroup`

```go
ctx := context.TODO()
id := grouponenotesection.NewGroupOnenoteSectionID("groupIdValue", "onenoteSectionIdValue")

payload := grouponenotesection.CreateGroupByIdOnenoteSectionByIdCopyToSectionGroupRequest{
	// ...
}


read, err := client.CreateGroupByIdOnenoteSectionByIdCopyToSectionGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteSectionClient.DeleteGroupByIdOnenoteSectionById`

```go
ctx := context.TODO()
id := grouponenotesection.NewGroupOnenoteSectionID("groupIdValue", "onenoteSectionIdValue")

read, err := client.DeleteGroupByIdOnenoteSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteSectionClient.GetGroupByIdOnenoteSectionById`

```go
ctx := context.TODO()
id := grouponenotesection.NewGroupOnenoteSectionID("groupIdValue", "onenoteSectionIdValue")

read, err := client.GetGroupByIdOnenoteSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteSectionClient.GetGroupByIdOnenoteSectionCount`

```go
ctx := context.TODO()
id := grouponenotesection.NewGroupID("groupIdValue")

read, err := client.GetGroupByIdOnenoteSectionCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenoteSectionClient.ListGroupByIdOnenoteSections`

```go
ctx := context.TODO()
id := grouponenotesection.NewGroupID("groupIdValue")

// alternatively `client.ListGroupByIdOnenoteSections(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdOnenoteSectionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupOnenoteSectionClient.UpdateGroupByIdOnenoteSectionById`

```go
ctx := context.TODO()
id := grouponenotesection.NewGroupOnenoteSectionID("groupIdValue", "onenoteSectionIdValue")

payload := grouponenotesection.OnenoteSection{
	// ...
}


read, err := client.UpdateGroupByIdOnenoteSectionById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
