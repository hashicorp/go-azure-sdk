
## `github.com/hashicorp/go-azure-sdk/resource-manager/applications/beta/applicationsynchronization` Documentation

The `applicationsynchronization` SDK allows for interaction with the Azure Resource Manager Service `applications` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/applications/beta/applicationsynchronization"
```


### Client Initialization

```go
client := applicationsynchronization.NewApplicationSynchronizationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApplicationSynchronizationClient.AcquireApplicationByIdSynchronizationAccessToken`

```go
ctx := context.TODO()
id := applicationsynchronization.NewApplicationID("applicationIdValue")

payload := applicationsynchronization.AcquireApplicationByIdSynchronizationAccessTokenRequest{
	// ...
}


read, err := client.AcquireApplicationByIdSynchronizationAccessToken(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationSynchronizationClient.CreateUpdateApplicationByIdSynchronization`

```go
ctx := context.TODO()
id := applicationsynchronization.NewApplicationID("applicationIdValue")

payload := applicationsynchronization.Synchronization{
	// ...
}


read, err := client.CreateUpdateApplicationByIdSynchronization(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationSynchronizationClient.DeleteApplicationByIdSynchronization`

```go
ctx := context.TODO()
id := applicationsynchronization.NewApplicationID("applicationIdValue")

read, err := client.DeleteApplicationByIdSynchronization(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationSynchronizationClient.GetApplicationByIdSynchronization`

```go
ctx := context.TODO()
id := applicationsynchronization.NewApplicationID("applicationIdValue")

read, err := client.GetApplicationByIdSynchronization(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
