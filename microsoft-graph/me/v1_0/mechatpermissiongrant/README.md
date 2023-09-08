
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mechatpermissiongrant` Documentation

The `mechatpermissiongrant` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/mechatpermissiongrant"
```


### Client Initialization

```go
client := mechatpermissiongrant.NewMeChatPermissionGrantClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeChatPermissionGrantClient.CheckMeChatByIdPermissionGrantByIdMemberGroup`

```go
ctx := context.TODO()
id := mechatpermissiongrant.NewMeChatPermissionGrantID("chatIdValue", "resourceSpecificPermissionGrantIdValue")

payload := mechatpermissiongrant.CheckMeChatByIdPermissionGrantByIdMemberGroupRequest{
	// ...
}


read, err := client.CheckMeChatByIdPermissionGrantByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatPermissionGrantClient.CheckMeChatByIdPermissionGrantByIdMemberObject`

```go
ctx := context.TODO()
id := mechatpermissiongrant.NewMeChatPermissionGrantID("chatIdValue", "resourceSpecificPermissionGrantIdValue")

payload := mechatpermissiongrant.CheckMeChatByIdPermissionGrantByIdMemberObjectRequest{
	// ...
}


read, err := client.CheckMeChatByIdPermissionGrantByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatPermissionGrantClient.CreateMeChatByIdPermissionGrant`

```go
ctx := context.TODO()
id := mechatpermissiongrant.NewMeChatID("chatIdValue")

payload := mechatpermissiongrant.ResourceSpecificPermissionGrant{
	// ...
}


read, err := client.CreateMeChatByIdPermissionGrant(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatPermissionGrantClient.DeleteMeChatByIdPermissionGrantById`

```go
ctx := context.TODO()
id := mechatpermissiongrant.NewMeChatPermissionGrantID("chatIdValue", "resourceSpecificPermissionGrantIdValue")

read, err := client.DeleteMeChatByIdPermissionGrantById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatPermissionGrantClient.GetMeChatByIdPermissionGrantById`

```go
ctx := context.TODO()
id := mechatpermissiongrant.NewMeChatPermissionGrantID("chatIdValue", "resourceSpecificPermissionGrantIdValue")

read, err := client.GetMeChatByIdPermissionGrantById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatPermissionGrantClient.GetMeChatByIdPermissionGrantByIdMemberGroup`

```go
ctx := context.TODO()
id := mechatpermissiongrant.NewMeChatPermissionGrantID("chatIdValue", "resourceSpecificPermissionGrantIdValue")

payload := mechatpermissiongrant.GetMeChatByIdPermissionGrantByIdMemberGroupRequest{
	// ...
}


read, err := client.GetMeChatByIdPermissionGrantByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatPermissionGrantClient.GetMeChatByIdPermissionGrantByIdMemberObject`

```go
ctx := context.TODO()
id := mechatpermissiongrant.NewMeChatPermissionGrantID("chatIdValue", "resourceSpecificPermissionGrantIdValue")

payload := mechatpermissiongrant.GetMeChatByIdPermissionGrantByIdMemberObjectRequest{
	// ...
}


read, err := client.GetMeChatByIdPermissionGrantByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatPermissionGrantClient.GetMeChatByIdPermissionGrantCount`

```go
ctx := context.TODO()
id := mechatpermissiongrant.NewMeChatID("chatIdValue")

read, err := client.GetMeChatByIdPermissionGrantCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatPermissionGrantClient.GetMeChatByIdPermissionGrantsAvailableExtensionProperties`

```go
ctx := context.TODO()
id := mechatpermissiongrant.NewMeChatID("chatIdValue")

// alternatively `client.GetMeChatByIdPermissionGrantsAvailableExtensionProperties(ctx, id)` can be used to do batched pagination
items, err := client.GetMeChatByIdPermissionGrantsAvailableExtensionPropertiesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeChatPermissionGrantClient.GetMeChatByIdPermissionGrantsByIds`

```go
ctx := context.TODO()
id := mechatpermissiongrant.NewMeChatID("chatIdValue")

// alternatively `client.GetMeChatByIdPermissionGrantsByIds(ctx, id)` can be used to do batched pagination
items, err := client.GetMeChatByIdPermissionGrantsByIdsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeChatPermissionGrantClient.ListMeChatByIdPermissionGrants`

```go
ctx := context.TODO()
id := mechatpermissiongrant.NewMeChatID("chatIdValue")

// alternatively `client.ListMeChatByIdPermissionGrants(ctx, id)` can be used to do batched pagination
items, err := client.ListMeChatByIdPermissionGrantsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeChatPermissionGrantClient.RestoreMeChatByIdPermissionGrantById`

```go
ctx := context.TODO()
id := mechatpermissiongrant.NewMeChatPermissionGrantID("chatIdValue", "resourceSpecificPermissionGrantIdValue")

read, err := client.RestoreMeChatByIdPermissionGrantById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatPermissionGrantClient.UpdateMeChatByIdPermissionGrantById`

```go
ctx := context.TODO()
id := mechatpermissiongrant.NewMeChatPermissionGrantID("chatIdValue", "resourceSpecificPermissionGrantIdValue")

payload := mechatpermissiongrant.ResourceSpecificPermissionGrant{
	// ...
}


read, err := client.UpdateMeChatByIdPermissionGrantById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeChatPermissionGrantClient.ValidateMeChatByIdPermissionGrantsProperty`

```go
ctx := context.TODO()
id := mechatpermissiongrant.NewMeChatID("chatIdValue")

payload := mechatpermissiongrant.ValidateMeChatByIdPermissionGrantsPropertyRequest{
	// ...
}


read, err := client.ValidateMeChatByIdPermissionGrantsProperty(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
