
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-02-01/jobdetails` Documentation

The `jobdetails` SDK allows for interaction with Azure Resource Manager `recoveryservicesbackup` (API Version `2023-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-02-01/jobdetails"
```


### Client Initialization

```go
client := jobdetails.NewJobDetailsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `JobDetailsClient.Get`

```go
ctx := context.TODO()
id := jobdetails.NewBackupJobID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "jobName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
