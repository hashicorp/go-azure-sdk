
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/publiccertificateoperationgroup` Documentation

The `publiccertificateoperationgroup` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/publiccertificateoperationgroup"
```


### Client Initialization

```go
client := publiccertificateoperationgroup.NewPublicCertificateOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PublicCertificateOperationGroupClient.WebAppsCreateOrUpdatePublicCertificateSlot`

```go
ctx := context.TODO()
id := publiccertificateoperationgroup.NewSlotPublicCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "publicCertificateName")

payload := publiccertificateoperationgroup.PublicCertificate{
	// ...
}


read, err := client.WebAppsCreateOrUpdatePublicCertificateSlot(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PublicCertificateOperationGroupClient.WebAppsDeletePublicCertificateSlot`

```go
ctx := context.TODO()
id := publiccertificateoperationgroup.NewSlotPublicCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "publicCertificateName")

read, err := client.WebAppsDeletePublicCertificateSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PublicCertificateOperationGroupClient.WebAppsGetPublicCertificateSlot`

```go
ctx := context.TODO()
id := publiccertificateoperationgroup.NewSlotPublicCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName", "publicCertificateName")

read, err := client.WebAppsGetPublicCertificateSlot(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PublicCertificateOperationGroupClient.WebAppsListPublicCertificatesSlot`

```go
ctx := context.TODO()
id := publiccertificateoperationgroup.NewSlotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "slotName")

// alternatively `client.WebAppsListPublicCertificatesSlot(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListPublicCertificatesSlotComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
