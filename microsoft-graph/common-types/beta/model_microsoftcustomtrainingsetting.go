package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftCustomTrainingSetting struct {
	CompletionDateTime         *string                                                   `json:"completionDateTime,omitempty"`
	ODataType                  *string                                                   `json:"@odata.type,omitempty"`
	SettingType                *MicrosoftCustomTrainingSettingSettingType                `json:"settingType,omitempty"`
	TrainingAssignmentMappings *[]MicrosoftTrainingAssignmentMapping                     `json:"trainingAssignmentMappings,omitempty"`
	TrainingCompletionDuration *MicrosoftCustomTrainingSettingTrainingCompletionDuration `json:"trainingCompletionDuration,omitempty"`
}
