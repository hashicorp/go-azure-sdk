package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionStageSetting struct {
	CustomExtension *CustomCalloutExtension           `json:"customExtension,omitempty"`
	Id              *string                           `json:"id,omitempty"`
	ODataType       *string                           `json:"@odata.type,omitempty"`
	Stage           *CustomExtensionStageSettingStage `json:"stage,omitempty"`
}
