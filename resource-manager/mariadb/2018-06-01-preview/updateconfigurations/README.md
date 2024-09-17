
## `github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/updateconfigurations` Documentation

The `updateconfigurations` SDK allows for interaction with Azure Resource Manager `mariadb` (API Version `2018-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/updateconfigurations"
```


### Client Initialization

```go
client := updateconfigurations.NewUpdateConfigurationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UpdateConfigurationsClient.ServerParametersListUpdateConfigurations`

```go
ctx := context.TODO()
id := updateconfigurations.NewServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue")

payload := updateconfigurations.ConfigurationListResult{
	// ...
}


if err := client.ServerParametersListUpdateConfigurationsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
