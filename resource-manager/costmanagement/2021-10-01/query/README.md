
## `github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2021-10-01/query` Documentation

The `query` SDK allows for interaction with the Azure Resource Manager Service `costmanagement` (API Version `2021-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/costmanagement/2021-10-01/query"
```


### Client Initialization

```go
client := query.NewQueryClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `QueryClient.Usage`

```go
ctx := context.TODO()
id := query.NewScopeID()

payload := query.QueryDefinition{
	// ...
}

read, err := client.Usage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `QueryClient.UsageByExternalCloudProviderType`

```go
ctx := context.TODO()
id := query.NewExternalCloudProviderTypeID("externalBillingAccounts", "externalCloudProviderIdValue")

payload := query.QueryDefinition{
	// ...
}

read, err := client.UsageByExternalCloudProviderType(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
