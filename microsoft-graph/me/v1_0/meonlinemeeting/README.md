
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/meonlinemeeting` Documentation

The `meonlinemeeting` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/meonlinemeeting"
```


### Client Initialization

```go
client := meonlinemeeting.NewMeOnlineMeetingClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeOnlineMeetingClient.CreateMeOnlineMeeting`

```go
ctx := context.TODO()

payload := meonlinemeeting.OnlineMeeting{
	// ...
}


read, err := client.CreateMeOnlineMeeting(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnlineMeetingClient.CreateMeOnlineMeetingCreateOrGet`

```go
ctx := context.TODO()

payload := meonlinemeeting.CreateMeOnlineMeetingCreateOrGetRequest{
	// ...
}


read, err := client.CreateMeOnlineMeetingCreateOrGet(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnlineMeetingClient.DeleteMeOnlineMeetingById`

```go
ctx := context.TODO()
id := meonlinemeeting.NewMeOnlineMeetingID("onlineMeetingIdValue")

read, err := client.DeleteMeOnlineMeetingById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnlineMeetingClient.GetMeOnlineMeetingById`

```go
ctx := context.TODO()
id := meonlinemeeting.NewMeOnlineMeetingID("onlineMeetingIdValue")

read, err := client.GetMeOnlineMeetingById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnlineMeetingClient.GetMeOnlineMeetingCount`

```go
ctx := context.TODO()


read, err := client.GetMeOnlineMeetingCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeOnlineMeetingClient.ListMeOnlineMeetings`

```go
ctx := context.TODO()


// alternatively `client.ListMeOnlineMeetings(ctx)` can be used to do batched pagination
items, err := client.ListMeOnlineMeetingsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeOnlineMeetingClient.UpdateMeOnlineMeetingById`

```go
ctx := context.TODO()
id := meonlinemeeting.NewMeOnlineMeetingID("onlineMeetingIdValue")

payload := meonlinemeeting.OnlineMeeting{
	// ...
}


read, err := client.UpdateMeOnlineMeetingById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
