package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementIntentCustomizedSetting struct {
	CustomizedJson *string `json:"customizedJson,omitempty"`
	DefaultJson    *string `json:"defaultJson,omitempty"`
	DefinitionId   *string `json:"definitionId,omitempty"`
	ODataType      *string `json:"@odata.type,omitempty"`
}
