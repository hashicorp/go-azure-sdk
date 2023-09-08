package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagedTenantAlertLog struct {
	Alert              *ManagedTenantsManagedTenantAlert `json:"alert,omitempty"`
	Content            *ManagedTenantsAlertLogContent    `json:"content,omitempty"`
	CreatedByUserId    *string                           `json:"createdByUserId,omitempty"`
	CreatedDateTime    *string                           `json:"createdDateTime,omitempty"`
	Id                 *string                           `json:"id,omitempty"`
	LastActionByUserId *string                           `json:"lastActionByUserId,omitempty"`
	LastActionDateTime *string                           `json:"lastActionDateTime,omitempty"`
	ODataType          *string                           `json:"@odata.type,omitempty"`
}
