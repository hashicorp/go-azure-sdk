
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mejoinedteampermissiongrant` Documentation

The `mejoinedteampermissiongrant` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mejoinedteampermissiongrant"
```


### Client Initialization

```go
client := mejoinedteampermissiongrant.NewMeJoinedTeamPermissionGrantClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeJoinedTeamPermissionGrantClient.CheckMeJoinedTeamByIdPermissionGrantByIdMemberGroup`

```go
ctx := context.TODO()
id := mejoinedteampermissiongrant.NewMeJoinedTeamPermissionGrantID("teamIdValue", "resourceSpecificPermissionGrantIdValue")

payload := mejoinedteampermissiongrant.CheckMeJoinedTeamByIdPermissionGrantByIdMemberGroupRequest{
	// ...
}


read, err := client.CheckMeJoinedTeamByIdPermissionGrantByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPermissionGrantClient.CheckMeJoinedTeamByIdPermissionGrantByIdMemberObject`

```go
ctx := context.TODO()
id := mejoinedteampermissiongrant.NewMeJoinedTeamPermissionGrantID("teamIdValue", "resourceSpecificPermissionGrantIdValue")

payload := mejoinedteampermissiongrant.CheckMeJoinedTeamByIdPermissionGrantByIdMemberObjectRequest{
	// ...
}


read, err := client.CheckMeJoinedTeamByIdPermissionGrantByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPermissionGrantClient.CreateMeJoinedTeamByIdPermissionGrant`

```go
ctx := context.TODO()
id := mejoinedteampermissiongrant.NewMeJoinedTeamID("teamIdValue")

payload := mejoinedteampermissiongrant.ResourceSpecificPermissionGrant{
	// ...
}


read, err := client.CreateMeJoinedTeamByIdPermissionGrant(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPermissionGrantClient.DeleteMeJoinedTeamByIdPermissionGrantById`

```go
ctx := context.TODO()
id := mejoinedteampermissiongrant.NewMeJoinedTeamPermissionGrantID("teamIdValue", "resourceSpecificPermissionGrantIdValue")

read, err := client.DeleteMeJoinedTeamByIdPermissionGrantById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPermissionGrantClient.GetMeJoinedTeamByIdPermissionGrantById`

```go
ctx := context.TODO()
id := mejoinedteampermissiongrant.NewMeJoinedTeamPermissionGrantID("teamIdValue", "resourceSpecificPermissionGrantIdValue")

read, err := client.GetMeJoinedTeamByIdPermissionGrantById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPermissionGrantClient.GetMeJoinedTeamByIdPermissionGrantByIdMemberGroup`

```go
ctx := context.TODO()
id := mejoinedteampermissiongrant.NewMeJoinedTeamPermissionGrantID("teamIdValue", "resourceSpecificPermissionGrantIdValue")

payload := mejoinedteampermissiongrant.GetMeJoinedTeamByIdPermissionGrantByIdMemberGroupRequest{
	// ...
}


read, err := client.GetMeJoinedTeamByIdPermissionGrantByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPermissionGrantClient.GetMeJoinedTeamByIdPermissionGrantByIdMemberObject`

```go
ctx := context.TODO()
id := mejoinedteampermissiongrant.NewMeJoinedTeamPermissionGrantID("teamIdValue", "resourceSpecificPermissionGrantIdValue")

payload := mejoinedteampermissiongrant.GetMeJoinedTeamByIdPermissionGrantByIdMemberObjectRequest{
	// ...
}


read, err := client.GetMeJoinedTeamByIdPermissionGrantByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPermissionGrantClient.GetMeJoinedTeamByIdPermissionGrantCount`

```go
ctx := context.TODO()
id := mejoinedteampermissiongrant.NewMeJoinedTeamID("teamIdValue")

read, err := client.GetMeJoinedTeamByIdPermissionGrantCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPermissionGrantClient.GetMeJoinedTeamByIdPermissionGrantsAvailableExtensionProperties`

```go
ctx := context.TODO()
id := mejoinedteampermissiongrant.NewMeJoinedTeamID("teamIdValue")

// alternatively `client.GetMeJoinedTeamByIdPermissionGrantsAvailableExtensionProperties(ctx, id)` can be used to do batched pagination
items, err := client.GetMeJoinedTeamByIdPermissionGrantsAvailableExtensionPropertiesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeJoinedTeamPermissionGrantClient.GetMeJoinedTeamByIdPermissionGrantsByIds`

```go
ctx := context.TODO()
id := mejoinedteampermissiongrant.NewMeJoinedTeamID("teamIdValue")

// alternatively `client.GetMeJoinedTeamByIdPermissionGrantsByIds(ctx, id)` can be used to do batched pagination
items, err := client.GetMeJoinedTeamByIdPermissionGrantsByIdsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeJoinedTeamPermissionGrantClient.ListMeJoinedTeamByIdPermissionGrants`

```go
ctx := context.TODO()
id := mejoinedteampermissiongrant.NewMeJoinedTeamID("teamIdValue")

// alternatively `client.ListMeJoinedTeamByIdPermissionGrants(ctx, id)` can be used to do batched pagination
items, err := client.ListMeJoinedTeamByIdPermissionGrantsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeJoinedTeamPermissionGrantClient.RestoreMeJoinedTeamByIdPermissionGrantById`

```go
ctx := context.TODO()
id := mejoinedteampermissiongrant.NewMeJoinedTeamPermissionGrantID("teamIdValue", "resourceSpecificPermissionGrantIdValue")

read, err := client.RestoreMeJoinedTeamByIdPermissionGrantById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPermissionGrantClient.UpdateMeJoinedTeamByIdPermissionGrantById`

```go
ctx := context.TODO()
id := mejoinedteampermissiongrant.NewMeJoinedTeamPermissionGrantID("teamIdValue", "resourceSpecificPermissionGrantIdValue")

payload := mejoinedteampermissiongrant.ResourceSpecificPermissionGrant{
	// ...
}


read, err := client.UpdateMeJoinedTeamByIdPermissionGrantById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeJoinedTeamPermissionGrantClient.ValidateMeJoinedTeamByIdPermissionGrantsProperty`

```go
ctx := context.TODO()
id := mejoinedteampermissiongrant.NewMeJoinedTeamID("teamIdValue")

payload := mejoinedteampermissiongrant.ValidateMeJoinedTeamByIdPermissionGrantsPropertyRequest{
	// ...
}


read, err := client.ValidateMeJoinedTeamByIdPermissionGrantsProperty(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
