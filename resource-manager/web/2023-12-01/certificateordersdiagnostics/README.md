
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2023-12-01/certificateordersdiagnostics` Documentation

The `certificateordersdiagnostics` SDK allows for interaction with Azure Resource Manager `web` (API Version `2023-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2023-12-01/certificateordersdiagnostics"
```


### Client Initialization

```go
client := certificateordersdiagnostics.NewCertificateOrdersDiagnosticsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CertificateOrdersDiagnosticsClient.GetAppServiceCertificateOrderDetectorResponse`

```go
ctx := context.TODO()
id := certificateordersdiagnostics.NewDetectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "certificateOrderName", "detectorName")

read, err := client.GetAppServiceCertificateOrderDetectorResponse(ctx, id, certificateordersdiagnostics.DefaultGetAppServiceCertificateOrderDetectorResponseOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CertificateOrdersDiagnosticsClient.ListAppServiceCertificateOrderDetectorResponse`

```go
ctx := context.TODO()
id := certificateordersdiagnostics.NewCertificateOrderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "certificateOrderName")

// alternatively `client.ListAppServiceCertificateOrderDetectorResponse(ctx, id)` can be used to do batched pagination
items, err := client.ListAppServiceCertificateOrderDetectorResponseComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
