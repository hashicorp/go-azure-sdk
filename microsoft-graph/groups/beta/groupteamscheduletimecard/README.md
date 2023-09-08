
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupteamscheduletimecard` Documentation

The `groupteamscheduletimecard` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupteamscheduletimecard"
```


### Client Initialization

```go
client := groupteamscheduletimecard.NewGroupTeamScheduleTimeCardClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupTeamScheduleTimeCardClient.CreateGroupByIdTeamScheduleTimeCard`

```go
ctx := context.TODO()
id := groupteamscheduletimecard.NewGroupID("groupIdValue")

payload := groupteamscheduletimecard.TimeCard{
	// ...
}


read, err := client.CreateGroupByIdTeamScheduleTimeCard(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamScheduleTimeCardClient.CreateGroupByIdTeamScheduleTimeCardByIdClockOut`

```go
ctx := context.TODO()
id := groupteamscheduletimecard.NewGroupTeamScheduleTimeCardID("groupIdValue", "timeCardIdValue")

payload := groupteamscheduletimecard.CreateGroupByIdTeamScheduleTimeCardByIdClockOutRequest{
	// ...
}


read, err := client.CreateGroupByIdTeamScheduleTimeCardByIdClockOut(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamScheduleTimeCardClient.CreateGroupByIdTeamScheduleTimeCardByIdConfirm`

```go
ctx := context.TODO()
id := groupteamscheduletimecard.NewGroupTeamScheduleTimeCardID("groupIdValue", "timeCardIdValue")

read, err := client.CreateGroupByIdTeamScheduleTimeCardByIdConfirm(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamScheduleTimeCardClient.CreateGroupByIdTeamScheduleTimeCardByIdEndBreak`

```go
ctx := context.TODO()
id := groupteamscheduletimecard.NewGroupTeamScheduleTimeCardID("groupIdValue", "timeCardIdValue")

payload := groupteamscheduletimecard.CreateGroupByIdTeamScheduleTimeCardByIdEndBreakRequest{
	// ...
}


read, err := client.CreateGroupByIdTeamScheduleTimeCardByIdEndBreak(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamScheduleTimeCardClient.CreateGroupByIdTeamScheduleTimeCardClockIn`

```go
ctx := context.TODO()
id := groupteamscheduletimecard.NewGroupID("groupIdValue")

payload := groupteamscheduletimecard.CreateGroupByIdTeamScheduleTimeCardClockInRequest{
	// ...
}


read, err := client.CreateGroupByIdTeamScheduleTimeCardClockIn(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamScheduleTimeCardClient.DeleteGroupByIdTeamScheduleTimeCardById`

```go
ctx := context.TODO()
id := groupteamscheduletimecard.NewGroupTeamScheduleTimeCardID("groupIdValue", "timeCardIdValue")

read, err := client.DeleteGroupByIdTeamScheduleTimeCardById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamScheduleTimeCardClient.GetGroupByIdTeamScheduleTimeCardById`

```go
ctx := context.TODO()
id := groupteamscheduletimecard.NewGroupTeamScheduleTimeCardID("groupIdValue", "timeCardIdValue")

read, err := client.GetGroupByIdTeamScheduleTimeCardById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamScheduleTimeCardClient.GetGroupByIdTeamScheduleTimeCardCount`

```go
ctx := context.TODO()
id := groupteamscheduletimecard.NewGroupID("groupIdValue")

read, err := client.GetGroupByIdTeamScheduleTimeCardCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamScheduleTimeCardClient.ListGroupByIdTeamScheduleTimeCards`

```go
ctx := context.TODO()
id := groupteamscheduletimecard.NewGroupID("groupIdValue")

// alternatively `client.ListGroupByIdTeamScheduleTimeCards(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdTeamScheduleTimeCardsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupTeamScheduleTimeCardClient.StartGroupByIdTeamScheduleTimeCardByIdBreak`

```go
ctx := context.TODO()
id := groupteamscheduletimecard.NewGroupTeamScheduleTimeCardID("groupIdValue", "timeCardIdValue")

payload := groupteamscheduletimecard.StartGroupByIdTeamScheduleTimeCardByIdBreakRequest{
	// ...
}


read, err := client.StartGroupByIdTeamScheduleTimeCardByIdBreak(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupTeamScheduleTimeCardClient.UpdateGroupByIdTeamScheduleTimeCardById`

```go
ctx := context.TODO()
id := groupteamscheduletimecard.NewGroupTeamScheduleTimeCardID("groupIdValue", "timeCardIdValue")

payload := groupteamscheduletimecard.TimeCard{
	// ...
}


read, err := client.UpdateGroupByIdTeamScheduleTimeCardById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
