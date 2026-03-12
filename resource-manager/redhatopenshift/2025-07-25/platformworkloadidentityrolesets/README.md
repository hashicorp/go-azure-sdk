
## `github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2025-07-25/platformworkloadidentityrolesets` Documentation

The `platformworkloadidentityrolesets` SDK allows for interaction with Azure Resource Manager `redhatopenshift` (API Version `2025-07-25`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2025-07-25/platformworkloadidentityrolesets"
```


### Client Initialization

```go
client := platformworkloadidentityrolesets.NewPlatformWorkloadIdentityRoleSetsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PlatformWorkloadIdentityRoleSetsClient.List`

```go
ctx := context.TODO()
id := platformworkloadidentityrolesets.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PlatformWorkloadIdentityRoleSetsClient.PlatformWorkloadIdentityRoleSetGet`

```go
ctx := context.TODO()
id := platformworkloadidentityrolesets.NewPlatformWorkloadIdentityRoleSetID("12345678-1234-9876-4563-123456789012", "locationName", "platformWorkloadIdentityRoleSetName")

read, err := client.PlatformWorkloadIdentityRoleSetGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
