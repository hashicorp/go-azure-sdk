
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/assignmentfilter` Documentation

The `assignmentfilter` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/assignmentfilter"
```


### Client Initialization

```go
client := assignmentfilter.NewAssignmentFilterClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AssignmentFilterClient.CreateAssignmentFilter`

```go
ctx := context.TODO()

payload := assignmentfilter.DeviceAndAppManagementAssignmentFilter{
	// ...
}


read, err := client.CreateAssignmentFilter(ctx, payload)
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
id := assignmentfilter.NewDeviceManagementAssignmentFilterID("deviceAndAppManagementAssignmentFilterIdValue")

read, err := client.DeleteAssignmentFilter(ctx, id, assignmentfilter.DefaultDeleteAssignmentFilterOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AssignmentFilterClient.EnableAssignmentFilter`

```go
ctx := context.TODO()

payload := assignmentfilter.EnableAssignmentFilterRequest{
	// ...
}


read, err := client.EnableAssignmentFilter(ctx, payload)
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
id := assignmentfilter.NewDeviceManagementAssignmentFilterID("deviceAndAppManagementAssignmentFilterIdValue")

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
id := assignmentfilter.NewDeviceManagementAssignmentFilterID("deviceAndAppManagementAssignmentFilterIdValue")

payload := assignmentfilter.DeviceAndAppManagementAssignmentFilter{
	// ...
}


read, err := client.UpdateAssignmentFilter(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AssignmentFilterClient.ValidateAssignmentFilter`

```go
ctx := context.TODO()

payload := assignmentfilter.ValidateAssignmentFilterRequest{
	// ...
}


read, err := client.ValidateAssignmentFilter(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
