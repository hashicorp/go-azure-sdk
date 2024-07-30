package templatemigratableto

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingInstance struct {
	DefinitionId *string `json:"definitionId,omitempty"`
	Id           *string `json:"id,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
	ValueJson    *string `json:"valueJson,omitempty"`
}
