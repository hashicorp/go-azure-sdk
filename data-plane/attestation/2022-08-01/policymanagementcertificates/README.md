
## `github.com/hashicorp/go-azure-sdk/data-plane/attestation/2022-08-01/policymanagementcertificates` Documentation

The `policymanagementcertificates` SDK allows for interaction with <unknown source data type> `attestation` (API Version `2022-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/attestation/2022-08-01/policymanagementcertificates"
```


### Client Initialization

```go
client := policymanagementcertificates.NewPolicyManagementCertificatesClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `PolicyManagementCertificatesClient.PolicyCertificatesAdd`

```go
ctx := context.TODO()
var payload string

read, err := client.PolicyCertificatesAdd(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyManagementCertificatesClient.PolicyCertificatesGet`

```go
ctx := context.TODO()


read, err := client.PolicyCertificatesGet(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyManagementCertificatesClient.PolicyCertificatesRemove`

```go
ctx := context.TODO()
var payload string

read, err := client.PolicyCertificatesRemove(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
