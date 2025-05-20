
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-11-01/sitecertificates` Documentation

The `sitecertificates` SDK allows for interaction with Azure Resource Manager `web` (API Version `2024-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-11-01/sitecertificates"
```


### Client Initialization

```go
client := sitecertificates.NewSiteCertificatesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteCertificatesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := sitecertificates.NewSiteCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "certificateName")

payload := sitecertificates.Certificate{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteCertificatesClient.CreateOrUpdateSlot`

```go
ctx := context.TODO()
id := sitecertificates.NewSlotCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "certificateName")

payload := sitecertificates.Certificate{
	// ...
}


read, err := client.CreateOrUpdateSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteCertificatesClient.Delete`

```go
ctx := context.TODO()
id := sitecertificates.NewSiteCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "certificateName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteCertificatesClient.DeleteSlot`

```go
ctx := context.TODO()
id := sitecertificates.NewSlotCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "certificateName")

read, err := client.DeleteSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteCertificatesClient.Get`

```go
ctx := context.TODO()
id := sitecertificates.NewSiteCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "certificateName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteCertificatesClient.GetSlot`

```go
ctx := context.TODO()
id := sitecertificates.NewSlotCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "certificateName")

read, err := client.GetSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteCertificatesClient.List`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteCertificatesClient.ListSlot`

```go
ctx := context.TODO()
id := sitecertificates.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.ListSlot(ctx, id)` can be used to do batched pagination
items, err := client.ListSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteCertificatesClient.Update`

```go
ctx := context.TODO()
id := sitecertificates.NewSiteCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "certificateName")

payload := sitecertificates.CertificatePatchResource{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteCertificatesClient.UpdateSlot`

```go
ctx := context.TODO()
id := sitecertificates.NewSlotCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "certificateName")

payload := sitecertificates.CertificatePatchResource{
	// ...
}


read, err := client.UpdateSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
