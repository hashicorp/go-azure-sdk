package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementIntegerSettingInstance struct {
	DefinitionId *string `json:"definitionId,omitempty"`
	Id           *string `json:"id,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
	Value        *int64  `json:"value,omitempty"`
	ValueJson    *string `json:"valueJson,omitempty"`
}
