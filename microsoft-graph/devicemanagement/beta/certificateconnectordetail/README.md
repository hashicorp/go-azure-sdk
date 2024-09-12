
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/certificateconnectordetail` Documentation

The `certificateconnectordetail` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/certificateconnectordetail"
```


### Client Initialization

```go
client := certificateconnectordetail.NewCertificateConnectorDetailClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CertificateConnectorDetailClient.CreateCertificateConnectorDetail`

```go
ctx := context.TODO()

payload := certificateconnectordetail.CertificateConnectorDetails{
	// ...
}


read, err := client.CreateCertificateConnectorDetail(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CertificateConnectorDetailClient.DeleteCertificateConnectorDetail`

```go
ctx := context.TODO()
id := certificateconnectordetail.NewDeviceManagementCertificateConnectorDetailID("certificateConnectorDetailsIdValue")

read, err := client.DeleteCertificateConnectorDetail(ctx, id, certificateconnectordetail.DefaultDeleteCertificateConnectorDetailOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CertificateConnectorDetailClient.GetCertificateConnectorDetail`

```go
ctx := context.TODO()
id := certificateconnectordetail.NewDeviceManagementCertificateConnectorDetailID("certificateConnectorDetailsIdValue")

read, err := client.GetCertificateConnectorDetail(ctx, id, certificateconnectordetail.DefaultGetCertificateConnectorDetailOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CertificateConnectorDetailClient.GetCertificateConnectorDetailHealthMetricTimeSeries`

```go
ctx := context.TODO()
id := certificateconnectordetail.NewDeviceManagementCertificateConnectorDetailID("certificateConnectorDetailsIdValue")

payload := certificateconnectordetail.GetCertificateConnectorDetailHealthMetricTimeSeriesRequest{
	// ...
}


// alternatively `client.GetCertificateConnectorDetailHealthMetricTimeSeries(ctx, id, payload, certificateconnectordetail.DefaultGetCertificateConnectorDetailHealthMetricTimeSeriesOperationOptions())` can be used to do batched pagination
items, err := client.GetCertificateConnectorDetailHealthMetricTimeSeriesComplete(ctx, id, payload, certificateconnectordetail.DefaultGetCertificateConnectorDetailHealthMetricTimeSeriesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CertificateConnectorDetailClient.GetCertificateConnectorDetailHealthMetrics`

```go
ctx := context.TODO()
id := certificateconnectordetail.NewDeviceManagementCertificateConnectorDetailID("certificateConnectorDetailsIdValue")

payload := certificateconnectordetail.GetCertificateConnectorDetailHealthMetricsRequest{
	// ...
}


// alternatively `client.GetCertificateConnectorDetailHealthMetrics(ctx, id, payload, certificateconnectordetail.DefaultGetCertificateConnectorDetailHealthMetricsOperationOptions())` can be used to do batched pagination
items, err := client.GetCertificateConnectorDetailHealthMetricsComplete(ctx, id, payload, certificateconnectordetail.DefaultGetCertificateConnectorDetailHealthMetricsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CertificateConnectorDetailClient.GetCertificateConnectorDetailsCount`

```go
ctx := context.TODO()


read, err := client.GetCertificateConnectorDetailsCount(ctx, certificateconnectordetail.DefaultGetCertificateConnectorDetailsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CertificateConnectorDetailClient.ListCertificateConnectorDetails`

```go
ctx := context.TODO()


// alternatively `client.ListCertificateConnectorDetails(ctx, certificateconnectordetail.DefaultListCertificateConnectorDetailsOperationOptions())` can be used to do batched pagination
items, err := client.ListCertificateConnectorDetailsComplete(ctx, certificateconnectordetail.DefaultListCertificateConnectorDetailsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CertificateConnectorDetailClient.UpdateCertificateConnectorDetail`

```go
ctx := context.TODO()
id := certificateconnectordetail.NewDeviceManagementCertificateConnectorDetailID("certificateConnectorDetailsIdValue")

payload := certificateconnectordetail.CertificateConnectorDetails{
	// ...
}


read, err := client.UpdateCertificateConnectorDetail(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
