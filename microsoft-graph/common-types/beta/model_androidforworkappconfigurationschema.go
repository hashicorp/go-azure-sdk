package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkAppConfigurationSchema struct {
	ExampleJson *string                                     `json:"exampleJson,omitempty"`
	Id          *string                                     `json:"id,omitempty"`
	ODataType   *string                                     `json:"@odata.type,omitempty"`
	SchemaItems *[]AndroidForWorkAppConfigurationSchemaItem `json:"schemaItems,omitempty"`
}
