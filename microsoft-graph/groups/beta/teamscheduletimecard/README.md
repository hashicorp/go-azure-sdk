
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamscheduletimecard` Documentation

The `teamscheduletimecard` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/teamscheduletimecard"
```


### Client Initialization

```go
client := teamscheduletimecard.NewTeamScheduleTimeCardClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TeamScheduleTimeCardClient.CreateTeamScheduleTimeCard`

```go
ctx := context.TODO()
id := teamscheduletimecard.NewGroupID("groupIdValue")

payload := teamscheduletimecard.TimeCard{
	// ...
}


read, err := client.CreateTeamScheduleTimeCard(ctx, id, payload)
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
id := teamscheduletimecard.NewGroupID("groupIdValue")

payload := teamscheduletimecard.CreateTeamScheduleTimeCardClockInRequest{
	// ...
}


read, err := client.CreateTeamScheduleTimeCardClockIn(ctx, id, payload)
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
id := teamscheduletimecard.NewGroupIdTeamScheduleTimeCardID("groupIdValue", "timeCardIdValue")

payload := teamscheduletimecard.CreateTeamScheduleTimeCardClockOutRequest{
	// ...
}


read, err := client.CreateTeamScheduleTimeCardClockOut(ctx, id, payload)
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
id := teamscheduletimecard.NewGroupIdTeamScheduleTimeCardID("groupIdValue", "timeCardIdValue")

read, err := client.CreateTeamScheduleTimeCardConfirm(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamScheduleTimeCardClient.CreateTeamScheduleTimeCardEndBreak`

```go
ctx := context.TODO()
id := teamscheduletimecard.NewGroupIdTeamScheduleTimeCardID("groupIdValue", "timeCardIdValue")

payload := teamscheduletimecard.CreateTeamScheduleTimeCardEndBreakRequest{
	// ...
}


read, err := client.CreateTeamScheduleTimeCardEndBreak(ctx, id, payload)
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
id := teamscheduletimecard.NewGroupIdTeamScheduleTimeCardID("groupIdValue", "timeCardIdValue")

read, err := client.DeleteTeamScheduleTimeCard(ctx, id)
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
id := teamscheduletimecard.NewGroupIdTeamScheduleTimeCardID("groupIdValue", "timeCardIdValue")

read, err := client.GetTeamScheduleTimeCard(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TeamScheduleTimeCardClient.GetTeamScheduleTimeCardCount`

```go
ctx := context.TODO()
id := teamscheduletimecard.NewGroupID("groupIdValue")

read, err := client.GetTeamScheduleTimeCardCount(ctx, id)
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
id := teamscheduletimecard.NewGroupID("groupIdValue")

// alternatively `client.ListTeamScheduleTimeCards(ctx, id)` can be used to do batched pagination
items, err := client.ListTeamScheduleTimeCardsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TeamScheduleTimeCardClient.StartGroupTeamScheduleTimeCardBreak`

```go
ctx := context.TODO()
id := teamscheduletimecard.NewGroupIdTeamScheduleTimeCardID("groupIdValue", "timeCardIdValue")

payload := teamscheduletimecard.StartGroupTeamScheduleTimeCardBreakRequest{
	// ...
}


read, err := client.StartGroupTeamScheduleTimeCardBreak(ctx, id, payload)
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
id := teamscheduletimecard.NewGroupIdTeamScheduleTimeCardID("groupIdValue", "timeCardIdValue")

payload := teamscheduletimecard.TimeCard{
	// ...
}


read, err := client.UpdateTeamScheduleTimeCard(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
