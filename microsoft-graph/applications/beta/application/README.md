
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/application` Documentation

The `application` SDK allows for interaction with the Azure Resource Manager Service `applications` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/application"
```


### Client Initialization

```go
client := application.NewApplicationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApplicationClient.AddKey`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.AddKeyRequest{
	// ...
}


read, err := client.AddKey(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.AddPassword`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.AddPasswordRequest{
	// ...
}


read, err := client.AddPassword(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.CheckMemberGroups`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.CheckMemberGroupsRequest{
	// ...
}


// alternatively `client.CheckMemberGroups(ctx, id, payload, application.DefaultCheckMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberGroupsComplete(ctx, id, payload, application.DefaultCheckMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ApplicationClient.CheckMemberObjects`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.CheckMemberObjectsRequest{
	// ...
}


// alternatively `client.CheckMemberObjects(ctx, id, payload, application.DefaultCheckMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberObjectsComplete(ctx, id, payload, application.DefaultCheckMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ApplicationClient.CreateApplication`

```go
ctx := context.TODO()

payload := application.Application{
	// ...
}


read, err := client.CreateApplication(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.DeleteApplication`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

read, err := client.DeleteApplication(ctx, id, application.DefaultDeleteApplicationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.GetApplication`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

read, err := client.GetApplication(ctx, id, application.DefaultGetApplicationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.GetByIds`

```go
ctx := context.TODO()

payload := application.GetByIdsRequest{
	// ...
}


// alternatively `client.GetByIds(ctx, payload, application.DefaultGetByIdsOperationOptions())` can be used to do batched pagination
items, err := client.GetByIdsComplete(ctx, payload, application.DefaultGetByIdsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ApplicationClient.GetCount`

```go
ctx := context.TODO()


read, err := client.GetCount(ctx, application.DefaultGetCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.GetMemberGroups`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.GetMemberGroupsRequest{
	// ...
}


// alternatively `client.GetMemberGroups(ctx, id, payload, application.DefaultGetMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberGroupsComplete(ctx, id, payload, application.DefaultGetMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ApplicationClient.GetMemberObjects`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.GetMemberObjectsRequest{
	// ...
}


// alternatively `client.GetMemberObjects(ctx, id, payload, application.DefaultGetMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberObjectsComplete(ctx, id, payload, application.DefaultGetMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ApplicationClient.GetUserOwnedObject`

```go
ctx := context.TODO()

payload := application.GetUserOwnedObjectRequest{
	// ...
}


read, err := client.GetUserOwnedObject(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.ListApplications`

```go
ctx := context.TODO()


// alternatively `client.ListApplications(ctx, application.DefaultListApplicationsOperationOptions())` can be used to do batched pagination
items, err := client.ListApplicationsComplete(ctx, application.DefaultListApplicationsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ApplicationClient.RemoveKey`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.RemoveKeyRequest{
	// ...
}


read, err := client.RemoveKey(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.RemovePassword`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.RemovePasswordRequest{
	// ...
}


read, err := client.RemovePassword(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.Restore`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.RestoreRequest{
	// ...
}


read, err := client.Restore(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.SetVerifiedPublisher`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.SetVerifiedPublisherRequest{
	// ...
}


read, err := client.SetVerifiedPublisher(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.UnsetVerifiedPublisher`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

read, err := client.UnsetVerifiedPublisher(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.UpdateApplication`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.Application{
	// ...
}


read, err := client.UpdateApplication(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.ValidateProperty`

```go
ctx := context.TODO()

payload := application.ValidatePropertyRequest{
	// ...
}


read, err := client.ValidateProperty(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
