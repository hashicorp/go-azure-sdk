
## `github.com/hashicorp/go-azure-sdk/resource-manager/automanage/2022-05-04/configurationprofilesversions` Documentation

The `configurationprofilesversions` SDK allows for interaction with Azure Resource Manager `automanage` (API Version `2022-05-04`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/automanage/2022-05-04/configurationprofilesversions"
```


### Client Initialization

```go
client := configurationprofilesversions.NewConfigurationProfilesVersionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConfigurationProfilesVersionsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := configurationprofilesversions.NewConfigurationProfileVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "configurationProfileName", "versionName")

payload := configurationprofilesversions.ConfigurationProfile{
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


### Example Usage: `ConfigurationProfilesVersionsClient.Delete`

```go
ctx := context.TODO()
id := configurationprofilesversions.NewConfigurationProfileVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "configurationProfileName", "versionName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConfigurationProfilesVersionsClient.Get`

```go
ctx := context.TODO()
id := configurationprofilesversions.NewConfigurationProfileVersionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "configurationProfileName", "versionName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConfigurationProfilesVersionsClient.ListChildResources`

```go
ctx := context.TODO()
id := configurationprofilesversions.NewConfigurationProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "configurationProfileName")

read, err := client.ListChildResources(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
