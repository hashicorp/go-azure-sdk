
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2025-02-01/maintenanceconfigurations` Documentation

The `maintenanceconfigurations` SDK allows for interaction with Azure Resource Manager `containerservice` (API Version `2025-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2025-02-01/maintenanceconfigurations"
```


### Client Initialization

```go
client := maintenanceconfigurations.NewMaintenanceConfigurationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MaintenanceConfigurationsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := maintenanceconfigurations.NewMaintenanceConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterName", "maintenanceConfigurationName")

payload := maintenanceconfigurations.MaintenanceConfiguration{
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


### Example Usage: `MaintenanceConfigurationsClient.Delete`

```go
ctx := context.TODO()
id := maintenanceconfigurations.NewMaintenanceConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterName", "maintenanceConfigurationName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MaintenanceConfigurationsClient.Get`

```go
ctx := context.TODO()
id := maintenanceconfigurations.NewMaintenanceConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterName", "maintenanceConfigurationName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MaintenanceConfigurationsClient.ListByManagedCluster`

```go
ctx := context.TODO()
id := commonids.NewKubernetesClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterName")

// alternatively `client.ListByManagedCluster(ctx, id)` can be used to do batched pagination
items, err := client.ListByManagedClusterComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
