package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingComparison struct {
	ComparisonResult *DeviceManagementSettingComparisonComparisonResult `json:"comparisonResult,omitempty"`
	CurrentValueJson *string                                            `json:"currentValueJson,omitempty"`
	DefinitionId     *string                                            `json:"definitionId,omitempty"`
	DisplayName      *string                                            `json:"displayName,omitempty"`
	Id               *string                                            `json:"id,omitempty"`
	NewValueJson     *string                                            `json:"newValueJson,omitempty"`
	ODataType        *string                                            `json:"@odata.type,omitempty"`
}
