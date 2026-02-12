
## `github.com/hashicorp/go-azure-sdk/data-plane/attestation/2020-10-01/attestation` Documentation

The `attestation` SDK allows for interaction with <unknown source data type> `attestation` (API Version `2020-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/attestation/2020-10-01/attestation"
```


### Client Initialization

```go
client := attestation.NewAttestationClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `AttestationClient.AttestOpenEnclave`

```go
ctx := context.TODO()

payload := attestation.AttestOpenEnclaveRequest{
	// ...
}


read, err := client.AttestOpenEnclave(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AttestationClient.AttestSgxEnclave`

```go
ctx := context.TODO()

payload := attestation.AttestSgxEnclaveRequest{
	// ...
}


read, err := client.AttestSgxEnclave(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AttestationClient.AttestTpm`

```go
ctx := context.TODO()

payload := attestation.TpmAttestationRequest{
	// ...
}


read, err := client.AttestTpm(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
