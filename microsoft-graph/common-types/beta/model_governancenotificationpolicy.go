package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GovernanceNotificationPolicy struct {
	EnabledTemplateTypes  *[]string                         `json:"enabledTemplateTypes,omitempty"`
	NotificationTemplates *[]GovernanceNotificationTemplate `json:"notificationTemplates,omitempty"`
	ODataType             *string                           `json:"@odata.type,omitempty"`
}
