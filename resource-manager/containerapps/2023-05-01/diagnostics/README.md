
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2023-05-01/diagnostics` Documentation

The `diagnostics` SDK allows for interaction with Azure Resource Manager `containerapps` (API Version `2023-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2023-05-01/diagnostics"
```


### Client Initialization

```go
client := diagnostics.NewDiagnosticsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DiagnosticsClient.ContainerAppsDiagnosticsGetDetector`

```go
ctx := context.TODO()
id := diagnostics.NewDetectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "detectorName")

read, err := client.ContainerAppsDiagnosticsGetDetector(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticsClient.ContainerAppsDiagnosticsGetRevision`

```go
ctx := context.TODO()
id := diagnostics.NewRevisionsApiRevisionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "revisionName")

read, err := client.ContainerAppsDiagnosticsGetRevision(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticsClient.ContainerAppsDiagnosticsGetRoot`

```go
ctx := context.TODO()
id := diagnostics.NewContainerAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName")

read, err := client.ContainerAppsDiagnosticsGetRoot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticsClient.ContainerAppsDiagnosticsListDetectors`

```go
ctx := context.TODO()
id := diagnostics.NewContainerAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName")

// alternatively `client.ContainerAppsDiagnosticsListDetectors(ctx, id)` can be used to do batched pagination
items, err := client.ContainerAppsDiagnosticsListDetectorsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DiagnosticsClient.ContainerAppsDiagnosticsListRevisions`

```go
ctx := context.TODO()
id := diagnostics.NewContainerAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName")

// alternatively `client.ContainerAppsDiagnosticsListRevisions(ctx, id, diagnostics.DefaultContainerAppsDiagnosticsListRevisionsOperationOptions())` can be used to do batched pagination
items, err := client.ContainerAppsDiagnosticsListRevisionsComplete(ctx, id, diagnostics.DefaultContainerAppsDiagnosticsListRevisionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DiagnosticsClient.ManagedEnvironmentDiagnosticsGetDetector`

```go
ctx := context.TODO()
id := diagnostics.NewManagedEnvironmentDetectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "environmentName", "detectorName")

read, err := client.ManagedEnvironmentDiagnosticsGetDetector(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticsClient.ManagedEnvironmentDiagnosticsListDetectors`

```go
ctx := context.TODO()
id := diagnostics.NewManagedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "environmentName")

// alternatively `client.ManagedEnvironmentDiagnosticsListDetectors(ctx, id)` can be used to do batched pagination
items, err := client.ManagedEnvironmentDiagnosticsListDetectorsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DiagnosticsClient.ManagedEnvironmentsDiagnosticsGetRoot`

```go
ctx := context.TODO()
id := diagnostics.NewManagedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "environmentName")

read, err := client.ManagedEnvironmentsDiagnosticsGetRoot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
