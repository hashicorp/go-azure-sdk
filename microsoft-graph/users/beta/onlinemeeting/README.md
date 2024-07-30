
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onlinemeeting` Documentation

The `onlinemeeting` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/onlinemeeting"
```


### Client Initialization

```go
client := onlinemeeting.NewOnlineMeetingClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OnlineMeetingClient.CreateOnlineMeeting`

```go
ctx := context.TODO()
id := onlinemeeting.NewUserID("userIdValue")

payload := onlinemeeting.OnlineMeeting{
	// ...
}


read, err := client.CreateOnlineMeeting(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnlineMeetingClient.CreateOnlineMeetingCreateOrGet`

```go
ctx := context.TODO()
id := onlinemeeting.NewUserID("userIdValue")

payload := onlinemeeting.CreateOnlineMeetingCreateOrGetRequest{
	// ...
}


read, err := client.CreateOnlineMeetingCreateOrGet(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnlineMeetingClient.CreateOnlineMeetingSendVirtualAppointmentReminderSm`

```go
ctx := context.TODO()
id := onlinemeeting.NewUserIdOnlineMeetingID("userIdValue", "onlineMeetingIdValue")

payload := onlinemeeting.CreateOnlineMeetingSendVirtualAppointmentReminderSmRequest{
	// ...
}


read, err := client.CreateOnlineMeetingSendVirtualAppointmentReminderSm(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnlineMeetingClient.CreateOnlineMeetingSendVirtualAppointmentSm`

```go
ctx := context.TODO()
id := onlinemeeting.NewUserIdOnlineMeetingID("userIdValue", "onlineMeetingIdValue")

payload := onlinemeeting.CreateOnlineMeetingSendVirtualAppointmentSmRequest{
	// ...
}


read, err := client.CreateOnlineMeetingSendVirtualAppointmentSm(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnlineMeetingClient.DeleteOnlineMeeting`

```go
ctx := context.TODO()
id := onlinemeeting.NewUserIdOnlineMeetingID("userIdValue", "onlineMeetingIdValue")

read, err := client.DeleteOnlineMeeting(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnlineMeetingClient.GetOnlineMeeting`

```go
ctx := context.TODO()
id := onlinemeeting.NewUserIdOnlineMeetingID("userIdValue", "onlineMeetingIdValue")

read, err := client.GetOnlineMeeting(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnlineMeetingClient.GetOnlineMeetingCount`

```go
ctx := context.TODO()
id := onlinemeeting.NewUserID("userIdValue")

read, err := client.GetOnlineMeetingCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnlineMeetingClient.ListOnlineMeetings`

```go
ctx := context.TODO()
id := onlinemeeting.NewUserID("userIdValue")

// alternatively `client.ListOnlineMeetings(ctx, id)` can be used to do batched pagination
items, err := client.ListOnlineMeetingsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OnlineMeetingClient.UpdateOnlineMeeting`

```go
ctx := context.TODO()
id := onlinemeeting.NewUserIdOnlineMeetingID("userIdValue", "onlineMeetingIdValue")

payload := onlinemeeting.OnlineMeeting{
	// ...
}


read, err := client.UpdateOnlineMeeting(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
