
## `github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/microsoftentraadministrators` Documentation

The `microsoftentraadministrators` SDK allows for interaction with Azure Resource Manager `horizondb` (API Version `2026-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/microsoftentraadministrators"
```


### Client Initialization

```go
client := microsoftentraadministrators.NewMicrosoftEntraAdministratorsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MicrosoftEntraAdministratorsClient.CreateOrReplace`

```go
ctx := context.TODO()
id := microsoftentraadministrators.NewMicrosoftEntraAdministratorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "objectId")

payload := microsoftentraadministrators.MicrosoftEntraAdministratorAdd{
	// ...
}


if err := client.CreateOrReplaceThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `MicrosoftEntraAdministratorsClient.Delete`

```go
ctx := context.TODO()
id := microsoftentraadministrators.NewMicrosoftEntraAdministratorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "objectId")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `MicrosoftEntraAdministratorsClient.Get`

```go
ctx := context.TODO()
id := microsoftentraadministrators.NewMicrosoftEntraAdministratorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "objectId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MicrosoftEntraAdministratorsClient.List`

```go
ctx := context.TODO()
id := microsoftentraadministrators.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
