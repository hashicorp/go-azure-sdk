
## `github.com/hashicorp/go-azure-sdk/data-plane/batch/2022-01-01.15.0/certificates` Documentation

The `certificates` SDK allows for interaction with <unknown source data type> `batch` (API Version `2022-01-01.15.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/batch/2022-01-01.15.0/certificates"
```


### Client Initialization

```go
client := certificates.NewCertificatesClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `CertificatesClient.CertificateAdd`

```go
ctx := context.TODO()

payload := certificates.CertificateAddParameter{
	// ...
}


read, err := client.CertificateAdd(ctx, payload, certificates.DefaultCertificateAddOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CertificatesClient.CertificateCancelDeletion`

```go
ctx := context.TODO()
id := certificates.NewThumbprintID("thumbprintAlgorithmName", "thumbprintName")

read, err := client.CertificateCancelDeletion(ctx, id, certificates.DefaultCertificateCancelDeletionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CertificatesClient.CertificateDelete`

```go
ctx := context.TODO()
id := certificates.NewThumbprintID("thumbprintAlgorithmName", "thumbprintName")

read, err := client.CertificateDelete(ctx, id, certificates.DefaultCertificateDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CertificatesClient.CertificateGet`

```go
ctx := context.TODO()
id := certificates.NewThumbprintID("thumbprintAlgorithmName", "thumbprintName")

read, err := client.CertificateGet(ctx, id, certificates.DefaultCertificateGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CertificatesClient.CertificateList`

```go
ctx := context.TODO()


// alternatively `client.CertificateList(ctx, certificates.DefaultCertificateListOperationOptions())` can be used to do batched pagination
items, err := client.CertificateListComplete(ctx, certificates.DefaultCertificateListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
