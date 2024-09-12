
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

read, err := client.DeleteDepOnboardingSettingImportedAppleDeviceIdentity(ctx, id, deponboardingsettingimportedappledeviceidentity.DefaultDeleteDepOnboardingSettingImportedAppleDeviceIdentityOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DepOnboardingSettingImportedAppleDeviceIdentityClient.GetDepOnboardingSettingImportedAppleDeviceIdentitiesCount`

```go
ctx := context.TODO()
id := deponboardingsettingimportedappledeviceidentity.NewDeviceManagementDepOnboardingSettingID("depOnboardingSettingIdValue")

read, err := client.GetDepOnboardingSettingImportedAppleDeviceIdentitiesCount(ctx, id, deponboardingsettingimportedappledeviceidentity.DefaultGetDepOnboardingSettingImportedAppleDeviceIdentitiesCountOperationOptions())
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

read, err := client.GetDepOnboardingSettingImportedAppleDeviceIdentity(ctx, id, deponboardingsettingimportedappledeviceidentity.DefaultGetDepOnboardingSettingImportedAppleDeviceIdentityOperationOptions())
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

// alternatively `client.ListDepOnboardingSettingImportedAppleDeviceIdentities(ctx, id, deponboardingsettingimportedappledeviceidentity.DefaultListDepOnboardingSettingImportedAppleDeviceIdentitiesOperationOptions())` can be used to do batched pagination
items, err := client.ListDepOnboardingSettingImportedAppleDeviceIdentitiesComplete(ctx, id, deponboardingsettingimportedappledeviceidentity.DefaultListDepOnboardingSettingImportedAppleDeviceIdentitiesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DepOnboardingSettingImportedAppleDeviceIdentityClient.ListDepOnboardingSettingImportedAppleDeviceIdentityImportLists`

```go
ctx := context.TODO()
id := deponboardingsettingimportedappledeviceidentity.NewDeviceManagementDepOnboardingSettingID("depOnboardingSettingIdValue")

payload := deponboardingsettingimportedappledeviceidentity.ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsRequest{
	// ...
}


// alternatively `client.ListDepOnboardingSettingImportedAppleDeviceIdentityImportLists(ctx, id, payload, deponboardingsettingimportedappledeviceidentity.DefaultListDepOnboardingSettingImportedAppleDeviceIdentityImportListsOperationOptions())` can be used to do batched pagination
items, err := client.ListDepOnboardingSettingImportedAppleDeviceIdentityImportListsComplete(ctx, id, payload, deponboardingsettingimportedappledeviceidentity.DefaultListDepOnboardingSettingImportedAppleDeviceIdentityImportListsOperationOptions())
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
