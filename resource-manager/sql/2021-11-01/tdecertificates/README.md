
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-11-01/tdecertificates` Documentation

The `tdecertificates` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2021-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-11-01/tdecertificates"
```


### Client Initialization

```go
client := tdecertificates.NewTdeCertificatesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TdeCertificatesClient.Create`

```go
ctx := context.TODO()
id := tdecertificates.NewServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue")

payload := tdecertificates.TdeCertificate{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
