
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/administrativeunits/beta/administrativeunit` Documentation

The `administrativeunit` SDK allows for interaction with the Azure Resource Manager Service `administrativeunits` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/administrativeunits/beta/administrativeunit"
```


### Client Initialization

```go
client := administrativeunit.NewAdministrativeUnitClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AdministrativeUnitClient.CheckAdministrativeUnitMemberGroup`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitIdValue")

payload := administrativeunit.CheckAdministrativeUnitMemberGroupRequest{
	// ...
}


read, err := client.CheckAdministrativeUnitMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrativeUnitClient.CheckAdministrativeUnitMemberObject`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitIdValue")

payload := administrativeunit.CheckAdministrativeUnitMemberObjectRequest{
	// ...
}


read, err := client.CheckAdministrativeUnitMemberObject(ctx, id, payload)
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


### Example Usage: `AdministrativeUnitClient.DeleteAdministrativeUnit`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitIdValue")

read, err := client.DeleteAdministrativeUnit(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrativeUnitClient.GetAdministrativeUnit`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitIdValue")

read, err := client.GetAdministrativeUnit(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrativeUnitClient.GetAdministrativeUnitMemberGroup`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitIdValue")

payload := administrativeunit.GetAdministrativeUnitMemberGroupRequest{
	// ...
}


read, err := client.GetAdministrativeUnitMemberGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrativeUnitClient.GetAdministrativeUnitMemberObject`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitIdValue")

payload := administrativeunit.GetAdministrativeUnitMemberObjectRequest{
	// ...
}


read, err := client.GetAdministrativeUnitMemberObject(ctx, id, payload)
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

payload := administrativeunit.GetAdministrativeUnitsByIdsRequest{
	// ...
}


// alternatively `client.GetAdministrativeUnitsByIds(ctx, payload)` can be used to do batched pagination
items, err := client.GetAdministrativeUnitsByIdsComplete(ctx, payload)
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


### Example Usage: `AdministrativeUnitClient.GetCount`

```go
ctx := context.TODO()


read, err := client.GetCount(ctx)
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


### Example Usage: `AdministrativeUnitClient.RestoreAdministrativeUnit`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitIdValue")

read, err := client.RestoreAdministrativeUnit(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrativeUnitClient.UpdateAdministrativeUnit`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitIdValue")

payload := administrativeunit.AdministrativeUnit{
	// ...
}


read, err := client.UpdateAdministrativeUnit(ctx, id, payload)
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
