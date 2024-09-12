
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


### Example Usage: `AdministrativeUnitClient.CheckMemberGroups`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitIdValue")

payload := administrativeunit.CheckMemberGroupsRequest{
	// ...
}


// alternatively `client.CheckMemberGroups(ctx, id, payload, administrativeunit.DefaultCheckMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberGroupsComplete(ctx, id, payload, administrativeunit.DefaultCheckMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AdministrativeUnitClient.CheckMemberObjects`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitIdValue")

payload := administrativeunit.CheckMemberObjectsRequest{
	// ...
}


// alternatively `client.CheckMemberObjects(ctx, id, payload, administrativeunit.DefaultCheckMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberObjectsComplete(ctx, id, payload, administrativeunit.DefaultCheckMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
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

read, err := client.DeleteAdministrativeUnit(ctx, id, administrativeunit.DefaultDeleteAdministrativeUnitOperationOptions())
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

read, err := client.GetAdministrativeUnit(ctx, id, administrativeunit.DefaultGetAdministrativeUnitOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrativeUnitClient.GetByIds`

```go
ctx := context.TODO()

payload := administrativeunit.GetByIdsRequest{
	// ...
}


// alternatively `client.GetByIds(ctx, payload, administrativeunit.DefaultGetByIdsOperationOptions())` can be used to do batched pagination
items, err := client.GetByIdsComplete(ctx, payload, administrativeunit.DefaultGetByIdsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AdministrativeUnitClient.GetCount`

```go
ctx := context.TODO()


read, err := client.GetCount(ctx, administrativeunit.DefaultGetCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrativeUnitClient.GetMemberGroups`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitIdValue")

payload := administrativeunit.GetMemberGroupsRequest{
	// ...
}


// alternatively `client.GetMemberGroups(ctx, id, payload, administrativeunit.DefaultGetMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberGroupsComplete(ctx, id, payload, administrativeunit.DefaultGetMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AdministrativeUnitClient.GetMemberObjects`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitIdValue")

payload := administrativeunit.GetMemberObjectsRequest{
	// ...
}


// alternatively `client.GetMemberObjects(ctx, id, payload, administrativeunit.DefaultGetMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberObjectsComplete(ctx, id, payload, administrativeunit.DefaultGetMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AdministrativeUnitClient.GetUserOwnedObject`

```go
ctx := context.TODO()

payload := administrativeunit.GetUserOwnedObjectRequest{
	// ...
}


read, err := client.GetUserOwnedObject(ctx, payload)
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


// alternatively `client.ListAdministrativeUnits(ctx, administrativeunit.DefaultListAdministrativeUnitsOperationOptions())` can be used to do batched pagination
items, err := client.ListAdministrativeUnitsComplete(ctx, administrativeunit.DefaultListAdministrativeUnitsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AdministrativeUnitClient.Restore`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitIdValue")

payload := administrativeunit.RestoreRequest{
	// ...
}


read, err := client.Restore(ctx, id, payload)
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


### Example Usage: `AdministrativeUnitClient.ValidateProperty`

```go
ctx := context.TODO()

payload := administrativeunit.ValidatePropertyRequest{
	// ...
}


read, err := client.ValidateProperty(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
