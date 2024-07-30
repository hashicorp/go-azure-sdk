
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deponboardingsettingimportedappledeviceidentity` Documentation

The `deponboardingsettingimportedappledeviceidentity` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/deponboardingsettingimportedappledeviceidentity"
```


### Client Initialization

```go
client := deponboardingsettingimportedappledeviceidentity.NewDepOnboardingSettingImportedAppleDeviceIdentityClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DepOnboardingSettingImportedAppleDeviceIdentityClient.CreateDepOnboardingSettingImportedAppleDeviceIdentity`

```go
ctx := context.TODO()
id := deponboardingsettingimportedappledeviceidentity.NewDeviceManagementDepOnboardingSettingID("depOnboardingSettingIdValue")

payload := deponboardingsettingimportedappledeviceidentity.ImportedAppleDeviceIdentity{
	// ...
}


read, err := client.CreateDepOnboardingSettingImportedAppleDeviceIdentity(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DepOnboardingSettingImportedAppleDeviceIdentityClient.DeleteDepOnboardingSettingImportedAppleDeviceIdentity`

```go
ctx := context.TODO()
id := deponboardingsettingimportedappledeviceidentity.NewDeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityID("depOnboardingSettingIdValue", "importedAppleDeviceIdentityIdValue")

read, err := client.DeleteDepOnboardingSettingImportedAppleDeviceIdentity(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DepOnboardingSettingImportedAppleDeviceIdentityClient.GetDepOnboardingSettingImportedAppleDeviceIdentity`

```go
ctx := context.TODO()
id := deponboardingsettingimportedappledeviceidentity.NewDeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityID("depOnboardingSettingIdValue", "importedAppleDeviceIdentityIdValue")

read, err := client.GetDepOnboardingSettingImportedAppleDeviceIdentity(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DepOnboardingSettingImportedAppleDeviceIdentityClient.GetDepOnboardingSettingImportedAppleDeviceIdentityCount`

```go
ctx := context.TODO()
id := deponboardingsettingimportedappledeviceidentity.NewDeviceManagementDepOnboardingSettingID("depOnboardingSettingIdValue")

read, err := client.GetDepOnboardingSettingImportedAppleDeviceIdentityCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DepOnboardingSettingImportedAppleDeviceIdentityClient.ListDepOnboardingSettingImportedAppleDeviceIdentities`

```go
ctx := context.TODO()
id := deponboardingsettingimportedappledeviceidentity.NewDeviceManagementDepOnboardingSettingID("depOnboardingSettingIdValue")

// alternatively `client.ListDepOnboardingSettingImportedAppleDeviceIdentities(ctx, id)` can be used to do batched pagination
items, err := client.ListDepOnboardingSettingImportedAppleDeviceIdentitiesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DepOnboardingSettingImportedAppleDeviceIdentityClient.ListDepOnboardingSettingImportedAppleDeviceIdentityImportAppleDeviceIdentityLists`

```go
ctx := context.TODO()
id := deponboardingsettingimportedappledeviceidentity.NewDeviceManagementDepOnboardingSettingID("depOnboardingSettingIdValue")

payload := deponboardingsettingimportedappledeviceidentity.ListDepOnboardingSettingImportedAppleDeviceIdentityImportAppleDeviceIdentityListsRequest{
	// ...
}


// alternatively `client.ListDepOnboardingSettingImportedAppleDeviceIdentityImportAppleDeviceIdentityLists(ctx, id, payload)` can be used to do batched pagination
items, err := client.ListDepOnboardingSettingImportedAppleDeviceIdentityImportAppleDeviceIdentityListsComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DepOnboardingSettingImportedAppleDeviceIdentityClient.UpdateDepOnboardingSettingImportedAppleDeviceIdentity`

```go
ctx := context.TODO()
id := deponboardingsettingimportedappledeviceidentity.NewDeviceManagementDepOnboardingSettingIdImportedAppleDeviceIdentityID("depOnboardingSettingIdValue", "importedAppleDeviceIdentityIdValue")

payload := deponboardingsettingimportedappledeviceidentity.ImportedAppleDeviceIdentity{
	// ...
}


read, err := client.UpdateDepOnboardingSettingImportedAppleDeviceIdentity(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
