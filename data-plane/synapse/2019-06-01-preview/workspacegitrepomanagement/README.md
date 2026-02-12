
## `github.com/hashicorp/go-azure-sdk/data-plane/synapse/2019-06-01-preview/workspacegitrepomanagement` Documentation

The `workspacegitrepomanagement` SDK allows for interaction with <unknown source data type> `synapse` (API Version `2019-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/synapse/2019-06-01-preview/workspacegitrepomanagement"
```


### Client Initialization

```go
client := workspacegitrepomanagement.NewWorkspaceGitRepoManagementClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkspaceGitRepoManagementClient.GetGitHubAccessToken`

```go
ctx := context.TODO()

payload := workspacegitrepomanagement.GitHubAccessTokenRequest{
	// ...
}


read, err := client.GetGitHubAccessToken(ctx, payload, workspacegitrepomanagement.DefaultGetGitHubAccessTokenOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
