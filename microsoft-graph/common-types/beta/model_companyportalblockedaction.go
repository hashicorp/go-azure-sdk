package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompanyPortalBlockedAction struct {
	Action    *CompanyPortalBlockedActionAction    `json:"action,omitempty"`
	ODataType *string                              `json:"@odata.type,omitempty"`
	OwnerType *CompanyPortalBlockedActionOwnerType `json:"ownerType,omitempty"`
	Platform  *CompanyPortalBlockedActionPlatform  `json:"platform,omitempty"`
}
