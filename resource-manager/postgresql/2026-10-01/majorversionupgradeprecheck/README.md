
## `github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2026-10-01/majorversionupgradeprecheck` Documentation

The `majorversionupgradeprecheck` SDK allows for interaction with Azure Resource Manager `postgresql` (API Version `2026-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2026-10-01/majorversionupgradeprecheck"
```


### Client Initialization

```go
client := majorversionupgradeprecheck.NewMajorVersionUpgradePrecheckClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MajorVersionUpgradePrecheckClient.Get`

```go
ctx := context.TODO()
id := majorversionupgradeprecheck.NewMajorVersionUpgradePrecheckID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName", "precheckValidationId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MajorVersionUpgradePrecheckClient.List`

```go
ctx := context.TODO()
id := majorversionupgradeprecheck.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
