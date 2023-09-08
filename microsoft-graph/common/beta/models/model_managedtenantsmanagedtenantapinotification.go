package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagedTenantApiNotification struct {
	Alert              *ManagedTenantsManagedTenantAlert `json:"alert,omitempty"`
	CreatedByUserId    *string                           `json:"createdByUserId,omitempty"`
	CreatedDateTime    *string                           `json:"createdDateTime,omitempty"`
	Id                 *string                           `json:"id,omitempty"`
	IsAcknowledged     *bool                             `json:"isAcknowledged,omitempty"`
	LastActionByUserId *string                           `json:"lastActionByUserId,omitempty"`
	LastActionDateTime *string                           `json:"lastActionDateTime,omitempty"`
	Message            *string                           `json:"message,omitempty"`
	ODataType          *string                           `json:"@odata.type,omitempty"`
	Title              *string                           `json:"title,omitempty"`
	UserId             *string                           `json:"userId,omitempty"`
}
