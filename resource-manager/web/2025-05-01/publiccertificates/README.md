
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/publiccertificates` Documentation

The `publiccertificates` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/publiccertificates"
```


### Client Initialization

```go
client := publiccertificates.NewPublicCertificatesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PublicCertificatesClient.WebAppsCreateOrUpdatePublicCertificate`

```go
ctx := context.TODO()
id := publiccertificates.NewPublicCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "publicCertificateName")

payload := publiccertificates.PublicCertificate{
	// ...
}


read, err := client.WebAppsCreateOrUpdatePublicCertificate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PublicCertificatesClient.WebAppsDeletePublicCertificate`

```go
ctx := context.TODO()
id := publiccertificates.NewPublicCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "publicCertificateName")

read, err := client.WebAppsDeletePublicCertificate(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PublicCertificatesClient.WebAppsGetPublicCertificate`

```go
ctx := context.TODO()
id := publiccertificates.NewPublicCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "publicCertificateName")

read, err := client.WebAppsGetPublicCertificate(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PublicCertificatesClient.WebAppsListPublicCertificates`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.WebAppsListPublicCertificates(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListPublicCertificatesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
