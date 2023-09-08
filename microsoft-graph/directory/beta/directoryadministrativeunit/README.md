
## `github.com/hashicorp/go-azure-sdk/resource-manager/directory/beta/directoryadministrativeunit` Documentation

The `directoryadministrativeunit` SDK allows for interaction with the Azure Resource Manager Service `directory` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/directory/beta/directoryadministrativeunit"
```


### Client Initialization

```go
client := directoryadministrativeunit.NewDirectoryAdministrativeUnitClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DirectoryAdministrativeUnitClient.CheckDirectoryAdministrativeUnitByIdMemberGroup`

```go
ctx := context.TODO()
id := directoryadministrativeunit.NewDirectoryAdministrativeUnitID("administrativeUnitIdValue")

payload := directoryadministrativeunit.CheckDirectoryAdministrativeUnitByIdMemberGroupRequest{
	// ...
}


read, err := client.CheckDirectoryAdministrativeUnitByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryAdministrativeUnitClient.CheckDirectoryAdministrativeUnitByIdMemberObject`

```go
ctx := context.TODO()
id := directoryadministrativeunit.NewDirectoryAdministrativeUnitID("administrativeUnitIdValue")

payload := directoryadministrativeunit.CheckDirectoryAdministrativeUnitByIdMemberObjectRequest{
	// ...
}


read, err := client.CheckDirectoryAdministrativeUnitByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryAdministrativeUnitClient.CreateDirectoryAdministrativeUnit`

```go
ctx := context.TODO()

payload := directoryadministrativeunit.AdministrativeUnit{
	// ...
}


read, err := client.CreateDirectoryAdministrativeUnit(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryAdministrativeUnitClient.DeleteDirectoryAdministrativeUnitById`

```go
ctx := context.TODO()
id := directoryadministrativeunit.NewDirectoryAdministrativeUnitID("administrativeUnitIdValue")

read, err := client.DeleteDirectoryAdministrativeUnitById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryAdministrativeUnitClient.GetDirectoryAdministrativeUnitById`

```go
ctx := context.TODO()
id := directoryadministrativeunit.NewDirectoryAdministrativeUnitID("administrativeUnitIdValue")

read, err := client.GetDirectoryAdministrativeUnitById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryAdministrativeUnitClient.GetDirectoryAdministrativeUnitByIdMemberGroup`

```go
ctx := context.TODO()
id := directoryadministrativeunit.NewDirectoryAdministrativeUnitID("administrativeUnitIdValue")

payload := directoryadministrativeunit.GetDirectoryAdministrativeUnitByIdMemberGroupRequest{
	// ...
}


read, err := client.GetDirectoryAdministrativeUnitByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryAdministrativeUnitClient.GetDirectoryAdministrativeUnitByIdMemberObject`

```go
ctx := context.TODO()
id := directoryadministrativeunit.NewDirectoryAdministrativeUnitID("administrativeUnitIdValue")

payload := directoryadministrativeunit.GetDirectoryAdministrativeUnitByIdMemberObjectRequest{
	// ...
}


read, err := client.GetDirectoryAdministrativeUnitByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryAdministrativeUnitClient.GetDirectoryAdministrativeUnitCount`

```go
ctx := context.TODO()


read, err := client.GetDirectoryAdministrativeUnitCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryAdministrativeUnitClient.GetDirectoryAdministrativeUnitsByIds`

```go
ctx := context.TODO()


// alternatively `client.GetDirectoryAdministrativeUnitsByIds(ctx)` can be used to do batched pagination
items, err := client.GetDirectoryAdministrativeUnitsByIdsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryAdministrativeUnitClient.GetDirectoryAdministrativeUnitsUserOwnedObject`

```go
ctx := context.TODO()

payload := directoryadministrativeunit.GetDirectoryAdministrativeUnitsUserOwnedObjectRequest{
	// ...
}


read, err := client.GetDirectoryAdministrativeUnitsUserOwnedObject(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryAdministrativeUnitClient.ListDirectoryAdministrativeUnits`

```go
ctx := context.TODO()


// alternatively `client.ListDirectoryAdministrativeUnits(ctx)` can be used to do batched pagination
items, err := client.ListDirectoryAdministrativeUnitsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryAdministrativeUnitClient.RestoreDirectoryAdministrativeUnitById`

```go
ctx := context.TODO()
id := directoryadministrativeunit.NewDirectoryAdministrativeUnitID("administrativeUnitIdValue")

read, err := client.RestoreDirectoryAdministrativeUnitById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryAdministrativeUnitClient.UpdateDirectoryAdministrativeUnitById`

```go
ctx := context.TODO()
id := directoryadministrativeunit.NewDirectoryAdministrativeUnitID("administrativeUnitIdValue")

payload := directoryadministrativeunit.AdministrativeUnit{
	// ...
}


read, err := client.UpdateDirectoryAdministrativeUnitById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryAdministrativeUnitClient.ValidateDirectoryAdministrativeUnitsProperty`

```go
ctx := context.TODO()

payload := directoryadministrativeunit.ValidateDirectoryAdministrativeUnitsPropertyRequest{
	// ...
}


read, err := client.ValidateDirectoryAdministrativeUnitsProperty(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
