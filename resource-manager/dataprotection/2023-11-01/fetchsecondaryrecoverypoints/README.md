
## `github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2023-11-01/fetchsecondaryrecoverypoints` Documentation

The `fetchsecondaryrecoverypoints` SDK allows for interaction with Azure Resource Manager `dataprotection` (API Version `2023-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2023-11-01/fetchsecondaryrecoverypoints"
```


### Client Initialization

```go
client := fetchsecondaryrecoverypoints.NewFetchSecondaryRecoveryPointsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FetchSecondaryRecoveryPointsClient.List`

```go
ctx := context.TODO()
id := fetchsecondaryrecoverypoints.NewProviderLocationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue")

payload := fetchsecondaryrecoverypoints.FetchSecondaryRPsRequestParameters{
	// ...
}


// alternatively `client.List(ctx, id, payload, fetchsecondaryrecoverypoints.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, payload, fetchsecondaryrecoverypoints.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
