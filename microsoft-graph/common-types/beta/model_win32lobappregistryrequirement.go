package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRegistryRequirement struct {
	Check32BitOn64System *bool                                        `json:"check32BitOn64System,omitempty"`
	DetectionType        *Win32LobAppRegistryRequirementDetectionType `json:"detectionType,omitempty"`
	DetectionValue       *string                                      `json:"detectionValue,omitempty"`
	KeyPath              *string                                      `json:"keyPath,omitempty"`
	ODataType            *string                                      `json:"@odata.type,omitempty"`
	Operator             *Win32LobAppRegistryRequirementOperator      `json:"operator,omitempty"`
	ValueName            *string                                      `json:"valueName,omitempty"`
}
