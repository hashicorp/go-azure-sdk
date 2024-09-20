
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/onlinemeeting` Documentation

The `onlinemeeting` SDK allows for interaction with Microsoft Graph `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/onlinemeeting"
```


### Client Initialization

```go
client := onlinemeeting.NewOnlineMeetingClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OnlineMeetingClient.CreateOnlineMeeting`

```go
ctx := context.TODO()

payload := onlinemeeting.OnlineMeeting{
	// ...
}


read, err := client.CreateOnlineMeeting(ctx, payload, onlinemeeting.DefaultCreateOnlineMeetingOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnlineMeetingClient.CreateOrGetOnlineMeetings`

```go
ctx := context.TODO()

payload := onlinemeeting.CreateOrGetOnlineMeetingsRequest{
	// ...
}


read, err := client.CreateOrGetOnlineMeetings(ctx, payload, onlinemeeting.DefaultCreateOrGetOnlineMeetingsOperationOptions())
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
id := onlinemeeting.NewMeOnlineMeetingID("onlineMeetingId")

read, err := client.DeleteOnlineMeeting(ctx, id, onlinemeeting.DefaultDeleteOnlineMeetingOperationOptions())
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
id := onlinemeeting.NewMeOnlineMeetingID("onlineMeetingId")

read, err := client.GetOnlineMeeting(ctx, id, onlinemeeting.DefaultGetOnlineMeetingOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnlineMeetingClient.GetOnlineMeetingsCount`

```go
ctx := context.TODO()


read, err := client.GetOnlineMeetingsCount(ctx, onlinemeeting.DefaultGetOnlineMeetingsCountOperationOptions())
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


// alternatively `client.ListOnlineMeetings(ctx, onlinemeeting.DefaultListOnlineMeetingsOperationOptions())` can be used to do batched pagination
items, err := client.ListOnlineMeetingsComplete(ctx, onlinemeeting.DefaultListOnlineMeetingsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OnlineMeetingClient.SendOnlineMeetingVirtualAppointmentReminderSms`

```go
ctx := context.TODO()
id := onlinemeeting.NewMeOnlineMeetingID("onlineMeetingId")

payload := onlinemeeting.SendOnlineMeetingVirtualAppointmentReminderSmsRequest{
	// ...
}


read, err := client.SendOnlineMeetingVirtualAppointmentReminderSms(ctx, id, payload, onlinemeeting.DefaultSendOnlineMeetingVirtualAppointmentReminderSmsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnlineMeetingClient.SendOnlineMeetingVirtualAppointmentSms`

```go
ctx := context.TODO()
id := onlinemeeting.NewMeOnlineMeetingID("onlineMeetingId")

payload := onlinemeeting.SendOnlineMeetingVirtualAppointmentSmsRequest{
	// ...
}


read, err := client.SendOnlineMeetingVirtualAppointmentSms(ctx, id, payload, onlinemeeting.DefaultSendOnlineMeetingVirtualAppointmentSmsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OnlineMeetingClient.UpdateOnlineMeeting`

```go
ctx := context.TODO()
id := onlinemeeting.NewMeOnlineMeetingID("onlineMeetingId")

payload := onlinemeeting.OnlineMeeting{
	// ...
}


read, err := client.UpdateOnlineMeeting(ctx, id, payload, onlinemeeting.DefaultUpdateOnlineMeetingOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
