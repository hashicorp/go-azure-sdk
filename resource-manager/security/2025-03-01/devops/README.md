
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2025-03-01/devops` Documentation

The `devops` SDK allows for interaction with Azure Resource Manager `security` (API Version `2025-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2025-03-01/devops"
```


### Client Initialization

```go
client := devops.NewDevOpsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DevOpsClient.AzureDevOpsOrgsCreateOrUpdate`

```go
ctx := context.TODO()
id := devops.NewAzureDevOpsOrgID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName", "azureDevOpsOrgName")

payload := devops.AzureDevOpsOrg{
	// ...
}


if err := client.AzureDevOpsOrgsCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DevOpsClient.AzureDevOpsOrgsGet`

```go
ctx := context.TODO()
id := devops.NewAzureDevOpsOrgID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName", "azureDevOpsOrgName")

read, err := client.AzureDevOpsOrgsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DevOpsClient.AzureDevOpsOrgsList`

```go
ctx := context.TODO()
id := devops.NewSecurityConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName")

// alternatively `client.AzureDevOpsOrgsList(ctx, id)` can be used to do batched pagination
items, err := client.AzureDevOpsOrgsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DevOpsClient.AzureDevOpsOrgsListAvailable`

```go
ctx := context.TODO()
id := devops.NewSecurityConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName")

// alternatively `client.AzureDevOpsOrgsListAvailable(ctx, id)` can be used to do batched pagination
items, err := client.AzureDevOpsOrgsListAvailableComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DevOpsClient.AzureDevOpsOrgsUpdate`

```go
ctx := context.TODO()
id := devops.NewAzureDevOpsOrgID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName", "azureDevOpsOrgName")

payload := devops.AzureDevOpsOrg{
	// ...
}


if err := client.AzureDevOpsOrgsUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DevOpsClient.AzureDevOpsProjectsCreateOrUpdate`

```go
ctx := context.TODO()
id := devops.NewProjectID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName", "azureDevOpsOrgName", "projectName")

payload := devops.AzureDevOpsProject{
	// ...
}


if err := client.AzureDevOpsProjectsCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DevOpsClient.AzureDevOpsProjectsGet`

```go
ctx := context.TODO()
id := devops.NewProjectID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName", "azureDevOpsOrgName", "projectName")

read, err := client.AzureDevOpsProjectsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DevOpsClient.AzureDevOpsProjectsList`

```go
ctx := context.TODO()
id := devops.NewAzureDevOpsOrgID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName", "azureDevOpsOrgName")

// alternatively `client.AzureDevOpsProjectsList(ctx, id)` can be used to do batched pagination
items, err := client.AzureDevOpsProjectsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DevOpsClient.AzureDevOpsProjectsUpdate`

```go
ctx := context.TODO()
id := devops.NewProjectID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName", "azureDevOpsOrgName", "projectName")

payload := devops.AzureDevOpsProject{
	// ...
}


if err := client.AzureDevOpsProjectsUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DevOpsClient.AzureDevOpsReposCreateOrUpdate`

```go
ctx := context.TODO()
id := devops.NewProjectRepoID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName", "azureDevOpsOrgName", "projectName", "repoName")

payload := devops.AzureDevOpsRepository{
	// ...
}


if err := client.AzureDevOpsReposCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DevOpsClient.AzureDevOpsReposGet`

```go
ctx := context.TODO()
id := devops.NewProjectRepoID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName", "azureDevOpsOrgName", "projectName", "repoName")

read, err := client.AzureDevOpsReposGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DevOpsClient.AzureDevOpsReposList`

```go
ctx := context.TODO()
id := devops.NewProjectID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName", "azureDevOpsOrgName", "projectName")

// alternatively `client.AzureDevOpsReposList(ctx, id)` can be used to do batched pagination
items, err := client.AzureDevOpsReposListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DevOpsClient.AzureDevOpsReposUpdate`

```go
ctx := context.TODO()
id := devops.NewProjectRepoID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName", "azureDevOpsOrgName", "projectName", "repoName")

