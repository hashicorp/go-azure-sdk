
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2023-03-01-preview/useridentity` Documentation

The `useridentity` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2023-03-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2023-03-01-preview/useridentity"
```


### Client Initialization

```go
client := useridentity.NewUserIdentityClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserIdentityClient.UserIdentitiesList`

```go
ctx := context.TODO()
id := useridentity.NewUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "userIdValue")

// alternatively `client.UserIdentitiesList(ctx, id)` can be used to do batched pagination
items, err := client.UserIdentitiesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
