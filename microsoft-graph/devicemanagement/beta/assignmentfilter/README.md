
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/assignmentfilter` Documentation

The `assignmentfilter` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/assignmentfilter"
```


### Client Initialization

```go
client := assignmentfilter.NewAssignmentFilterClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AssignmentFilterClient.CreateAssignmentFilter`

```go
ctx := context.TODO()

payload := assignmentfilter.DeviceAndAppManagementAssignmentFilter{
	// ...
}


read, err := client.CreateAssignmentFilter(ctx, payload, assignmentfilter.DefaultCreateAssignmentFilterOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AssignmentFilterClient.DeleteAssignmentFilter`

```go
ctx := context.TODO()
id := assignmentfilter.NewDeviceManagementAssignmentFilterID("deviceAndAppManagementAssignmentFilterId")

read, err := client.DeleteAssignmentFilter(ctx, id, assignmentfilter.DefaultDeleteAssignmentFilterOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AssignmentFilterClient.EnableAssignmentFilters`

```go
ctx := context.TODO()

payload := assignmentfilter.EnableAssignmentFiltersRequest{
	// ...
}


read, err := client.EnableAssignmentFilters(ctx, payload, assignmentfilter.DefaultEnableAssignmentFiltersOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AssignmentFilterClient.GetAssignmentFilter`

```go
ctx := context.TODO()
id := assignmentfilter.NewDeviceManagementAssignmentFilterID("deviceAndAppManagementAssignmentFilterId")

read, err := client.GetAssignmentFilter(ctx, id, assignmentfilter.DefaultGetAssignmentFilterOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AssignmentFilterClient.GetAssignmentFiltersCount`

```go
ctx := context.TODO()


read, err := client.GetAssignmentFiltersCount(ctx, assignmentfilter.DefaultGetAssignmentFiltersCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AssignmentFilterClient.ListAssignmentFilters`

```go
ctx := context.TODO()


// alternatively `client.ListAssignmentFilters(ctx, assignmentfilter.DefaultListAssignmentFiltersOperationOptions())` can be used to do batched pagination
items, err := client.ListAssignmentFiltersComplete(ctx, assignmentfilter.DefaultListAssignmentFiltersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AssignmentFilterClient.UpdateAssignmentFilter`

```go
ctx := context.TODO()
id := assignmentfilter.NewDeviceManagementAssignmentFilterID("deviceAndAppManagementAssignmentFilterId")

payload := assignmentfilter.DeviceAndAppManagementAssignmentFilter{
	// ...
}


read, err := client.UpdateAssignmentFilter(ctx, id, payload, assignmentfilter.DefaultUpdateAssignmentFilterOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AssignmentFilterClient.ValidateAssignmentFiltersFilter`

```go
ctx := context.TODO()

payload := assignmentfilter.ValidateAssignmentFiltersFilterRequest{
	// ...
}


read, err := client.ValidateAssignmentFiltersFilter(ctx, payload, assignmentfilter.DefaultValidateAssignmentFiltersFilterOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
