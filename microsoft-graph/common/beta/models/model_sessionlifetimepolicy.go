package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SessionLifetimePolicy struct {
	Detail                *string                                     `json:"detail,omitempty"`
	ExpirationRequirement *SessionLifetimePolicyExpirationRequirement `json:"expirationRequirement,omitempty"`
	ODataType             *string                                     `json:"@odata.type,omitempty"`
}
