package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAbstractComplexSettingInstance struct {
	DefinitionId     *string                            `json:"definitionId,omitempty"`
	Id               *string                            `json:"id,omitempty"`
	ImplementationId *string                            `json:"implementationId,omitempty"`
	ODataType        *string                            `json:"@odata.type,omitempty"`
	Value            *[]DeviceManagementSettingInstance `json:"value,omitempty"`
	ValueJson        *string                            `json:"valueJson,omitempty"`
}
