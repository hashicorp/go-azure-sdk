package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftTrainingAssignmentMapping struct {
	AssignedTo  *[]MicrosoftTrainingAssignmentMappingAssignedTo `json:"assignedTo,omitempty"`
	ODataType   *string                                         `json:"@odata.type,omitempty"`
	SettingType *MicrosoftTrainingAssignmentMappingSettingType  `json:"settingType,omitempty"`
	Training    *Training                                       `json:"training,omitempty"`
}
