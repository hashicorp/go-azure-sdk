
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/importeddeviceidentity` Documentation

The `importeddeviceidentity` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/importeddeviceidentity"
```


### Client Initialization

```go
client := importeddeviceidentity.NewImportedDeviceIdentityClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ImportedDeviceIdentityClient.CreateImportedDeviceIdentity`

```go
ctx := context.TODO()

payload := importeddeviceidentity.ImportedDeviceIdentity{
	// ...
}


read, err := client.CreateImportedDeviceIdentity(ctx, payload, importeddeviceidentity.DefaultCreateImportedDeviceIdentityOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ImportedDeviceIdentityClient.DeleteImportedDeviceIdentity`

```go
ctx := context.TODO()
id := importeddeviceidentity.NewDeviceManagementImportedDeviceIdentityID("importedDeviceIdentityId")

read, err := client.DeleteImportedDeviceIdentity(ctx, id, importeddeviceidentity.DefaultDeleteImportedDeviceIdentityOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ImportedDeviceIdentityClient.GetImportedDeviceIdentitiesCount`

```go
ctx := context.TODO()


read, err := client.GetImportedDeviceIdentitiesCount(ctx, importeddeviceidentity.DefaultGetImportedDeviceIdentitiesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ImportedDeviceIdentityClient.GetImportedDeviceIdentity`

```go
ctx := context.TODO()
id := importeddeviceidentity.NewDeviceManagementImportedDeviceIdentityID("importedDeviceIdentityId")

read, err := client.GetImportedDeviceIdentity(ctx, id, importeddeviceidentity.DefaultGetImportedDeviceIdentityOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ImportedDeviceIdentityClient.ListImportedDeviceIdentities`

```go
ctx := context.TODO()


// alternatively `client.ListImportedDeviceIdentities(ctx, importeddeviceidentity.DefaultListImportedDeviceIdentitiesOperationOptions())` can be used to do batched pagination
items, err := client.ListImportedDeviceIdentitiesComplete(ctx, importeddeviceidentity.DefaultListImportedDeviceIdentitiesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ImportedDeviceIdentityClient.ListImportedDeviceIdentityImportLists`

```go
ctx := context.TODO()

payload := importeddeviceidentity.ListImportedDeviceIdentityImportListsRequest{
	// ...
}


// alternatively `client.ListImportedDeviceIdentityImportLists(ctx, payload, importeddeviceidentity.DefaultListImportedDeviceIdentityImportListsOperationOptions())` can be used to do batched pagination
items, err := client.ListImportedDeviceIdentityImportListsComplete(ctx, payload, importeddeviceidentity.DefaultListImportedDeviceIdentityImportListsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ImportedDeviceIdentityClient.ListImportedDeviceIdentitySearchExistingIdentities`

```go
ctx := context.TODO()

payload := importeddeviceidentity.ListImportedDeviceIdentitySearchExistingIdentitiesRequest{
	// ...
}


// alternatively `client.ListImportedDeviceIdentitySearchExistingIdentities(ctx, payload, importeddeviceidentity.DefaultListImportedDeviceIdentitySearchExistingIdentitiesOperationOptions())` can be used to do batched pagination
items, err := client.ListImportedDeviceIdentitySearchExistingIdentitiesComplete(ctx, payload, importeddeviceidentity.DefaultListImportedDeviceIdentitySearchExistingIdentitiesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ImportedDeviceIdentityClient.UpdateImportedDeviceIdentity`

```go
ctx := context.TODO()
id := importeddeviceidentity.NewDeviceManagementImportedDeviceIdentityID("importedDeviceIdentityId")

payload := importeddeviceidentity.ImportedDeviceIdentity{
	// ...
}


read, err := client.UpdateImportedDeviceIdentity(ctx, id, payload, importeddeviceidentity.DefaultUpdateImportedDeviceIdentityOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
