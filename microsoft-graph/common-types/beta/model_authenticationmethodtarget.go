package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodTarget struct {
	Id                     *string                               `json:"id,omitempty"`
	IsRegistrationRequired *bool                                 `json:"isRegistrationRequired,omitempty"`
	ODataType              *string                               `json:"@odata.type,omitempty"`
	TargetType             *AuthenticationMethodTargetTargetType `json:"targetType,omitempty"`
}
