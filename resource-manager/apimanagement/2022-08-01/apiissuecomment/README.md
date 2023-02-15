
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/apiissuecomment` Documentation

The `apiissuecomment` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2022-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/apiissuecomment"
```


### Client Initialization

```go
client := apiissuecomment.NewApiIssueCommentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApiIssueCommentClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := apiissuecomment.NewCommentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "apiIdValue", "issueIdValue", "commentIdValue")

payload := apiissuecomment.IssueCommentContract{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload, apiissuecomment.DefaultCreateOrUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApiIssueCommentClient.Delete`

```go
ctx := context.TODO()
id := apiissuecomment.NewCommentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "apiIdValue", "issueIdValue", "commentIdValue")

read, err := client.Delete(ctx, id, apiissuecomment.DefaultDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApiIssueCommentClient.Get`

```go
ctx := context.TODO()
id := apiissuecomment.NewCommentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "apiIdValue", "issueIdValue", "commentIdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApiIssueCommentClient.GetEntityTag`

```go
ctx := context.TODO()
id := apiissuecomment.NewCommentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "apiIdValue", "issueIdValue", "commentIdValue")

read, err := client.GetEntityTag(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApiIssueCommentClient.ListByService`

```go
ctx := context.TODO()
id := apiissuecomment.NewIssueID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "issueIdValue")

// alternatively `client.ListByService(ctx, id, apiissuecomment.DefaultListByServiceOperationOptions())` can be used to do batched pagination
items, err := client.ListByServiceComplete(ctx, id, apiissuecomment.DefaultListByServiceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
