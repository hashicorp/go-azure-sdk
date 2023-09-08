
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meonenotepage` Documentation

The `meonenotepage` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meonenotepage"
```


### Client Initialization

```go
client := meonenotepage.NewMeOnenotePageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeOnenotePageClient.CreateMeOnenotePage`

```go
ctx := context.TODO()

payload := meonenotepage.OnenotePage{
	// ...
}


read, err := client.CreateMeOnenotePage(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenotePageClient.CreateMeOnenotePageByIdCopyToSection`

```go
ctx := context.TODO()
id := meonenotepage.NewMeOnenotePageID("onenotePageIdValue")

payload := meonenotepage.CreateMeOnenotePageByIdCopyToSectionRequest{
	// ...
}


read, err := client.CreateMeOnenotePageByIdCopyToSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenotePageClient.CreateMeOnenotePageByIdOnenotePatchContent`

```go
ctx := context.TODO()
id := meonenotepage.NewMeOnenotePageID("onenotePageIdValue")

payload := meonenotepage.CreateMeOnenotePageByIdOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateMeOnenotePageByIdOnenotePatchContent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenotePageClient.DeleteMeOnenotePageById`

```go
ctx := context.TODO()
id := meonenotepage.NewMeOnenotePageID("onenotePageIdValue")

read, err := client.DeleteMeOnenotePageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenotePageClient.GetMeOnenotePageById`

```go
ctx := context.TODO()
id := meonenotepage.NewMeOnenotePageID("onenotePageIdValue")

read, err := client.GetMeOnenotePageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenotePageClient.GetMeOnenotePageCount`

```go
ctx := context.TODO()


read, err := client.GetMeOnenotePageCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnenotePageClient.ListMeOnenotePages`

```go
ctx := context.TODO()


// alternatively `client.ListMeOnenotePages(ctx)` can be used to do batched pagination
items, err := client.ListMeOnenotePagesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeOnenotePageClient.UpdateMeOnenotePageById`

```go
ctx := context.TODO()
id := meonenotepage.NewMeOnenotePageID("onenotePageIdValue")

payload := meonenotepage.OnenotePage{
	// ...
}


read, err := client.UpdateMeOnenotePageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
