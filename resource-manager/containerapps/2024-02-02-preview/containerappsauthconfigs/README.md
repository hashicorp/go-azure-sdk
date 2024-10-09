
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/containerappsauthconfigs` Documentation

The `containerappsauthconfigs` SDK allows for interaction with Azure Resource Manager `containerapps` (API Version `2024-02-02-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/containerappsauthconfigs"
```


### Client Initialization

```go
client := containerappsauthconfigs.NewContainerAppsAuthConfigsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ContainerAppsAuthConfigsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := containerappsauthconfigs.NewAuthConfigID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "authConfigName")

payload := containerappsauthconfigs.AuthConfig{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContainerAppsAuthConfigsClient.Delete`

```go
ctx := context.TODO()
id := containerappsauthconfigs.NewAuthConfigID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "authConfigName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContainerAppsAuthConfigsClient.Get`

```go
ctx := context.TODO()
id := containerappsauthconfigs.NewAuthConfigID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "authConfigName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContainerAppsAuthConfigsClient.ListByContainerApp`

```go
ctx := context.TODO()
id := containerappsauthconfigs.NewContainerAppID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName")

// alternatively `client.ListByContainerApp(ctx, id)` can be used to do batched pagination
items, err := client.ListByContainerAppComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
