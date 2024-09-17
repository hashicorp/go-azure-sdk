
## `github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2018-04-19/standardoperation` Documentation

The `standardoperation` SDK allows for interaction with Azure Resource Manager `datamigration` (API Version `2018-04-19`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2018-04-19/standardoperation"
```


### Client Initialization

```go
client := standardoperation.NewStandardOperationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
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


### Example Usage: `StandardOperationClient.ProjectsListByResourceGroup`

```go
ctx := context.TODO()
id := standardoperation.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue")

// alternatively `client.ProjectsListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ProjectsListByResourceGroupComplete(ctx, id)
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
