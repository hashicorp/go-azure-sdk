
## `github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/iotconnectors` Documentation

The `iotconnectors` SDK allows for interaction with the Azure Resource Manager Service `healthcareapis` (API Version `2021-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/healthcareapis/2021-06-01-preview/iotconnectors"
```


### Client Initialization

```go
client := iotconnectors.NewIotConnectorsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `IotConnectorsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := iotconnectors.NewIotConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "iotConnectorValue")

payload := iotconnectors.IotConnector{
	// ...
}

future, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `IotConnectorsClient.Delete`

```go
ctx := context.TODO()
id := iotconnectors.NewIotConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "iotConnectorValue")
future, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `IotConnectorsClient.FhirDestinationsListByIotConnector`

```go
ctx := context.TODO()
id := iotconnectors.NewIotConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "iotConnectorValue")
// alternatively `client.FhirDestinationsListByIotConnector(ctx, id)` can be used to do batched pagination
items, err := client.FhirDestinationsListByIotConnectorComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotConnectorsClient.Get`

```go
ctx := context.TODO()
id := iotconnectors.NewIotConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "iotConnectorValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotConnectorsClient.IotConnectorFhirDestinationCreateOrUpdate`

```go
ctx := context.TODO()
id := iotconnectors.NewFhirDestinationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "iotConnectorValue", "fhirDestinationValue")

payload := iotconnectors.IotFhirDestination{
	// ...
}

future, err := client.IotConnectorFhirDestinationCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `IotConnectorsClient.IotConnectorFhirDestinationDelete`

```go
ctx := context.TODO()
id := iotconnectors.NewFhirDestinationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "iotConnectorValue", "fhirDestinationValue")
future, err := client.IotConnectorFhirDestinationDelete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `IotConnectorsClient.IotConnectorFhirDestinationGet`

```go
ctx := context.TODO()
id := iotconnectors.NewFhirDestinationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "iotConnectorValue", "fhirDestinationValue")
read, err := client.IotConnectorFhirDestinationGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotConnectorsClient.ListByWorkspace`

```go
ctx := context.TODO()
id := iotconnectors.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")
// alternatively `client.ListByWorkspace(ctx, id)` can be used to do batched pagination
items, err := client.ListByWorkspaceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotConnectorsClient.Update`

```go
ctx := context.TODO()
id := iotconnectors.NewIotConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "iotConnectorValue")

payload := iotconnectors.IotConnectorPatchResource{
	// ...
}

future, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```
