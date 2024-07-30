package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTemplateParameter struct {
	Description       *string                                   `json:"description,omitempty"`
	DisplayName       *string                                   `json:"displayName,omitempty"`
	JsonAllowedValues *string                                   `json:"jsonAllowedValues,omitempty"`
	JsonDefaultValue  *string                                   `json:"jsonDefaultValue,omitempty"`
	ODataType         *string                                   `json:"@odata.type,omitempty"`
	ValueType         *ManagedTenantsTemplateParameterValueType `json:"valueType,omitempty"`
}
