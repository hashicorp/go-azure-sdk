
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/tenantconfiguration` Documentation

The `tenantconfiguration` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2022-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/tenantconfiguration"
```


### Client Initialization

```go
client := tenantconfiguration.NewTenantConfigurationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TenantConfigurationClient.Deploy`

```go
ctx := context.TODO()
id := tenantconfiguration.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue")

payload := tenantconfiguration.DeployConfigurationParameters{
	// ...
}


if err := client.DeployThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `TenantConfigurationClient.Save`

```go
ctx := context.TODO()
id := tenantconfiguration.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue")

payload := tenantconfiguration.SaveConfigurationParameter{
	// ...
}


if err := client.SaveThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `TenantConfigurationClient.Validate`

```go
ctx := context.TODO()
id := tenantconfiguration.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue")

payload := tenantconfiguration.DeployConfigurationParameters{
	// ...
}


if err := client.ValidateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
