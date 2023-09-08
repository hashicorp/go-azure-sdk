package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NoTrainingSetting struct {
	ODataType   *string                       `json:"@odata.type,omitempty"`
	SettingType *NoTrainingSettingSettingType `json:"settingType,omitempty"`
}
