package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SocialIdentitySource struct {
	DisplayName              *string                                       `json:"displayName,omitempty"`
	ODataType                *string                                       `json:"@odata.type,omitempty"`
	SocialIdentitySourceType *SocialIdentitySourceSocialIdentitySourceType `json:"socialIdentitySourceType,omitempty"`
}
