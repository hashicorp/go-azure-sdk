package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FormsSettings struct {
	IsBingImageSearchEnabled            *bool   `json:"isBingImageSearchEnabled,omitempty"`
	IsExternalSendFormEnabled           *bool   `json:"isExternalSendFormEnabled,omitempty"`
	IsExternalShareCollaborationEnabled *bool   `json:"isExternalShareCollaborationEnabled,omitempty"`
	IsExternalShareResultEnabled        *bool   `json:"isExternalShareResultEnabled,omitempty"`
	IsExternalShareTemplateEnabled      *bool   `json:"isExternalShareTemplateEnabled,omitempty"`
	IsInOrgFormsPhishingScanEnabled     *bool   `json:"isInOrgFormsPhishingScanEnabled,omitempty"`
	IsRecordIdentityByDefaultEnabled    *bool   `json:"isRecordIdentityByDefaultEnabled,omitempty"`
	ODataType                           *string `json:"@odata.type,omitempty"`
}
