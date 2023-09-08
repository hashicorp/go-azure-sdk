package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagedTenantTicketingEndpoint struct {
	CreatedByUserId    *string `json:"createdByUserId,omitempty"`
	CreatedDateTime    *string `json:"createdDateTime,omitempty"`
	DisplayName        *string `json:"displayName,omitempty"`
	EmailAddress       *string `json:"emailAddress,omitempty"`
	Id                 *string `json:"id,omitempty"`
	LastActionByUserId *string `json:"lastActionByUserId,omitempty"`
	LastActionDateTime *string `json:"lastActionDateTime,omitempty"`
	ODataType          *string `json:"@odata.type,omitempty"`
	PhoneNumber        *string `json:"phoneNumber,omitempty"`
}
