
## `github.com/hashicorp/go-azure-sdk/resource-manager/videoanalyzer/2021-05-01-preview/videos` Documentation

The `videos` SDK allows for interaction with Azure Resource Manager `videoanalyzer` (API Version `2021-05-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/videoanalyzer/2021-05-01-preview/videos"
```


### Client Initialization

```go
client := videos.NewVideosClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VideosClient.AccessPoliciesCreateOrUpdate`

```go
ctx := context.TODO()
id := videos.NewAccessPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "accessPolicyName")

payload := videos.AccessPolicyEntity{
	// ...
}


read, err := client.AccessPoliciesCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VideosClient.AccessPoliciesDelete`

```go
ctx := context.TODO()
id := videos.NewAccessPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "accessPolicyName")

read, err := client.AccessPoliciesDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VideosClient.AccessPoliciesGet`

```go
ctx := context.TODO()
id := videos.NewAccessPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "accessPolicyName")

read, err := client.AccessPoliciesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VideosClient.AccessPoliciesList`

```go
ctx := context.TODO()
id := videos.NewVideoAnalyzerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName")

// alternatively `client.AccessPoliciesList(ctx, id, videos.DefaultAccessPoliciesListOperationOptions())` can be used to do batched pagination
items, err := client.AccessPoliciesListComplete(ctx, id, videos.DefaultAccessPoliciesListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VideosClient.AccessPoliciesUpdate`

```go
ctx := context.TODO()
id := videos.NewAccessPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "accessPolicyName")

payload := videos.AccessPolicyEntity{
	// ...
}


read, err := client.AccessPoliciesUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VideosClient.VideosCreateOrUpdate`

```go
ctx := context.TODO()
id := videos.NewVideoID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "videoName")

payload := videos.VideoEntity{
	// ...
}


read, err := client.VideosCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VideosClient.VideosDelete`

```go
ctx := context.TODO()
id := videos.NewVideoID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "videoName")

read, err := client.VideosDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VideosClient.VideosGet`

```go
ctx := context.TODO()
id := videos.NewVideoID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "videoName")

read, err := client.VideosGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VideosClient.VideosList`

```go
ctx := context.TODO()
id := videos.NewVideoAnalyzerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName")

// alternatively `client.VideosList(ctx, id, videos.DefaultVideosListOperationOptions())` can be used to do batched pagination
items, err := client.VideosListComplete(ctx, id, videos.DefaultVideosListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VideosClient.VideosListStreamingToken`

```go
ctx := context.TODO()
id := videos.NewVideoID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "videoName")

read, err := client.VideosListStreamingToken(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VideosClient.VideosUpdate`

```go
ctx := context.TODO()
id := videos.NewVideoID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "videoName")

payload := videos.VideoEntity{
	// ...
}


read, err := client.VideosUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
