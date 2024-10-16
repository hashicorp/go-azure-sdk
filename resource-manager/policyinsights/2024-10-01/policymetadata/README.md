
## `github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2024-10-01/policymetadata` Documentation

The `policymetadata` SDK allows for interaction with Azure Resource Manager `policyinsights` (API Version `2024-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2024-10-01/policymetadata"
```


### Client Initialization

```go
client := policymetadata.NewPolicyMetadataClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PolicyMetadataClient.GetResource`

```go
ctx := context.TODO()
id := policymetadata.NewPolicyMetadataID("policyMetadataName")

read, err := client.GetResource(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyMetadataClient.List`

```go
ctx := context.TODO()


// alternatively `client.List(ctx, policymetadata.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, policymetadata.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
