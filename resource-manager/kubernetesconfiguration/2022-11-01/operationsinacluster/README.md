
## `github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-11-01/operationsinacluster` Documentation

The `operationsinacluster` SDK allows for interaction with the Azure Resource Manager Service `kubernetesconfiguration` (API Version `2022-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-11-01/operationsinacluster"
```


### Client Initialization

```go
client := operationsinacluster.NewOperationsInAClusterClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OperationsInAClusterClient.OperationStatusList`

```go
ctx := context.TODO()
id := operationsinacluster.NewProviderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "providerValue", "clusterResourceValue", "clusterValue")

// alternatively `client.OperationStatusList(ctx, id)` can be used to do batched pagination
items, err := client.OperationStatusListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
