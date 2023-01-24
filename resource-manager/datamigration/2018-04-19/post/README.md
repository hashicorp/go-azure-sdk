
## `github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2018-04-19/post` Documentation

The `post` SDK allows for interaction with the Azure Resource Manager Service `datamigration` (API Version `2018-04-19`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2018-04-19/post"
```


### Client Initialization

```go
client := post.NewPOSTClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `POSTClient.ServicesCheckNameAvailability`

```go
ctx := context.TODO()
id := post.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := post.NameAvailabilityRequest{
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


### Example Usage: `POSTClient.ServicesCheckStatus`

```go
ctx := context.TODO()
id := post.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue")

read, err := client.ServicesCheckStatus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `POSTClient.ServicesNestedCheckNameAvailability`

```go
ctx := context.TODO()
id := post.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue")

payload := post.NameAvailabilityRequest{
	// ...
}


read, err := client.ServicesNestedCheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `POSTClient.ServicesStart`

```go
ctx := context.TODO()
id := post.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue")

if err := client.ServicesStartThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `POSTClient.ServicesStop`

```go
ctx := context.TODO()
id := post.NewServiceID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue")

if err := client.ServicesStopThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `POSTClient.TasksCancel`

```go
ctx := context.TODO()
id := post.NewTaskID("12345678-1234-9876-4563-123456789012", "resourceGroupValue", "serviceValue", "projectValue", "taskValue")

read, err := client.TasksCancel(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
