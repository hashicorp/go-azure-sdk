
## `github.com/hashicorp/go-azure-sdk/data-plane/attestation/2022-08-01/openidmetadatadiscovery` Documentation

The `openidmetadatadiscovery` SDK allows for interaction with <unknown source data type> `attestation` (API Version `2022-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/attestation/2022-08-01/openidmetadatadiscovery"
```


### Client Initialization

```go
client := openidmetadatadiscovery.NewOpenIDMetadataDiscoveryClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `OpenIDMetadataDiscoveryClient.MetadataConfigurationGet`

```go
ctx := context.TODO()


read, err := client.MetadataConfigurationGet(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenIDMetadataDiscoveryClient.SigningCertificatesGet`

```go
ctx := context.TODO()


read, err := client.SigningCertificatesGet(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
