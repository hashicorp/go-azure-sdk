
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/synchronizationjob` Documentation

The `synchronizationjob` SDK allows for interaction with the Azure Resource Manager Service `applications` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/synchronizationjob"
```


### Client Initialization

```go
client := synchronizationjob.NewSynchronizationJobClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SynchronizationJobClient.CreateSynchronizationJob`

```go
ctx := context.TODO()
id := synchronizationjob.NewApplicationID("applicationIdValue")

payload := synchronizationjob.SynchronizationJob{
	// ...
}


read, err := client.CreateSynchronizationJob(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.DeleteSynchronizationJob`

```go
ctx := context.TODO()
id := synchronizationjob.NewApplicationIdSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

read, err := client.DeleteSynchronizationJob(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.GetSynchronizationJob`

```go
ctx := context.TODO()
id := synchronizationjob.NewApplicationIdSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

read, err := client.GetSynchronizationJob(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.GetSynchronizationJobCount`

```go
ctx := context.TODO()
id := synchronizationjob.NewApplicationID("applicationIdValue")

read, err := client.GetSynchronizationJobCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.ListSynchronizationJobs`

```go
ctx := context.TODO()
id := synchronizationjob.NewApplicationID("applicationIdValue")

// alternatively `client.ListSynchronizationJobs(ctx, id)` can be used to do batched pagination
items, err := client.ListSynchronizationJobsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SynchronizationJobClient.PauseApplicationSynchronizationJob`

```go
ctx := context.TODO()
id := synchronizationjob.NewApplicationIdSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

read, err := client.PauseApplicationSynchronizationJob(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.ProvisionApplicationSynchronizationJobOnDemand`

```go
ctx := context.TODO()
id := synchronizationjob.NewApplicationIdSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

payload := synchronizationjob.ProvisionApplicationSynchronizationJobOnDemandRequest{
	// ...
}


read, err := client.ProvisionApplicationSynchronizationJobOnDemand(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.RestartApplicationSynchronizationJob`

```go
ctx := context.TODO()
id := synchronizationjob.NewApplicationIdSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

payload := synchronizationjob.RestartApplicationSynchronizationJobRequest{
	// ...
}


read, err := client.RestartApplicationSynchronizationJob(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.StartApplicationSynchronizationJob`

```go
ctx := context.TODO()
id := synchronizationjob.NewApplicationIdSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

read, err := client.StartApplicationSynchronizationJob(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.UpdateSynchronizationJob`

```go
ctx := context.TODO()
id := synchronizationjob.NewApplicationIdSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

payload := synchronizationjob.SynchronizationJob{
	// ...
}


read, err := client.UpdateSynchronizationJob(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.ValidateApplicationSynchronizationJobCredential`

```go
ctx := context.TODO()
id := synchronizationjob.NewApplicationIdSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

payload := synchronizationjob.ValidateApplicationSynchronizationJobCredentialRequest{
	// ...
}


read, err := client.ValidateApplicationSynchronizationJobCredential(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.ValidateApplicationSynchronizationJobsCredential`

```go
ctx := context.TODO()
id := synchronizationjob.NewApplicationID("applicationIdValue")

payload := synchronizationjob.ValidateApplicationSynchronizationJobsCredentialRequest{
	// ...
}


read, err := client.ValidateApplicationSynchronizationJobsCredential(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
