
## `github.com/hashicorp/go-azure-sdk/resource-manager/servicelinker/2024-04-01/connector` Documentation

The `connector` SDK allows for interaction with Azure Resource Manager `servicelinker` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/servicelinker/2024-04-01/connector"
```


### Client Initialization

```go
client := connector.NewConnectorClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConnectorClient.CreateDryrun`

```go
ctx := context.TODO()
id := connector.NewDryrunID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationName", "dryrunName")

payload := connector.DryrunResource{
	// ...
}


if err := client.CreateDryrunThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ConnectorClient.Delete`

```go
ctx := context.TODO()
id := connector.NewConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationName", "connectorName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ConnectorClient.DeleteDryrun`

```go
ctx := context.TODO()
id := connector.NewDryrunID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationName", "dryrunName")

read, err := client.DeleteDryrun(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConnectorClient.GenerateConfigurations`

```go
ctx := context.TODO()
id := connector.NewConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationName", "connectorName")

payload := connector.ConfigurationInfo{
	// ...
}


read, err := client.GenerateConfigurations(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConnectorClient.GetDryrun`

```go
ctx := context.TODO()
id := connector.NewDryrunID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationName", "dryrunName")

read, err := client.GetDryrun(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConnectorClient.ListDryrun`

```go
ctx := context.TODO()
id := connector.NewLocationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationName")

// alternatively `client.ListDryrun(ctx, id)` can be used to do batched pagination
items, err := client.ListDryrunComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ConnectorClient.Update`

```go
ctx := context.TODO()
id := connector.NewConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationName", "connectorName")

payload := connector.LinkerPatch{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ConnectorClient.UpdateDryrun`

```go
ctx := context.TODO()
id := connector.NewDryrunID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationName", "dryrunName")

payload := connector.DryrunPatch{
	// ...
}


if err := client.UpdateDryrunThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ConnectorClient.Validate`

```go
ctx := context.TODO()
id := connector.NewConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationName", "connectorName")

if err := client.ValidateThenPoll(ctx, id); err != nil {
	// handle the error
}
```
