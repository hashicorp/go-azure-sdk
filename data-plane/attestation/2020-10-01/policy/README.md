
## `github.com/hashicorp/go-azure-sdk/data-plane/attestation/2020-10-01/policy` Documentation

The `policy` SDK allows for interaction with <unknown source data type> `attestation` (API Version `2020-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/attestation/2020-10-01/policy"
```


### Client Initialization

```go
client := policy.NewPolicyClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `PolicyClient.Get`

```go
ctx := context.TODO()
id := policy.NewAttestationTypeID("OpenEnclave")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyClient.Reset`

```go
ctx := context.TODO()
var payload string

read, err := client.Reset(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyClient.Set`

```go
ctx := context.TODO()
id := policy.NewAttestationTypeID("OpenEnclave")
var payload string

read, err := client.Set(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
