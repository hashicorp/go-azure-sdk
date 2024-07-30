
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/application` Documentation

The `application` SDK allows for interaction with the Azure Resource Manager Service `applications` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/application"
```


### Client Initialization

```go
client := application.NewApplicationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApplicationClient.AddApplicationKey`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.AddApplicationKeyRequest{
	// ...
}


read, err := client.AddApplicationKey(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.AddApplicationPassword`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.AddApplicationPasswordRequest{
	// ...
}


read, err := client.AddApplicationPassword(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.CheckApplicationMemberGroup`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.CheckApplicationMemberGroupRequest{
	// ...
}


read, err := client.CheckApplicationMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.CheckApplicationMemberObject`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.CheckApplicationMemberObjectRequest{
	// ...
}


read, err := client.CheckApplicationMemberObject(ctx, id, payload)
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


### Example Usage: `ApplicationClient.DeleteApplication`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

read, err := client.DeleteApplication(ctx, id)
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

read, err := client.GetApplication(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.GetApplicationMemberGroup`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.GetApplicationMemberGroupRequest{
	// ...
}


read, err := client.GetApplicationMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.GetApplicationMemberObject`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.GetApplicationMemberObjectRequest{
	// ...
}


read, err := client.GetApplicationMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.GetApplicationsAvailableExtensionProperties`

```go
ctx := context.TODO()

payload := application.GetApplicationsAvailableExtensionPropertiesRequest{
	// ...
}


// alternatively `client.GetApplicationsAvailableExtensionProperties(ctx, payload)` can be used to do batched pagination
items, err := client.GetApplicationsAvailableExtensionPropertiesComplete(ctx, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ApplicationClient.GetApplicationsByIds`

```go
ctx := context.TODO()

payload := application.GetApplicationsByIdsRequest{
	// ...
}


// alternatively `client.GetApplicationsByIds(ctx, payload)` can be used to do batched pagination
items, err := client.GetApplicationsByIdsComplete(ctx, payload)
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


read, err := client.GetCount(ctx)
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


### Example Usage: `ApplicationClient.RemoveApplicationKey`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.RemoveApplicationKeyRequest{
	// ...
}


read, err := client.RemoveApplicationKey(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.RemoveApplicationPassword`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.RemoveApplicationPasswordRequest{
	// ...
}


read, err := client.RemoveApplicationPassword(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.RestoreApplication`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

read, err := client.RestoreApplication(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.SetApplicationVerifiedPublisher`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

payload := application.SetApplicationVerifiedPublisherRequest{
	// ...
}


read, err := client.SetApplicationVerifiedPublisher(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.UnsetApplicationVerifiedPublisher`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationIdValue")

read, err := client.UnsetApplicationVerifiedPublisher(ctx, id)
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
