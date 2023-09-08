
## `github.com/hashicorp/go-azure-sdk/resource-manager/applications/v1.0/applicationsynchronizationjob` Documentation

The `applicationsynchronizationjob` SDK allows for interaction with the Azure Resource Manager Service `applications` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/applications/v1.0/applicationsynchronizationjob"
```


### Client Initialization

```go
client := applicationsynchronizationjob.NewApplicationSynchronizationJobClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApplicationSynchronizationJobClient.CreateApplicationByIdSynchronizationJob`

```go
ctx := context.TODO()
id := applicationsynchronizationjob.NewApplicationID("applicationIdValue")

payload := applicationsynchronizationjob.SynchronizationJob{
	// ...
}


read, err := client.CreateApplicationByIdSynchronizationJob(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationSynchronizationJobClient.DeleteApplicationByIdSynchronizationJobById`

```go
ctx := context.TODO()
id := applicationsynchronizationjob.NewApplicationSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

read, err := client.DeleteApplicationByIdSynchronizationJobById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationSynchronizationJobClient.GetApplicationByIdSynchronizationJobById`

```go
ctx := context.TODO()
id := applicationsynchronizationjob.NewApplicationSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

read, err := client.GetApplicationByIdSynchronizationJobById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationSynchronizationJobClient.GetApplicationByIdSynchronizationJobCount`

```go
ctx := context.TODO()
id := applicationsynchronizationjob.NewApplicationID("applicationIdValue")

read, err := client.GetApplicationByIdSynchronizationJobCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationSynchronizationJobClient.ListApplicationByIdSynchronizationJobs`

```go
ctx := context.TODO()
id := applicationsynchronizationjob.NewApplicationID("applicationIdValue")

// alternatively `client.ListApplicationByIdSynchronizationJobs(ctx, id)` can be used to do batched pagination
items, err := client.ListApplicationByIdSynchronizationJobsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ApplicationSynchronizationJobClient.PauseApplicationByIdSynchronizationJobById`

```go
ctx := context.TODO()
id := applicationsynchronizationjob.NewApplicationSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

read, err := client.PauseApplicationByIdSynchronizationJobById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationSynchronizationJobClient.ProvisionApplicationByIdSynchronizationJobByIdOnDemand`

```go
ctx := context.TODO()
id := applicationsynchronizationjob.NewApplicationSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

payload := applicationsynchronizationjob.ProvisionApplicationByIdSynchronizationJobByIdOnDemandRequest{
	// ...
}


read, err := client.ProvisionApplicationByIdSynchronizationJobByIdOnDemand(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationSynchronizationJobClient.RestartApplicationByIdSynchronizationJobById`

```go
ctx := context.TODO()
id := applicationsynchronizationjob.NewApplicationSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

payload := applicationsynchronizationjob.RestartApplicationByIdSynchronizationJobByIdRequest{
	// ...
}


read, err := client.RestartApplicationByIdSynchronizationJobById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationSynchronizationJobClient.StartApplicationByIdSynchronizationJobById`

```go
ctx := context.TODO()
id := applicationsynchronizationjob.NewApplicationSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

read, err := client.StartApplicationByIdSynchronizationJobById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationSynchronizationJobClient.UpdateApplicationByIdSynchronizationJobById`

```go
ctx := context.TODO()
id := applicationsynchronizationjob.NewApplicationSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

payload := applicationsynchronizationjob.SynchronizationJob{
	// ...
}


read, err := client.UpdateApplicationByIdSynchronizationJobById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationSynchronizationJobClient.ValidateApplicationByIdSynchronizationJobByIdCredential`

```go
ctx := context.TODO()
id := applicationsynchronizationjob.NewApplicationSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

payload := applicationsynchronizationjob.ValidateApplicationByIdSynchronizationJobByIdCredentialRequest{
	// ...
}


read, err := client.ValidateApplicationByIdSynchronizationJobByIdCredential(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationSynchronizationJobClient.ValidateApplicationByIdSynchronizationJobsCredential`

```go
ctx := context.TODO()
id := applicationsynchronizationjob.NewApplicationID("applicationIdValue")

payload := applicationsynchronizationjob.ValidateApplicationByIdSynchronizationJobsCredentialRequest{
	// ...
}


read, err := client.ValidateApplicationByIdSynchronizationJobsCredential(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
