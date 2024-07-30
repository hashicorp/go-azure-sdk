package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAppConfigurationSchemaItem struct {
	DataType                *AndroidManagedStoreAppConfigurationSchemaItemDataType `json:"dataType,omitempty"`
	DefaultBoolValue        *bool                                                  `json:"defaultBoolValue,omitempty"`
	DefaultIntValue         *int64                                                 `json:"defaultIntValue,omitempty"`
	DefaultStringArrayValue *[]string                                              `json:"defaultStringArrayValue,omitempty"`
	DefaultStringValue      *string                                                `json:"defaultStringValue,omitempty"`
	Description             *string                                                `json:"description,omitempty"`
	DisplayName             *string                                                `json:"displayName,omitempty"`
	Index                   *int64                                                 `json:"index,omitempty"`
	ODataType               *string                                                `json:"@odata.type,omitempty"`
	ParentIndex             *int64                                                 `json:"parentIndex,omitempty"`
	SchemaItemKey           *string                                                `json:"schemaItemKey,omitempty"`
	Selections              *[]KeyValuePair                                        `json:"selections,omitempty"`
}
