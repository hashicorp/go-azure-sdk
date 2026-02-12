
## `github.com/hashicorp/go-azure-sdk/data-plane/synapse/2019-06-01-preview/datasets` Documentation

The `datasets` SDK allows for interaction with <unknown source data type> `synapse` (API Version `2019-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/synapse/2019-06-01-preview/datasets"
```


### Client Initialization

```go
client := datasets.NewDatasetsClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `DatasetsClient.DatasetCreateOrUpdateDataset`

```go
ctx := context.TODO()
id := datasets.NewDatasetID("datasetName")

payload := datasets.DatasetResource{
	// ...
}


if err := client.DatasetCreateOrUpdateDatasetThenPoll(ctx, id, payload, datasets.DefaultDatasetCreateOrUpdateDatasetOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `DatasetsClient.DatasetDeleteDataset`

```go
ctx := context.TODO()
id := datasets.NewDatasetID("datasetName")

if err := client.DatasetDeleteDatasetThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DatasetsClient.DatasetGetDataset`

```go
ctx := context.TODO()
id := datasets.NewDatasetID("datasetName")

read, err := client.DatasetGetDataset(ctx, id, datasets.DefaultDatasetGetDatasetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DatasetsClient.DatasetGetDatasetsByWorkspace`

```go
ctx := context.TODO()


// alternatively `client.DatasetGetDatasetsByWorkspace(ctx)` can be used to do batched pagination
items, err := client.DatasetGetDatasetsByWorkspaceComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DatasetsClient.DatasetRenameDataset`

```go
ctx := context.TODO()
id := datasets.NewDatasetID("datasetName")

payload := datasets.ArtifactRenameRequest{
	// ...
}


if err := client.DatasetRenameDatasetThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
