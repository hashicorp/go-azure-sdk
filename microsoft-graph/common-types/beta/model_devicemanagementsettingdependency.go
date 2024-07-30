package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingDependency struct {
	Constraints  *[]DeviceManagementConstraint `json:"constraints,omitempty"`
	DefinitionId *string                       `json:"definitionId,omitempty"`
	ODataType    *string                       `json:"@odata.type,omitempty"`
}
