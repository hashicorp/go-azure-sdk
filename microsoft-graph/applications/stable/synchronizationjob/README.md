
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/synchronizationjob` Documentation

The `synchronizationjob` SDK allows for interaction with the Azure Resource Manager Service `applications` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/synchronizationjob"
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

read, err := client.DeleteSynchronizationJob(ctx, id, synchronizationjob.DefaultDeleteSynchronizationJobOperationOptions())
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

read, err := client.GetSynchronizationJob(ctx, id, synchronizationjob.DefaultGetSynchronizationJobOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.GetSynchronizationJobsCount`

```go
ctx := context.TODO()
id := synchronizationjob.NewApplicationID("applicationIdValue")

read, err := client.GetSynchronizationJobsCount(ctx, id, synchronizationjob.DefaultGetSynchronizationJobsCountOperationOptions())
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

// alternatively `client.ListSynchronizationJobs(ctx, id, synchronizationjob.DefaultListSynchronizationJobsOperationOptions())` can be used to do batched pagination
items, err := client.ListSynchronizationJobsComplete(ctx, id, synchronizationjob.DefaultListSynchronizationJobsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SynchronizationJobClient.PauseSynchronizationJob`

```go
ctx := context.TODO()
id := synchronizationjob.NewApplicationIdSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

read, err := client.PauseSynchronizationJob(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.ProvisionSynchronizationJobOnDemand`

```go
ctx := context.TODO()
id := synchronizationjob.NewApplicationIdSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

payload := synchronizationjob.ProvisionSynchronizationJobOnDemandRequest{
	// ...
}


read, err := client.ProvisionSynchronizationJobOnDemand(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.RestartSynchronizationJob`

```go
ctx := context.TODO()
id := synchronizationjob.NewApplicationIdSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

payload := synchronizationjob.RestartSynchronizationJobRequest{
	// ...
}


read, err := client.RestartSynchronizationJob(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.StartSynchronizationJob`

```go
ctx := context.TODO()
id := synchronizationjob.NewApplicationIdSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

read, err := client.StartSynchronizationJob(ctx, id)
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


### Example Usage: `SynchronizationJobClient.ValidateSynchronizationJobCredential`

```go
ctx := context.TODO()
id := synchronizationjob.NewApplicationIdSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

payload := synchronizationjob.ValidateSynchronizationJobCredentialRequest{
	// ...
}


read, err := client.ValidateSynchronizationJobCredential(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.ValidateSynchronizationJobsCredential`

```go
ctx := context.TODO()
id := synchronizationjob.NewApplicationID("applicationIdValue")

payload := synchronizationjob.ValidateSynchronizationJobsCredentialRequest{
	// ...
}


read, err := client.ValidateSynchronizationJobsCredential(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
