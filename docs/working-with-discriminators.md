## Working with Discriminated Types

The Azure API does not makes use of Discriminated Types to allow a single API endpoint to provision multiple resources - which are modelled within the Swagger/OpenAPI Definitions as Discriminators.

A field within the Json payload determines which Implementation should be deserialized on both in both the Client (Go SDK) and the Server (Azure API) - as such this field must always be set to a fixed value.

This SDK outputs these as Go Interfaces (parent types) and Go Structs (for implementations) with custom Marshal and Unmarshal functions switching out the correct implementation as required.

For example the Interface (`Serialization`) and Implementations (`AvroSerialization` and `JsonSerialization`) are output in this SDK as:

```go
type Serialization interface {}
```

```go
var _ Serialization = AvroSerialization{}

type AvroSerialization struct {
	Properties *interface{} `json:"properties,omitempty"`

	// Fields inherited from Serialization
}
```

```go
var _ Serialization = JsonSerialization{}

type JsonSerialization struct {
	Properties *JsonSerializationProperties `json:"properties,omitempty"`

	// Fields inherited from Serialization
}
```

### Example: Using a Discriminator Implementation

Given the following Go model, using the Interface defined above:

```go
type Example struct {
	Serialization Serialization `json:"serialization"`
}
```

The desired implementation can be specified directly, for example:

```go
example := Example{
	Serialization: JsonSerialization{
		// ..
	},
}
```

### Example: Determining which Discriminator Implementation is returned

Given the following Go model, using the Interface defined above:

```go
type Example struct {
	Serialization Serialization `json:"serialization"`
}
```

The specific implementation can be determined via Type Assertion:

```go
// this snippet assumes that the type Example is returned/populated from the SDK
var response Example
if v, ok := response.Serialization.(inputs.AvroSerialization); ok {
	// do something with `v` which is an `AvroSerialization`
}
if v, ok := response.Serialization.(inputs.JsonSerialization); ok {
	// do something with `v` which is an `JsonSerialization`
}
```