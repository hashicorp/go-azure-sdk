
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/diagnostics` Documentation

The `diagnostics` SDK allows for interaction with the Azure Resource Manager Service `containerapps` (API Version `2022-11-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/diagnostics"
```


### Client Initialization

```go
client := diagnostics.NewDiagnosticsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DiagnosticsClient.ContainerAppsDiagnosticsGetDetector`

```go
ctx := context.TODO()
id := diagnostics.NewDetectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppValue", "detectorValue")

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

read, err := client.ManagedEnvironmentDiagnosticsListDetectors(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
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
