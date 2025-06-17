
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamscheduletimecard` Documentation

The `teamscheduletimecard` SDK allows for interaction with Microsoft Graph `groups` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/stable/teamscheduletimecard"
```


### Client Initialization

```go
client := teamscheduletimecard.NewTeamScheduleTimeCardClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamScheduleTimeCardClient.CreateTeamScheduleTimeCard`

```go
ctx := context.TODO()
id := teamscheduletimecard.NewGroupID("groupId")

payload := teamscheduletimecard.TimeCard{
	// ...
}


read, err := client.CreateTeamScheduleTimeCard(ctx, id, payload, teamscheduletimecard.DefaultCreateTeamScheduleTimeCardOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamScheduleTimeCardClient.CreateTeamScheduleTimeCardClockIn`

```go
ctx := context.TODO()
id := teamscheduletimecard.NewGroupID("groupId")

payload := teamscheduletimecard.CreateTeamScheduleTimeCardClockInRequest{
	// ...
}


read, err := client.CreateTeamScheduleTimeCardClockIn(ctx, id, payload, teamscheduletimecard.DefaultCreateTeamScheduleTimeCardClockInOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamScheduleTimeCardClient.CreateTeamScheduleTimeCardClockOut`

```go
ctx := context.TODO()
id := teamscheduletimecard.NewGroupIdTeamScheduleTimeCardID("groupId", "timeCardId")

payload := teamscheduletimecard.CreateTeamScheduleTimeCardClockOutRequest{
	// ...
}


read, err := client.CreateTeamScheduleTimeCardClockOut(ctx, id, payload, teamscheduletimecard.DefaultCreateTeamScheduleTimeCardClockOutOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamScheduleTimeCardClient.CreateTeamScheduleTimeCardConfirm`

```go
ctx := context.TODO()
id := teamscheduletimecard.NewGroupIdTeamScheduleTimeCardID("groupId", "timeCardId")

read, err := client.CreateTeamScheduleTimeCardConfirm(ctx, id, teamscheduletimecard.DefaultCreateTeamScheduleTimeCardConfirmOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamScheduleTimeCardClient.DeleteTeamScheduleTimeCard`

```go
ctx := context.TODO()
id := teamscheduletimecard.NewGroupIdTeamScheduleTimeCardID("groupId", "timeCardId")

read, err := client.DeleteTeamScheduleTimeCard(ctx, id, teamscheduletimecard.DefaultDeleteTeamScheduleTimeCardOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamScheduleTimeCardClient.EndTeamScheduleTimeCardBreak`

```go
ctx := context.TODO()
id := teamscheduletimecard.NewGroupIdTeamScheduleTimeCardID("groupId", "timeCardId")

payload := teamscheduletimecard.EndTeamScheduleTimeCardBreakRequest{
	// ...
}


read, err := client.EndTeamScheduleTimeCardBreak(ctx, id, payload, teamscheduletimecard.DefaultEndTeamScheduleTimeCardBreakOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamScheduleTimeCardClient.GetTeamScheduleTimeCard`

```go
ctx := context.TODO()
id := teamscheduletimecard.NewGroupIdTeamScheduleTimeCardID("groupId", "timeCardId")

read, err := client.GetTeamScheduleTimeCard(ctx, id, teamscheduletimecard.DefaultGetTeamScheduleTimeCardOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamScheduleTimeCardClient.GetTeamScheduleTimeCardsCount`

```go
ctx := context.TODO()
id := teamscheduletimecard.NewGroupID("groupId")

read, err := client.GetTeamScheduleTimeCardsCount(ctx, id, teamscheduletimecard.DefaultGetTeamScheduleTimeCardsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamScheduleTimeCardClient.ListTeamScheduleTimeCards`

```go
ctx := context.TODO()
id := teamscheduletimecard.NewGroupID("groupId")

// alternatively `client.ListTeamScheduleTimeCards(ctx, id, teamscheduletimecard.DefaultListTeamScheduleTimeCardsOperationOptions())` can be used to do batched pagination
items, err := client.ListTeamScheduleTimeCardsComplete(ctx, id, teamscheduletimecard.DefaultListTeamScheduleTimeCardsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamScheduleTimeCardClient.StartTeamScheduleTimeCardBreak`

```go
ctx := context.TODO()
id := teamscheduletimecard.NewGroupIdTeamScheduleTimeCardID("groupId", "timeCardId")

payload := teamscheduletimecard.StartTeamScheduleTimeCardBreakRequest{
	// ...
}


read, err := client.StartTeamScheduleTimeCardBreak(ctx, id, payload, teamscheduletimecard.DefaultStartTeamScheduleTimeCardBreakOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamScheduleTimeCardClient.UpdateTeamScheduleTimeCard`

```go
ctx := context.TODO()
id := teamscheduletimecard.NewGroupIdTeamScheduleTimeCardID("groupId", "timeCardId")

payload := teamscheduletimecard.TimeCard{
	// ...
}


read, err := client.UpdateTeamScheduleTimeCard(ctx, id, payload, teamscheduletimecard.DefaultUpdateTeamScheduleTimeCardOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
