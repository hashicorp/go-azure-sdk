package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSetting struct {
	DisplayName *string         `json:"displayName,omitempty"`
	Id          *string         `json:"id,omitempty"`
	ODataType   *string         `json:"@odata.type,omitempty"`
	TemplateId  *string         `json:"templateId,omitempty"`
	Values      *[]SettingValue `json:"values,omitempty"`
}
