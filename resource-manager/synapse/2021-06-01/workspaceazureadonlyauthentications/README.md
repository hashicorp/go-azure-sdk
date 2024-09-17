
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/workspaceazureadonlyauthentications` Documentation

The `workspaceazureadonlyauthentications` SDK allows for interaction with Azure Resource Manager `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/workspaceazureadonlyauthentications"
```


### Client Initialization

```go
client := workspaceazureadonlyauthentications.NewWorkspaceAzureADOnlyAuthenticationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkspaceAzureADOnlyAuthenticationsClient.AzureADOnlyAuthenticationsCreate`

```go
ctx := context.TODO()
id := workspaceazureadonlyauthentications.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

payload := workspaceazureadonlyauthentications.AzureADOnlyAuthentication{
	// ...
}


if err := client.AzureADOnlyAuthenticationsCreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkspaceAzureADOnlyAuthenticationsClient.AzureADOnlyAuthenticationsGet`

```go
ctx := context.TODO()
id := workspaceazureadonlyauthentications.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

read, err := client.AzureADOnlyAuthenticationsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceAzureADOnlyAuthenticationsClient.AzureADOnlyAuthenticationsList`

```go
ctx := context.TODO()
id := workspaceazureadonlyauthentications.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

// alternatively `client.AzureADOnlyAuthenticationsList(ctx, id)` can be used to do batched pagination
items, err := client.AzureADOnlyAuthenticationsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
