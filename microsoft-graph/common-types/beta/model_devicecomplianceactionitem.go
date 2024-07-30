package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceActionItem struct {
	ActionType                *DeviceComplianceActionItemActionType `json:"actionType,omitempty"`
	GracePeriodHours          *int64                                `json:"gracePeriodHours,omitempty"`
	Id                        *string                               `json:"id,omitempty"`
	NotificationMessageCCList *[]string                             `json:"notificationMessageCCList,omitempty"`
	NotificationTemplateId    *string                               `json:"notificationTemplateId,omitempty"`
	ODataType                 *string                               `json:"@odata.type,omitempty"`
}
