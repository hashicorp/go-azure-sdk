
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/grouponenotepage` Documentation

The `grouponenotepage` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/v1.0/grouponenotepage"
```


### Client Initialization

```go
client := grouponenotepage.NewGroupOnenotePageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupOnenotePageClient.CreateGroupByIdOnenotePage`

```go
ctx := context.TODO()
id := grouponenotepage.NewGroupID("groupIdValue")

payload := grouponenotepage.OnenotePage{
	// ...
}


read, err := client.CreateGroupByIdOnenotePage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenotePageClient.CreateGroupByIdOnenotePageByIdCopyToSection`

```go
ctx := context.TODO()
id := grouponenotepage.NewGroupOnenotePageID("groupIdValue", "onenotePageIdValue")

payload := grouponenotepage.CreateGroupByIdOnenotePageByIdCopyToSectionRequest{
	// ...
}


read, err := client.CreateGroupByIdOnenotePageByIdCopyToSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenotePageClient.CreateGroupByIdOnenotePageByIdOnenotePatchContent`

```go
ctx := context.TODO()
id := grouponenotepage.NewGroupOnenotePageID("groupIdValue", "onenotePageIdValue")

payload := grouponenotepage.CreateGroupByIdOnenotePageByIdOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateGroupByIdOnenotePageByIdOnenotePatchContent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenotePageClient.DeleteGroupByIdOnenotePageById`

```go
ctx := context.TODO()
id := grouponenotepage.NewGroupOnenotePageID("groupIdValue", "onenotePageIdValue")

read, err := client.DeleteGroupByIdOnenotePageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenotePageClient.GetGroupByIdOnenotePageById`

```go
ctx := context.TODO()
id := grouponenotepage.NewGroupOnenotePageID("groupIdValue", "onenotePageIdValue")

read, err := client.GetGroupByIdOnenotePageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenotePageClient.GetGroupByIdOnenotePageCount`

```go
ctx := context.TODO()
id := grouponenotepage.NewGroupID("groupIdValue")

read, err := client.GetGroupByIdOnenotePageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupOnenotePageClient.ListGroupByIdOnenotePages`

```go
ctx := context.TODO()
id := grouponenotepage.NewGroupID("groupIdValue")

// alternatively `client.ListGroupByIdOnenotePages(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdOnenotePagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupOnenotePageClient.UpdateGroupByIdOnenotePageById`

```go
ctx := context.TODO()
id := grouponenotepage.NewGroupOnenotePageID("groupIdValue", "onenotePageIdValue")

payload := grouponenotepage.OnenotePage{
	// ...
}


read, err := client.UpdateGroupByIdOnenotePageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
