
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2023-05-01/availableworkloadprofiles` Documentation

The `availableworkloadprofiles` SDK allows for interaction with the Azure Resource Manager Service `containerapps` (API Version `2023-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2023-05-01/availableworkloadprofiles"
```


### Client Initialization

```go
client := availableworkloadprofiles.NewAvailableWorkloadProfilesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AvailableWorkloadProfilesClient.Get`

```go
ctx := context.TODO()
id := availableworkloadprofiles.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

// alternatively `client.Get(ctx, id)` can be used to do batched pagination
items, err := client.GetComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