payload := devops.AzureDevOpsRepository{
	// ...
}


if err := client.AzureDevOpsReposUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DevOpsClient.ConfigurationsCreateOrUpdate`

```go
ctx := context.TODO()
id := devops.NewSecurityConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName")

payload := devops.DevOpsConfiguration{
	// ...
}


if err := client.ConfigurationsCreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DevOpsClient.ConfigurationsDelete`

```go
ctx := context.TODO()
id := devops.NewSecurityConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName")

if err := client.ConfigurationsDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DevOpsClient.ConfigurationsGet`

```go
ctx := context.TODO()
id := devops.NewSecurityConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName")

read, err := client.ConfigurationsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DevOpsClient.ConfigurationsList`

```go
ctx := context.TODO()
id := devops.NewSecurityConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName")

// alternatively `client.ConfigurationsList(ctx, id)` can be used to do batched pagination
items, err := client.ConfigurationsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DevOpsClient.ConfigurationsUpdate`

```go
ctx := context.TODO()
id := devops.NewSecurityConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName")

payload := devops.DevOpsConfiguration{
	// ...
}


if err := client.ConfigurationsUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DevOpsClient.GitHubOwnersGet`

```go
ctx := context.TODO()
id := devops.NewGitHubOwnerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName", "gitHubOwnerName")

read, err := client.GitHubOwnersGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DevOpsClient.GitHubOwnersList`

```go
ctx := context.TODO()
id := devops.NewSecurityConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName")

// alternatively `client.GitHubOwnersList(ctx, id)` can be used to do batched pagination
items, err := client.GitHubOwnersListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DevOpsClient.GitHubOwnersListAvailable`

```go
ctx := context.TODO()
id := devops.NewSecurityConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName")

// alternatively `client.GitHubOwnersListAvailable(ctx, id)` can be used to do batched pagination
items, err := client.GitHubOwnersListAvailableComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DevOpsClient.GitHubReposGet`

```go
ctx := context.TODO()
id := devops.NewRepoID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName", "gitHubOwnerName", "repoName")

read, err := client.GitHubReposGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DevOpsClient.GitHubReposList`

```go
ctx := context.TODO()
id := devops.NewGitHubOwnerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName", "gitHubOwnerName")

// alternatively `client.GitHubReposList(ctx, id)` can be used to do batched pagination
items, err := client.GitHubReposListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DevOpsClient.GitLabGroupsGet`

```go
ctx := context.TODO()
id := devops.NewGitLabGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName", "gitLabGroupName")

read, err := client.GitLabGroupsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DevOpsClient.GitLabGroupsList`

```go
ctx := context.TODO()
id := devops.NewSecurityConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName")

// alternatively `client.GitLabGroupsList(ctx, id)` can be used to do batched pagination
items, err := client.GitLabGroupsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DevOpsClient.GitLabGroupsListAvailable`

```go
ctx := context.TODO()
id := devops.NewSecurityConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName")

// alternatively `client.GitLabGroupsListAvailable(ctx, id)` can be used to do batched pagination
items, err := client.GitLabGroupsListAvailableComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DevOpsClient.GitLabProjectsGet`

```go
ctx := context.TODO()
id := devops.NewGitLabGroupProjectID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName", "gitLabGroupName", "projectName")

read, err := client.GitLabProjectsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DevOpsClient.GitLabProjectsList`

```go
ctx := context.TODO()
id := devops.NewGitLabGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName", "gitLabGroupName")

// alternatively `client.GitLabProjectsList(ctx, id)` can be used to do batched pagination
items, err := client.GitLabProjectsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DevOpsClient.GitLabSubgroupsList`

```go
ctx := context.TODO()
id := devops.NewGitLabGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "securityConnectorName", "gitLabGroupName")

// alternatively `client.GitLabSubgroupsList(ctx, id)` can be used to do batched pagination
items, err := client.GitLabSubgroupsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
