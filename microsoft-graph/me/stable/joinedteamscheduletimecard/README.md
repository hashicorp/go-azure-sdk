
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/joinedteamscheduletimecard` Documentation

The `joinedteamscheduletimecard` SDK allows for interaction with Microsoft Graph `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/joinedteamscheduletimecard"
```


### Client Initialization

```go
client := joinedteamscheduletimecard.NewJoinedTeamScheduleTimeCardClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JoinedTeamScheduleTimeCardClient.CreateJoinedTeamScheduleTimeCard`

```go
ctx := context.TODO()
id := joinedteamscheduletimecard.NewMeJoinedTeamID("teamId")

payload := joinedteamscheduletimecard.TimeCard{
	// ...
}


read, err := client.CreateJoinedTeamScheduleTimeCard(ctx, id, payload, joinedteamscheduletimecard.DefaultCreateJoinedTeamScheduleTimeCardOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamScheduleTimeCardClient.CreateJoinedTeamScheduleTimeCardClockIn`

```go
ctx := context.TODO()
id := joinedteamscheduletimecard.NewMeJoinedTeamID("teamId")

payload := joinedteamscheduletimecard.CreateJoinedTeamScheduleTimeCardClockInRequest{
	// ...
}


read, err := client.CreateJoinedTeamScheduleTimeCardClockIn(ctx, id, payload, joinedteamscheduletimecard.DefaultCreateJoinedTeamScheduleTimeCardClockInOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamScheduleTimeCardClient.CreateJoinedTeamScheduleTimeCardClockOut`

```go
ctx := context.TODO()
id := joinedteamscheduletimecard.NewMeJoinedTeamIdScheduleTimeCardID("teamId", "timeCardId")

payload := joinedteamscheduletimecard.CreateJoinedTeamScheduleTimeCardClockOutRequest{
	// ...
}


read, err := client.CreateJoinedTeamScheduleTimeCardClockOut(ctx, id, payload, joinedteamscheduletimecard.DefaultCreateJoinedTeamScheduleTimeCardClockOutOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamScheduleTimeCardClient.CreateJoinedTeamScheduleTimeCardConfirm`

```go
ctx := context.TODO()
id := joinedteamscheduletimecard.NewMeJoinedTeamIdScheduleTimeCardID("teamId", "timeCardId")

read, err := client.CreateJoinedTeamScheduleTimeCardConfirm(ctx, id, joinedteamscheduletimecard.DefaultCreateJoinedTeamScheduleTimeCardConfirmOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamScheduleTimeCardClient.DeleteJoinedTeamScheduleTimeCard`

```go
ctx := context.TODO()
id := joinedteamscheduletimecard.NewMeJoinedTeamIdScheduleTimeCardID("teamId", "timeCardId")

read, err := client.DeleteJoinedTeamScheduleTimeCard(ctx, id, joinedteamscheduletimecard.DefaultDeleteJoinedTeamScheduleTimeCardOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamScheduleTimeCardClient.EndJoinedTeamScheduleTimeCardBreak`

```go
ctx := context.TODO()
id := joinedteamscheduletimecard.NewMeJoinedTeamIdScheduleTimeCardID("teamId", "timeCardId")

payload := joinedteamscheduletimecard.EndJoinedTeamScheduleTimeCardBreakRequest{
	// ...
}


read, err := client.EndJoinedTeamScheduleTimeCardBreak(ctx, id, payload, joinedteamscheduletimecard.DefaultEndJoinedTeamScheduleTimeCardBreakOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamScheduleTimeCardClient.GetJoinedTeamScheduleTimeCard`

```go
ctx := context.TODO()
id := joinedteamscheduletimecard.NewMeJoinedTeamIdScheduleTimeCardID("teamId", "timeCardId")

read, err := client.GetJoinedTeamScheduleTimeCard(ctx, id, joinedteamscheduletimecard.DefaultGetJoinedTeamScheduleTimeCardOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamScheduleTimeCardClient.GetJoinedTeamScheduleTimeCardsCount`

```go
ctx := context.TODO()
id := joinedteamscheduletimecard.NewMeJoinedTeamID("teamId")

read, err := client.GetJoinedTeamScheduleTimeCardsCount(ctx, id, joinedteamscheduletimecard.DefaultGetJoinedTeamScheduleTimeCardsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamScheduleTimeCardClient.ListJoinedTeamScheduleTimeCards`

```go
ctx := context.TODO()
id := joinedteamscheduletimecard.NewMeJoinedTeamID("teamId")

// alternatively `client.ListJoinedTeamScheduleTimeCards(ctx, id, joinedteamscheduletimecard.DefaultListJoinedTeamScheduleTimeCardsOperationOptions())` can be used to do batched pagination
items, err := client.ListJoinedTeamScheduleTimeCardsComplete(ctx, id, joinedteamscheduletimecard.DefaultListJoinedTeamScheduleTimeCardsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `JoinedTeamScheduleTimeCardClient.StartJoinedTeamScheduleTimeCardBreak`

```go
ctx := context.TODO()
id := joinedteamscheduletimecard.NewMeJoinedTeamIdScheduleTimeCardID("teamId", "timeCardId")

payload := joinedteamscheduletimecard.StartJoinedTeamScheduleTimeCardBreakRequest{
	// ...
}


read, err := client.StartJoinedTeamScheduleTimeCardBreak(ctx, id, payload, joinedteamscheduletimecard.DefaultStartJoinedTeamScheduleTimeCardBreakOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JoinedTeamScheduleTimeCardClient.UpdateJoinedTeamScheduleTimeCard`

```go
ctx := context.TODO()
id := joinedteamscheduletimecard.NewMeJoinedTeamIdScheduleTimeCardID("teamId", "timeCardId")

payload := joinedteamscheduletimecard.TimeCard{
	// ...
}


read, err := client.UpdateJoinedTeamScheduleTimeCard(ctx, id, payload, joinedteamscheduletimecard.DefaultUpdateJoinedTeamScheduleTimeCardOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
