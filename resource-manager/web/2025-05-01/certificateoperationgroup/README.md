
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/certificateoperationgroup` Documentation

The `certificateoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/certificateoperationgroup"
```


### Client Initialization

```go
client := certificateoperationgroup.NewCertificateOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CertificateOperationGroupClient.SiteCertificatesCreateOrUpdateSlot`

```go
ctx := context.TODO()
id := certificateoperationgroup.NewSlotCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "certificateName")

payload := certificateoperationgroup.Certificate{
	// ...
}


read, err := client.SiteCertificatesCreateOrUpdateSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CertificateOperationGroupClient.SiteCertificatesDeleteSlot`

```go
ctx := context.TODO()
id := certificateoperationgroup.NewSlotCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "certificateName")

read, err := client.SiteCertificatesDeleteSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CertificateOperationGroupClient.SiteCertificatesGetSlot`

```go
ctx := context.TODO()
id := certificateoperationgroup.NewSlotCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "certificateName")

read, err := client.SiteCertificatesGetSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CertificateOperationGroupClient.SiteCertificatesListSlot`

```go
ctx := context.TODO()
id := certificateoperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.SiteCertificatesListSlot(ctx, id)` can be used to do batched pagination
items, err := client.SiteCertificatesListSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CertificateOperationGroupClient.SiteCertificatesUpdateSlot`

```go
ctx := context.TODO()
id := certificateoperationgroup.NewSlotCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "certificateName")

payload := certificateoperationgroup.CertificatePatchResource{
	// ...
}


read, err := client.SiteCertificatesUpdateSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
