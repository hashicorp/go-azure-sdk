
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/jobprivateendpoints` Documentation

The `jobprivateendpoints` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/jobprivateendpoints"
```


### Client Initialization

```go
client := jobprivateendpoints.NewJobPrivateEndpointsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JobPrivateEndpointsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := jobprivateendpoints.NewPrivateEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "jobAgentName", "privateEndpointName")

payload := jobprivateendpoints.JobPrivateEndpoint{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `JobPrivateEndpointsClient.Delete`

```go
ctx := context.TODO()
id := jobprivateendpoints.NewPrivateEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "jobAgentName", "privateEndpointName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `JobPrivateEndpointsClient.Get`

```go
ctx := context.TODO()
id := jobprivateendpoints.NewPrivateEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "jobAgentName", "privateEndpointName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `JobPrivateEndpointsClient.ListByAgent`

```go
ctx := context.TODO()
id := jobprivateendpoints.NewJobAgentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "jobAgentName")

// alternatively `client.ListByAgent(ctx, id)` can be used to do batched pagination
items, err := client.ListByAgentComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
