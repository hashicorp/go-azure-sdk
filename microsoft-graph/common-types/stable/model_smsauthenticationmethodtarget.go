package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SmsAuthenticationMethodTarget struct {
	Id                     *string                                  `json:"id,omitempty"`
	IsRegistrationRequired *bool                                    `json:"isRegistrationRequired,omitempty"`
	IsUsableForSignIn      *bool                                    `json:"isUsableForSignIn,omitempty"`
	ODataType              *string                                  `json:"@odata.type,omitempty"`
	TargetType             *SmsAuthenticationMethodTargetTargetType `json:"targetType,omitempty"`
}
