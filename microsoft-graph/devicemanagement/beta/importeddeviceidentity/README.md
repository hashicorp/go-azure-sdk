
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/importeddeviceidentity` Documentation

The `importeddeviceidentity` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/importeddeviceidentity"
```


### Client Initialization

```go
client := importeddeviceidentity.NewImportedDeviceIdentityClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ImportedDeviceIdentityClient.CreateImportedDeviceIdentity`

```go
ctx := context.TODO()

payload := importeddeviceidentity.ImportedDeviceIdentity{
	// ...
}


read, err := client.CreateImportedDeviceIdentity(ctx, payload)
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
id := importeddeviceidentity.NewDeviceManagementImportedDeviceIdentityID("importedDeviceIdentityIdValue")

read, err := client.DeleteImportedDeviceIdentity(ctx, id)
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
id := importeddeviceidentity.NewDeviceManagementImportedDeviceIdentityID("importedDeviceIdentityIdValue")

read, err := client.GetImportedDeviceIdentity(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ImportedDeviceIdentityClient.GetImportedDeviceIdentityCount`

```go
ctx := context.TODO()


read, err := client.GetImportedDeviceIdentityCount(ctx)
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


// alternatively `client.ListImportedDeviceIdentities(ctx)` can be used to do batched pagination
items, err := client.ListImportedDeviceIdentitiesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ImportedDeviceIdentityClient.ListImportedDeviceIdentityImportDeviceIdentityLists`

```go
ctx := context.TODO()

payload := importeddeviceidentity.ListImportedDeviceIdentityImportDeviceIdentityListsRequest{
	// ...
}


// alternatively `client.ListImportedDeviceIdentityImportDeviceIdentityLists(ctx, payload)` can be used to do batched pagination
items, err := client.ListImportedDeviceIdentityImportDeviceIdentityListsComplete(ctx, payload)
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


// alternatively `client.ListImportedDeviceIdentitySearchExistingIdentities(ctx, payload)` can be used to do batched pagination
items, err := client.ListImportedDeviceIdentitySearchExistingIdentitiesComplete(ctx, payload)
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
id := importeddeviceidentity.NewDeviceManagementImportedDeviceIdentityID("importedDeviceIdentityIdValue")

payload := importeddeviceidentity.ImportedDeviceIdentity{
	// ...
}


read, err := client.UpdateImportedDeviceIdentity(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
