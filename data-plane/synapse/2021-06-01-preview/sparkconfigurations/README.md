
## `github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/sparkconfigurations` Documentation

The `sparkconfigurations` SDK allows for interaction with <unknown source data type> `synapse` (API Version `2021-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/sparkconfigurations"
```


### Client Initialization

```go
client := sparkconfigurations.NewSparkConfigurationsClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `SparkConfigurationsClient.SparkConfigurationCreateOrUpdateSparkConfiguration`

```go
ctx := context.TODO()
id := sparkconfigurations.NewSparkConfigurationID("sparkConfigurationName")

payload := sparkconfigurations.SparkConfigurationResource{
	// ...
}


if err := client.SparkConfigurationCreateOrUpdateSparkConfigurationThenPoll(ctx, id, payload, sparkconfigurations.DefaultSparkConfigurationCreateOrUpdateSparkConfigurationOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `SparkConfigurationsClient.SparkConfigurationDeleteSparkConfiguration`

```go
ctx := context.TODO()
id := sparkconfigurations.NewSparkConfigurationID("sparkConfigurationName")

if err := client.SparkConfigurationDeleteSparkConfigurationThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `SparkConfigurationsClient.SparkConfigurationGetSparkConfiguration`

```go
ctx := context.TODO()
id := sparkconfigurations.NewSparkConfigurationID("sparkConfigurationName")

read, err := client.SparkConfigurationGetSparkConfiguration(ctx, id, sparkconfigurations.DefaultSparkConfigurationGetSparkConfigurationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SparkConfigurationsClient.SparkConfigurationGetSparkConfigurationsByWorkspace`

```go
ctx := context.TODO()


// alternatively `client.SparkConfigurationGetSparkConfigurationsByWorkspace(ctx)` can be used to do batched pagination
items, err := client.SparkConfigurationGetSparkConfigurationsByWorkspaceComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SparkConfigurationsClient.SparkConfigurationRenameSparkConfiguration`

```go
ctx := context.TODO()
id := sparkconfigurations.NewSparkConfigurationID("sparkConfigurationName")

payload := sparkconfigurations.ArtifactRenameRequest{
	// ...
}


if err := client.SparkConfigurationRenameSparkConfigurationThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
