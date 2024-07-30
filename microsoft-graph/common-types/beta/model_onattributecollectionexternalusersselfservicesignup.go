package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnAttributeCollectionExternalUsersSelfServiceSignUp struct {
	AttributeCollectionPage *AuthenticationAttributeCollectionPage `json:"attributeCollectionPage,omitempty"`
	Attributes              *[]IdentityUserFlowAttribute           `json:"attributes,omitempty"`
	ODataType               *string                                `json:"@odata.type,omitempty"`
}
