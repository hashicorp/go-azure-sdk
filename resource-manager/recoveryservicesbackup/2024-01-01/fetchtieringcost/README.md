
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2024-01-01/fetchtieringcost` Documentation

The `fetchtieringcost` SDK allows for interaction with the Azure Resource Manager Service `recoveryservicesbackup` (API Version `2024-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2024-01-01/fetchtieringcost"
```


### Client Initialization

```go
client := fetchtieringcost.NewFetchTieringCostClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FetchTieringCostClient.Post`

```go
ctx := context.TODO()
id := fetchtieringcost.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue")

payload := fetchtieringcost.FetchTieringCostInfoRequest{
	// ...
}


if err := client.PostThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
