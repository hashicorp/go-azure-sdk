package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppSettings struct {
	AllowUserRequestsForAppAccess                     *bool   `json:"allowUserRequestsForAppAccess,omitempty"`
	Id                                                *string `json:"id,omitempty"`
	IsChatResourceSpecificConsentEnabled              *bool   `json:"isChatResourceSpecificConsentEnabled,omitempty"`
	IsUserPersonalScopeResourceSpecificConsentEnabled *bool   `json:"isUserPersonalScopeResourceSpecificConsentEnabled,omitempty"`
	ODataType                                         *string `json:"@odata.type,omitempty"`
}
