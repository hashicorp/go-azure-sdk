package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftManagedTrainingSetting struct {
	CompletionDateTime         *string                                                    `json:"completionDateTime,omitempty"`
	ODataType                  *string                                                    `json:"@odata.type,omitempty"`
	SettingType                *MicrosoftManagedTrainingSettingSettingType                `json:"settingType,omitempty"`
	TrainingCompletionDuration *MicrosoftManagedTrainingSettingTrainingCompletionDuration `json:"trainingCompletionDuration,omitempty"`
}
