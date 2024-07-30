package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessGuestsOrExternalUsers struct {
	ExternalTenants          *ConditionalAccessExternalTenants                               `json:"externalTenants,omitempty"`
	GuestOrExternalUserTypes *ConditionalAccessGuestsOrExternalUsersGuestOrExternalUserTypes `json:"guestOrExternalUserTypes,omitempty"`
	ODataType                *string                                                         `json:"@odata.type,omitempty"`
}
