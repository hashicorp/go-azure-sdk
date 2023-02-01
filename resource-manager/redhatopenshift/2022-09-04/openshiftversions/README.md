
## `github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2022-09-04/openshiftversions` Documentation

The `openshiftversions` SDK allows for interaction with the Azure Resource Manager Service `redhatopenshift` (API Version `2022-09-04`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2022-09-04/openshiftversions"
```


### Client Initialization

```go
client := openshiftversions.NewOpenShiftVersionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OpenShiftVersionsClient.List`

```go
ctx := context.TODO()
id := openshiftversions.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
