
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-03-01/diagnostics` Documentation

The `diagnostics` SDK allows for interaction with Azure Resource Manager `containerapps` (API Version `2024-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-03-01/diagnostics"
```


### Client Initialization

```go
client := diagnostics.NewDiagnosticsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DiagnosticsClient.ContainerAppsDiagnosticsGetDetector`

```go
ctx := context.TODO()
id := diagnostics.NewContainerAppDetectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppValue", "detectorValue")

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
id := diagnostics.NewRevisionsApiRevisionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppValue", "revisionValue")

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
id := diagnostics.NewContainerAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppValue")

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
id := diagnostics.NewContainerAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppValue")

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
id := diagnostics.NewContainerAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppValue")

// alternatively `client.ContainerAppsDiagnosticsListRevisions(ctx, id, diagnostics.DefaultContainerAppsDiagnosticsListRevisionsOperationOptions())` can be used to do batched pagination
items, err := client.ContainerAppsDiagnosticsListRevisionsComplete(ctx, id, diagnostics.DefaultContainerAppsDiagnosticsListRevisionsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DiagnosticsClient.JobsGetDetector`

```go
ctx := context.TODO()
id := diagnostics.NewDetectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "jobValue", "detectorValue")

read, err := client.JobsGetDetector(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticsClient.JobsListDetectors`

```go
ctx := context.TODO()
id := diagnostics.NewJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "jobValue")

// alternatively `client.JobsListDetectors(ctx, id)` can be used to do batched pagination
items, err := client.JobsListDetectorsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DiagnosticsClient.JobsProxyGet`

```go
ctx := context.TODO()
id := diagnostics.NewDetectorPropertyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "jobValue", "detectorPropertyValue")

read, err := client.JobsProxyGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiagnosticsClient.ManagedEnvironmentDiagnosticsGetDetector`

```go
ctx := context.TODO()
id := diagnostics.NewManagedEnvironmentDetectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedEnvironmentValue", "detectorValue")

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
id := diagnostics.NewManagedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedEnvironmentValue")

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
id := diagnostics.NewManagedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedEnvironmentValue")

read, err := client.ManagedEnvironmentsDiagnosticsGetRoot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
