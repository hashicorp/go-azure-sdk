package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationPolicyTemplateReference struct {
	ODataType              *string                                                             `json:"@odata.type,omitempty"`
	TemplateDisplayName    *string                                                             `json:"templateDisplayName,omitempty"`
	TemplateDisplayVersion *string                                                             `json:"templateDisplayVersion,omitempty"`
	TemplateFamily         *DeviceManagementConfigurationPolicyTemplateReferenceTemplateFamily `json:"templateFamily,omitempty"`
	TemplateId             *string                                                             `json:"templateId,omitempty"`
}
