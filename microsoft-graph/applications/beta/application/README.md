
## `github.com/hashicorp/go-azure-sdk/resource-manager/applications/beta/application` Documentation

The `application` SDK allows for interaction with the Azure Resource Manager Service `applications` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/applications/beta/application"
```


### Client Initialization

```go
client := application.NewApplicationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApplicationClient.AddApplicationByIdKey`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.AddApplicationByIdKeyRequest{
	// ...
}


read, err := client.AddApplicationByIdKey(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.AddApplicationByIdPassword`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.AddApplicationByIdPasswordRequest{
	// ...
}


read, err := client.AddApplicationByIdPassword(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.CheckApplicationByIdMemberGroup`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.CheckApplicationByIdMemberGroupRequest{
	// ...
}


read, err := client.CheckApplicationByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.CheckApplicationByIdMemberObject`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.CheckApplicationByIdMemberObjectRequest{
	// ...
}


read, err := client.CheckApplicationByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
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


### Example Usage: `ApplicationClient.DeleteApplicationById`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

read, err := client.DeleteApplicationById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.GetApplicationById`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

read, err := client.GetApplicationById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.GetApplicationByIdMemberGroup`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.GetApplicationByIdMemberGroupRequest{
	// ...
}


read, err := client.GetApplicationByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.GetApplicationByIdMemberObject`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.GetApplicationByIdMemberObjectRequest{
	// ...
}


read, err := client.GetApplicationByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.GetApplicationCount`

```go
ctx := context.TODO()


read, err := client.GetApplicationCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.GetApplicationsByIds`

```go
ctx := context.TODO()


// alternatively `client.GetApplicationsByIds(ctx)` can be used to do batched pagination
items, err := client.GetApplicationsByIdsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ApplicationClient.GetApplicationsUserOwnedObject`

```go
ctx := context.TODO()

payload := application.GetApplicationsUserOwnedObjectRequest{
	// ...
}


read, err := client.GetApplicationsUserOwnedObject(ctx, payload)
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


// alternatively `client.ListApplications(ctx)` can be used to do batched pagination
items, err := client.ListApplicationsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ApplicationClient.RemoveApplicationByIdKey`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.RemoveApplicationByIdKeyRequest{
	// ...
}


read, err := client.RemoveApplicationByIdKey(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.RemoveApplicationByIdPassword`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.RemoveApplicationByIdPasswordRequest{
	// ...
}


read, err := client.RemoveApplicationByIdPassword(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.RestoreApplicationById`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

read, err := client.RestoreApplicationById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.SetApplicationByIdVerifiedPublisher`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.SetApplicationByIdVerifiedPublisherRequest{
	// ...
}


read, err := client.SetApplicationByIdVerifiedPublisher(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.UnsetApplicationByIdVerifiedPublisher`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

read, err := client.UnsetApplicationByIdVerifiedPublisher(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.UpdateApplicationById`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.Application{
	// ...
}


read, err := client.UpdateApplicationById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.ValidateApplicationsProperty`

```go
ctx := context.TODO()

payload := application.ValidateApplicationsPropertyRequest{
	// ...
}


read, err := client.ValidateApplicationsProperty(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
