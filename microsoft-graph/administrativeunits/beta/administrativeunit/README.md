
## `github.com/hashicorp/go-azure-sdk/resource-manager/administrativeunits/beta/administrativeunit` Documentation

The `administrativeunit` SDK allows for interaction with the Azure Resource Manager Service `administrativeunits` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/administrativeunits/beta/administrativeunit"
```


### Client Initialization

```go
client := administrativeunit.NewAdministrativeUnitClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AdministrativeUnitClient.CheckAdministrativeUnitByIdMemberGroup`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitIdValue")

payload := administrativeunit.CheckAdministrativeUnitByIdMemberGroupRequest{
	// ...
}


read, err := client.CheckAdministrativeUnitByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrativeUnitClient.CheckAdministrativeUnitByIdMemberObject`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitIdValue")

payload := administrativeunit.CheckAdministrativeUnitByIdMemberObjectRequest{
	// ...
}


read, err := client.CheckAdministrativeUnitByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrativeUnitClient.CreateAdministrativeUnit`

```go
ctx := context.TODO()

payload := administrativeunit.AdministrativeUnit{
	// ...
}


read, err := client.CreateAdministrativeUnit(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrativeUnitClient.DeleteAdministrativeUnitById`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitIdValue")

read, err := client.DeleteAdministrativeUnitById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrativeUnitClient.GetAdministrativeUnitById`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitIdValue")

read, err := client.GetAdministrativeUnitById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrativeUnitClient.GetAdministrativeUnitByIdMemberGroup`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitIdValue")

payload := administrativeunit.GetAdministrativeUnitByIdMemberGroupRequest{
	// ...
}


read, err := client.GetAdministrativeUnitByIdMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrativeUnitClient.GetAdministrativeUnitByIdMemberObject`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitIdValue")

payload := administrativeunit.GetAdministrativeUnitByIdMemberObjectRequest{
	// ...
}


read, err := client.GetAdministrativeUnitByIdMemberObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrativeUnitClient.GetAdministrativeUnitCount`

```go
ctx := context.TODO()


read, err := client.GetAdministrativeUnitCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrativeUnitClient.GetAdministrativeUnitsByIds`

```go
ctx := context.TODO()


// alternatively `client.GetAdministrativeUnitsByIds(ctx)` can be used to do batched pagination
items, err := client.GetAdministrativeUnitsByIdsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AdministrativeUnitClient.GetAdministrativeUnitsUserOwnedObject`

```go
ctx := context.TODO()

payload := administrativeunit.GetAdministrativeUnitsUserOwnedObjectRequest{
	// ...
}


read, err := client.GetAdministrativeUnitsUserOwnedObject(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrativeUnitClient.ListAdministrativeUnits`

```go
ctx := context.TODO()


// alternatively `client.ListAdministrativeUnits(ctx)` can be used to do batched pagination
items, err := client.ListAdministrativeUnitsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AdministrativeUnitClient.RestoreAdministrativeUnitById`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitIdValue")

read, err := client.RestoreAdministrativeUnitById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrativeUnitClient.UpdateAdministrativeUnitById`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitIdValue")

payload := administrativeunit.AdministrativeUnit{
	// ...
}


read, err := client.UpdateAdministrativeUnitById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrativeUnitClient.ValidateAdministrativeUnitsProperty`

```go
ctx := context.TODO()

payload := administrativeunit.ValidateAdministrativeUnitsPropertyRequest{
	// ...
}


read, err := client.ValidateAdministrativeUnitsProperty(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
