
## `github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/standardoperation` Documentation

The `standardoperation` SDK allows for interaction with the Azure Resource Manager Service `datamigration` (API Version `2021-06-30`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/standardoperation"
```


### Client Initialization

```go
client := standardoperation.NewStandardOperationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `StandardOperationClient.FilesCreateOrUpdate`

```go
ctx := context.TODO()
id := standardoperation.NewFileID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue", "fileValue")

payload := standardoperation.ProjectFile{
	// ...
}


read, err := client.FilesCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandardOperationClient.FilesDelete`

```go
ctx := context.TODO()
id := standardoperation.NewFileID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue", "fileValue")

read, err := client.FilesDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandardOperationClient.FilesGet`

```go
ctx := context.TODO()
id := standardoperation.NewFileID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue", "fileValue")

read, err := client.FilesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandardOperationClient.FilesList`

```go
ctx := context.TODO()
id := standardoperation.NewProjectID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue")

// alternatively `client.FilesList(ctx, id)` can be used to do batched pagination
items, err := client.FilesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StandardOperationClient.FilesRead`

```go
ctx := context.TODO()
id := standardoperation.NewFileID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue", "fileValue")

read, err := client.FilesRead(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandardOperationClient.FilesReadWrite`

```go
ctx := context.TODO()
id := standardoperation.NewFileID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue", "fileValue")

read, err := client.FilesReadWrite(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandardOperationClient.FilesUpdate`

```go
ctx := context.TODO()
id := standardoperation.NewFileID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue", "fileValue")

payload := standardoperation.ProjectFile{
	// ...
}


read, err := client.FilesUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandardOperationClient.ProjectsCreateOrUpdate`

```go
ctx := context.TODO()
id := standardoperation.NewProjectID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue")

payload := standardoperation.Project{
	// ...
}


read, err := client.ProjectsCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandardOperationClient.ProjectsDelete`

```go
ctx := context.TODO()
id := standardoperation.NewProjectID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue")

read, err := client.ProjectsDelete(ctx, id, standardoperation.DefaultProjectsDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandardOperationClient.ProjectsGet`

```go
ctx := context.TODO()
id := standardoperation.NewProjectID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue")

read, err := client.ProjectsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandardOperationClient.ProjectsList`

```go
ctx := context.TODO()
id := standardoperation.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue")

// alternatively `client.ProjectsList(ctx, id)` can be used to do batched pagination
items, err := client.ProjectsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StandardOperationClient.ProjectsUpdate`

```go
ctx := context.TODO()
id := standardoperation.NewProjectID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue")

payload := standardoperation.Project{
	// ...
}


read, err := client.ProjectsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandardOperationClient.ResourceSkusListSkus`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ResourceSkusListSkus(ctx, id)` can be used to do batched pagination
items, err := client.ResourceSkusListSkusComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StandardOperationClient.ServiceTasksCreateOrUpdate`

```go
ctx := context.TODO()
id := standardoperation.NewServiceTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "serviceTaskValue")

payload := standardoperation.ProjectTask{
	// ...
}


read, err := client.ServiceTasksCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandardOperationClient.ServiceTasksDelete`

```go
ctx := context.TODO()
id := standardoperation.NewServiceTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "serviceTaskValue")

read, err := client.ServiceTasksDelete(ctx, id, standardoperation.DefaultServiceTasksDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandardOperationClient.ServiceTasksGet`

```go
ctx := context.TODO()
id := standardoperation.NewServiceTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "serviceTaskValue")

read, err := client.ServiceTasksGet(ctx, id, standardoperation.DefaultServiceTasksGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandardOperationClient.ServiceTasksList`

```go
ctx := context.TODO()
id := standardoperation.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue")

// alternatively `client.ServiceTasksList(ctx, id, standardoperation.DefaultServiceTasksListOperationOptions())` can be used to do batched pagination
items, err := client.ServiceTasksListComplete(ctx, id, standardoperation.DefaultServiceTasksListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StandardOperationClient.ServiceTasksUpdate`

```go
ctx := context.TODO()
id := standardoperation.NewServiceTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "serviceTaskValue")

payload := standardoperation.ProjectTask{
	// ...
}


read, err := client.ServiceTasksUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandardOperationClient.ServicesCheckNameAvailability`

```go
ctx := context.TODO()
id := standardoperation.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := standardoperation.NameAvailabilityRequest{
	// ...
}


read, err := client.ServicesCheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandardOperationClient.ServicesCreateOrUpdate`

```go
ctx := context.TODO()
id := standardoperation.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue")

payload := standardoperation.DataMigrationService{
	// ...
}


if err := client.ServicesCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `StandardOperationClient.ServicesDelete`

```go
ctx := context.TODO()
id := standardoperation.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue")

if err := client.ServicesDeleteThenPoll(ctx, id, standardoperation.DefaultServicesDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `StandardOperationClient.ServicesGet`

```go
ctx := context.TODO()
id := standardoperation.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue")

read, err := client.ServicesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandardOperationClient.ServicesList`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ServicesList(ctx, id)` can be used to do batched pagination
items, err := client.ServicesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StandardOperationClient.ServicesListByResourceGroup`

```go
ctx := context.TODO()
id := standardoperation.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "resourceGroupValue")

// alternatively `client.ServicesListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ServicesListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StandardOperationClient.ServicesListSkus`

```go
ctx := context.TODO()
id := standardoperation.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue")

// alternatively `client.ServicesListSkus(ctx, id)` can be used to do batched pagination
items, err := client.ServicesListSkusComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StandardOperationClient.ServicesUpdate`

```go
ctx := context.TODO()
id := standardoperation.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue")

payload := standardoperation.DataMigrationService{
	// ...
}


if err := client.ServicesUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `StandardOperationClient.TasksCreateOrUpdate`

```go
ctx := context.TODO()
id := standardoperation.NewTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue", "taskValue")

payload := standardoperation.ProjectTask{
	// ...
}


read, err := client.TasksCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandardOperationClient.TasksDelete`

```go
ctx := context.TODO()
id := standardoperation.NewTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue", "taskValue")

read, err := client.TasksDelete(ctx, id, standardoperation.DefaultTasksDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandardOperationClient.TasksGet`

```go
ctx := context.TODO()
id := standardoperation.NewTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue", "taskValue")

read, err := client.TasksGet(ctx, id, standardoperation.DefaultTasksGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandardOperationClient.TasksList`

```go
ctx := context.TODO()
id := standardoperation.NewProjectID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue")

// alternatively `client.TasksList(ctx, id, standardoperation.DefaultTasksListOperationOptions())` can be used to do batched pagination
items, err := client.TasksListComplete(ctx, id, standardoperation.DefaultTasksListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `StandardOperationClient.TasksUpdate`

```go
ctx := context.TODO()
id := standardoperation.NewTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue", "taskValue")

payload := standardoperation.ProjectTask{
	// ...
}


read, err := client.TasksUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StandardOperationClient.UsagesList`

```go
ctx := context.TODO()
id := standardoperation.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

// alternatively `client.UsagesList(ctx, id)` can be used to do batched pagination
items, err := client.UsagesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
