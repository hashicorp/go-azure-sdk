package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementBooleanSettingInstance struct {
	DefinitionId *string `json:"definitionId,omitempty"`
	Id           *string `json:"id,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
	Value        *bool   `json:"value,omitempty"`
	ValueJson    *string `json:"valueJson,omitempty"`
}
