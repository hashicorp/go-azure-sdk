
## `github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2018-04-19/put` Documentation

The `put` SDK allows for interaction with Azure Resource Manager `datamigration` (API Version `2018-04-19`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2018-04-19/put"
```


### Client Initialization

```go
client := put.NewPUTClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PUTClient.ProjectsCreateOrUpdate`

```go
ctx := context.TODO()
id := put.NewProjectID("12345678-1234-9876-4563-123456789012", "resourceGroupName", "serviceName", "projectName")

payload := put.Project{
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


### Example Usage: `PUTClient.ServicesCreateOrUpdate`

```go
ctx := context.TODO()
id := put.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupName", "serviceName")

payload := put.DataMigrationService{
	// ...
}


if err := client.ServicesCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `PUTClient.TasksCreateOrUpdate`

```go
ctx := context.TODO()
id := put.NewTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupName", "serviceName", "projectName", "taskName")

payload := put.ProjectTask{
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
