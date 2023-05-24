
## `github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2023-05-01/commitmenttiers` Documentation

The `commitmenttiers` SDK allows for interaction with the Azure Resource Manager Service `cognitive` (API Version `2023-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2023-05-01/commitmenttiers"
```


### Client Initialization

```go
client := commitmenttiers.NewCommitmentTiersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CommitmentTiersClient.List`

```go
ctx := context.TODO()
id := commitmenttiers.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
