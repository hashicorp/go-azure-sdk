
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mepermissiongrant` Documentation

The `mepermissiongrant` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mepermissiongrant"
```


### Client Initialization

```go
client := mepermissiongrant.NewMePermissionGrantClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MePermissionGrantClient.CheckMePermissionGrantByIdMemberGroup`

```go
ctx := context.TODO()
id := mepermissiongrant.NewMePermissionGrantID("resourceSpecificPermissionGrantIdValue")

payload := mepermissiongrant.CheckMePermissionGrantByIdMemberGroupRequest{
	// ...
}


read, err := client.CheckMePermissionGrantByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePermissionGrantClient.CheckMePermissionGrantByIdMemberObject`

```go
ctx := context.TODO()
id := mepermissiongrant.NewMePermissionGrantID("resourceSpecificPermissionGrantIdValue")

payload := mepermissiongrant.CheckMePermissionGrantByIdMemberObjectRequest{
	// ...
}


read, err := client.CheckMePermissionGrantByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePermissionGrantClient.CreateMePermissionGrant`

```go
ctx := context.TODO()

payload := mepermissiongrant.ResourceSpecificPermissionGrant{
	// ...
}


read, err := client.CreateMePermissionGrant(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePermissionGrantClient.DeleteMePermissionGrantById`

```go
ctx := context.TODO()
id := mepermissiongrant.NewMePermissionGrantID("resourceSpecificPermissionGrantIdValue")

read, err := client.DeleteMePermissionGrantById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePermissionGrantClient.GetMePermissionGrantById`

```go
ctx := context.TODO()
id := mepermissiongrant.NewMePermissionGrantID("resourceSpecificPermissionGrantIdValue")

read, err := client.GetMePermissionGrantById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePermissionGrantClient.GetMePermissionGrantByIdMemberGroup`

```go
ctx := context.TODO()
id := mepermissiongrant.NewMePermissionGrantID("resourceSpecificPermissionGrantIdValue")

payload := mepermissiongrant.GetMePermissionGrantByIdMemberGroupRequest{
	// ...
}


read, err := client.GetMePermissionGrantByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePermissionGrantClient.GetMePermissionGrantByIdMemberObject`

```go
ctx := context.TODO()
id := mepermissiongrant.NewMePermissionGrantID("resourceSpecificPermissionGrantIdValue")

payload := mepermissiongrant.GetMePermissionGrantByIdMemberObjectRequest{
	// ...
}


read, err := client.GetMePermissionGrantByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePermissionGrantClient.GetMePermissionGrantCount`

```go
ctx := context.TODO()


read, err := client.GetMePermissionGrantCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePermissionGrantClient.GetMePermissionGrantsByIds`

```go
ctx := context.TODO()


// alternatively `client.GetMePermissionGrantsByIds(ctx)` can be used to do batched pagination
items, err := client.GetMePermissionGrantsByIdsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MePermissionGrantClient.GetMePermissionGrantsUserOwnedObject`

```go
ctx := context.TODO()

payload := mepermissiongrant.GetMePermissionGrantsUserOwnedObjectRequest{
	// ...
}


read, err := client.GetMePermissionGrantsUserOwnedObject(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePermissionGrantClient.ListMePermissionGrants`

```go
ctx := context.TODO()


// alternatively `client.ListMePermissionGrants(ctx)` can be used to do batched pagination
items, err := client.ListMePermissionGrantsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MePermissionGrantClient.RestoreMePermissionGrantById`

```go
ctx := context.TODO()
id := mepermissiongrant.NewMePermissionGrantID("resourceSpecificPermissionGrantIdValue")

read, err := client.RestoreMePermissionGrantById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePermissionGrantClient.UpdateMePermissionGrantById`

```go
ctx := context.TODO()
id := mepermissiongrant.NewMePermissionGrantID("resourceSpecificPermissionGrantIdValue")

payload := mepermissiongrant.ResourceSpecificPermissionGrant{
	// ...
}


read, err := client.UpdateMePermissionGrantById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MePermissionGrantClient.ValidateMePermissionGrantsProperty`

```go
ctx := context.TODO()

payload := mepermissiongrant.ValidateMePermissionGrantsPropertyRequest{
	// ...
}


read, err := client.ValidateMePermissionGrantsProperty(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
