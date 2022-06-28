
## `github.com/hashicorp/go-azure-sdk/resource-manager/dns/2018-05-01/dns` Documentation

The `dns` SDK allows for interaction with the Azure Resource Manager Service `dns` (API Version `2018-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dns/2018-05-01/dns"
```


### Client Initialization

```go
client := dns.NewDNSClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `DNSClient.DnsResourceReferenceGetByTargetResources`

```go
ctx := context.TODO()
id := dns.NewSubscriptionID()

payload := dns.DnsResourceReferenceRequest{
	// ...
}

read, err := client.DnsResourceReferenceGetByTargetResources(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
