
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meonenotesection` Documentation

The `meonenotesection` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meonenotesection"
```


### Client Initialization

```go
client := meonenotesection.NewMeOnenoteSectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeOnenoteSectionClient.CreateMeOnenoteSection`

```go
ctx := context.TODO()

payload := meonenotesection.OnenoteSection{
	// ...
}


read, err := client.CreateMeOnenoteSection(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteSectionClient.CreateMeOnenoteSectionByIdCopyToNotebook`

```go
ctx := context.TODO()
id := meonenotesection.NewMeOnenoteSectionID("onenoteSectionIdValue")

payload := meonenotesection.CreateMeOnenoteSectionByIdCopyToNotebookRequest{
	// ...
}


read, err := client.CreateMeOnenoteSectionByIdCopyToNotebook(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteSectionClient.CreateMeOnenoteSectionByIdCopyToSectionGroup`

```go
ctx := context.TODO()
id := meonenotesection.NewMeOnenoteSectionID("onenoteSectionIdValue")

payload := meonenotesection.CreateMeOnenoteSectionByIdCopyToSectionGroupRequest{
	// ...
}


read, err := client.CreateMeOnenoteSectionByIdCopyToSectionGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteSectionClient.DeleteMeOnenoteSectionById`

```go
ctx := context.TODO()
id := meonenotesection.NewMeOnenoteSectionID("onenoteSectionIdValue")

read, err := client.DeleteMeOnenoteSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteSectionClient.GetMeOnenoteSectionById`

```go
ctx := context.TODO()
id := meonenotesection.NewMeOnenoteSectionID("onenoteSectionIdValue")

read, err := client.GetMeOnenoteSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteSectionClient.GetMeOnenoteSectionCount`

```go
ctx := context.TODO()


read, err := client.GetMeOnenoteSectionCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenoteSectionClient.ListMeOnenoteSections`

```go
ctx := context.TODO()


// alternatively `client.ListMeOnenoteSections(ctx)` can be used to do batched pagination
items, err := client.ListMeOnenoteSectionsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeOnenoteSectionClient.UpdateMeOnenoteSectionById`

```go
ctx := context.TODO()
id := meonenotesection.NewMeOnenoteSectionID("onenoteSectionIdValue")

payload := meonenotesection.OnenoteSection{
	// ...
}


read, err := client.UpdateMeOnenoteSectionById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
