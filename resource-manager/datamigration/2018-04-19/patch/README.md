
## `github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2018-04-19/patch` Documentation

The `patch` SDK allows for interaction with Azure Resource Manager `datamigration` (API Version `2018-04-19`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2018-04-19/patch"
```


### Client Initialization

```go
client := patch.NewPATCHClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PATCHClient.ProjectsUpdate`

```go
ctx := context.TODO()
id := patch.NewProjectID("12345678-1234-9876-4563-123456789012", "resourceGroupName", "serviceName", "projectName")

payload := patch.Project{
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


### Example Usage: `PATCHClient.ServicesUpdate`

```go
ctx := context.TODO()
id := patch.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupName", "serviceName")

payload := patch.DataMigrationService{
	// ...
}


if err := client.ServicesUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `PATCHClient.TasksUpdate`

```go
ctx := context.TODO()
id := patch.NewTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupName", "serviceName", "projectName", "taskName")

payload := patch.ProjectTask{
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
