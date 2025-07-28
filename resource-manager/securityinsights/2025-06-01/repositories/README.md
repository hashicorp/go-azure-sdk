
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2025-06-01/repositories` Documentation

The `repositories` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2025-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2025-06-01/repositories"
```


### Client Initialization

```go
client := repositories.NewRepositoriesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RepositoriesClient.SourceControllistRepositories`

```go
ctx := context.TODO()
id := repositories.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

payload := repositories.RepositoryAccessProperties{
	// ...
}


// alternatively `client.SourceControllistRepositories(ctx, id, payload)` can be used to do batched pagination
items, err := client.SourceControllistRepositoriesComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
