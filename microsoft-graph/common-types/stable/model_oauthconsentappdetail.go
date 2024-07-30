package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OAuthConsentAppDetail struct {
	AppScope    *OAuthConsentAppDetailAppScope `json:"appScope,omitempty"`
	DisplayLogo *string                        `json:"displayLogo,omitempty"`
	DisplayName *string                        `json:"displayName,omitempty"`
	ODataType   *string                        `json:"@odata.type,omitempty"`
}
