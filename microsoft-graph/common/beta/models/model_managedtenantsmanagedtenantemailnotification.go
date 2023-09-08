package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagedTenantEmailNotification struct {
	Alert              *ManagedTenantsManagedTenantAlert `json:"alert,omitempty"`
	CreatedByUserId    *string                           `json:"createdByUserId,omitempty"`
	CreatedDateTime    *string                           `json:"createdDateTime,omitempty"`
	EmailAddresses     *[]ManagedTenantsEmail            `json:"emailAddresses,omitempty"`
	EmailBody          *string                           `json:"emailBody,omitempty"`
	Id                 *string                           `json:"id,omitempty"`
	LastActionByUserId *string                           `json:"lastActionByUserId,omitempty"`
	LastActionDateTime *string                           `json:"lastActionDateTime,omitempty"`
	ODataType          *string                           `json:"@odata.type,omitempty"`
	Subject            *string                           `json:"subject,omitempty"`
}
