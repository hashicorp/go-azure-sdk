package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomerVoiceSettings struct {
	IsInOrgFormsPhishingScanEnabled  *bool   `json:"isInOrgFormsPhishingScanEnabled,omitempty"`
	IsRecordIdentityByDefaultEnabled *bool   `json:"isRecordIdentityByDefaultEnabled,omitempty"`
	IsRestrictedSurveyAccessEnabled  *bool   `json:"isRestrictedSurveyAccessEnabled,omitempty"`
	ODataType                        *string `json:"@odata.type,omitempty"`
}
