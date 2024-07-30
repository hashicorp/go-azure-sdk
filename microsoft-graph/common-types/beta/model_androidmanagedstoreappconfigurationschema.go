package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAppConfigurationSchema struct {
	ExampleJson       *string                                          `json:"exampleJson,omitempty"`
	Id                *string                                          `json:"id,omitempty"`
	NestedSchemaItems *[]AndroidManagedStoreAppConfigurationSchemaItem `json:"nestedSchemaItems,omitempty"`
	ODataType         *string                                          `json:"@odata.type,omitempty"`
	SchemaItems       *[]AndroidManagedStoreAppConfigurationSchemaItem `json:"schemaItems,omitempty"`
}
