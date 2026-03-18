
## `github.com/hashicorp/go-azure-sdk/data-plane/keyvault/2025-07-01/administrations` Documentation

The `administrations` SDK allows for interaction with <unknown source data type> `keyvault` (API Version `2025-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/keyvault/2025-07-01/administrations"
```


### Client Initialization

```go
client := administrations.NewAdministrationsClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `AdministrationsClient.FullBackup`

```go
ctx := context.TODO()

payload := administrations.SASTokenParameter{
	// ...
}


if err := client.FullBackupThenPoll(ctx, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AdministrationsClient.FullBackupStatus`

```go
ctx := context.TODO()
id := administrations.NewBackupID("jobId")

read, err := client.FullBackupStatus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrationsClient.FullRestoreOperation`

```go
ctx := context.TODO()

payload := administrations.RestoreOperationParameters{
	// ...
}


if err := client.FullRestoreOperationThenPoll(ctx, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AdministrationsClient.GetSetting`

```go
ctx := context.TODO()
id := administrations.NewSettingID("settingName")

read, err := client.GetSetting(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrationsClient.GetSettings`

```go
ctx := context.TODO()


read, err := client.GetSettings(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrationsClient.PreFullBackup`

```go
ctx := context.TODO()

payload := administrations.PreBackupOperationParameters{
	// ...
}


if err := client.PreFullBackupThenPoll(ctx, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AdministrationsClient.PreFullRestoreOperation`

```go
ctx := context.TODO()

payload := administrations.PreRestoreOperationParameters{
	// ...
}


if err := client.PreFullRestoreOperationThenPoll(ctx, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AdministrationsClient.RestoreStatus`

```go
ctx := context.TODO()
id := administrations.NewRestoreID("jobId")

read, err := client.RestoreStatus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrationsClient.RoleAssignmentsCreate`

```go
ctx := context.TODO()
id := administrations.NewScopedRoleAssignmentID("https://endpoint-url.example.com", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

payload := administrations.RoleAssignmentCreateParameters{
	// ...
}


read, err := client.RoleAssignmentsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrationsClient.RoleAssignmentsDelete`

```go
ctx := context.TODO()
id := administrations.NewScopedRoleAssignmentID("https://endpoint-url.example.com", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.RoleAssignmentsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrationsClient.RoleAssignmentsGet`

```go
ctx := context.TODO()
id := administrations.NewScopedRoleAssignmentID("https://endpoint-url.example.com", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.RoleAssignmentsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrationsClient.RoleAssignmentsListForScope`

```go
ctx := context.TODO()
id := administrations.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.RoleAssignmentsListForScope(ctx, id, administrations.DefaultRoleAssignmentsListForScopeOperationOptions())` can be used to do batched pagination
items, err := client.RoleAssignmentsListForScopeComplete(ctx, id, administrations.DefaultRoleAssignmentsListForScopeOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AdministrationsClient.RoleDefinitionsCreateOrUpdate`

```go
ctx := context.TODO()
id := administrations.NewScopedRoleDefinitionID("https://endpoint-url.example.com", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

payload := administrations.RoleDefinitionCreateParameters{
	// ...
}


read, err := client.RoleDefinitionsCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrationsClient.RoleDefinitionsDelete`

```go
ctx := context.TODO()
id := administrations.NewScopedRoleDefinitionID("https://endpoint-url.example.com", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.RoleDefinitionsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrationsClient.RoleDefinitionsGet`

```go
ctx := context.TODO()
id := administrations.NewScopedRoleDefinitionID("https://endpoint-url.example.com", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.RoleDefinitionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrationsClient.RoleDefinitionsList`

```go
ctx := context.TODO()
id := administrations.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.RoleDefinitionsList(ctx, id, administrations.DefaultRoleDefinitionsListOperationOptions())` can be used to do batched pagination
items, err := client.RoleDefinitionsListComplete(ctx, id, administrations.DefaultRoleDefinitionsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AdministrationsClient.SelectiveKeyRestoreOperation`

```go
ctx := context.TODO()
id := administrations.NewKeyID("keyName")

payload := administrations.SelectiveKeyRestoreOperationParameters{
	// ...
}


if err := client.SelectiveKeyRestoreOperationThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AdministrationsClient.UpdateSetting`

```go
ctx := context.TODO()
id := administrations.NewSettingID("settingName")

payload := administrations.UpdateSettingRequest{
	// ...
}


read, err := client.UpdateSetting(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
