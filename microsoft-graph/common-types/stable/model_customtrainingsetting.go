package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomTrainingSetting struct {
	AssignedTo        *CustomTrainingSettingAssignedTo  `json:"assignedTo,omitempty"`
	Description       *string                           `json:"description,omitempty"`
	DisplayName       *string                           `json:"displayName,omitempty"`
	DurationInMinutes *int64                            `json:"durationInMinutes,omitempty"`
	ODataType         *string                           `json:"@odata.type,omitempty"`
	SettingType       *CustomTrainingSettingSettingType `json:"settingType,omitempty"`
	Url               *string                           `json:"url,omitempty"`
}
