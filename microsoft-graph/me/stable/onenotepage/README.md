
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/onenotepage` Documentation

The `onenotepage` SDK allows for interaction with Microsoft Graph `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/onenotepage"
```


### Client Initialization

```go
client := onenotepage.NewOnenotePageClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OnenotePageClient.CopyOnenotePageToSection`

```go
ctx := context.TODO()
id := onenotepage.NewMeOnenotePageID("onenotePageId")

payload := onenotepage.CopyOnenotePageToSectionRequest{
	// ...
}


read, err := client.CopyOnenotePageToSection(ctx, id, payload, onenotepage.DefaultCopyOnenotePageToSectionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenotePageClient.CreateOnenotePage`

```go
ctx := context.TODO()

payload := onenotepage.OnenotePage{
	// ...
}


read, err := client.CreateOnenotePage(ctx, payload, onenotepage.DefaultCreateOnenotePageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenotePageClient.CreateOnenotePageOnenotePatchContent`

```go
ctx := context.TODO()
id := onenotepage.NewMeOnenotePageID("onenotePageId")

payload := onenotepage.CreateOnenotePageOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateOnenotePageOnenotePatchContent(ctx, id, payload, onenotepage.DefaultCreateOnenotePageOnenotePatchContentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenotePageClient.DeleteOnenotePage`

```go
ctx := context.TODO()
id := onenotepage.NewMeOnenotePageID("onenotePageId")

read, err := client.DeleteOnenotePage(ctx, id, onenotepage.DefaultDeleteOnenotePageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenotePageClient.GetOnenotePage`

```go
ctx := context.TODO()
id := onenotepage.NewMeOnenotePageID("onenotePageId")

read, err := client.GetOnenotePage(ctx, id, onenotepage.DefaultGetOnenotePageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenotePageClient.GetOnenotePagesCount`

```go
ctx := context.TODO()


read, err := client.GetOnenotePagesCount(ctx, onenotepage.DefaultGetOnenotePagesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnenotePageClient.ListOnenotePages`

```go
ctx := context.TODO()


// alternatively `client.ListOnenotePages(ctx, onenotepage.DefaultListOnenotePagesOperationOptions())` can be used to do batched pagination
items, err := client.ListOnenotePagesComplete(ctx, onenotepage.DefaultListOnenotePagesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OnenotePageClient.UpdateOnenotePage`

```go
ctx := context.TODO()
id := onenotepage.NewMeOnenotePageID("onenotePageId")

payload := onenotepage.OnenotePage{
	// ...
}


read, err := client.UpdateOnenotePage(ctx, id, payload, onenotepage.DefaultUpdateOnenotePageOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
