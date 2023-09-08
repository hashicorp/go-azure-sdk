package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationalUrl struct {
	LogoUrl             *string `json:"logoUrl,omitempty"`
	MarketingUrl        *string `json:"marketingUrl,omitempty"`
	ODataType           *string `json:"@odata.type,omitempty"`
	PrivacyStatementUrl *string `json:"privacyStatementUrl,omitempty"`
	SupportUrl          *string `json:"supportUrl,omitempty"`
	TermsOfServiceUrl   *string `json:"termsOfServiceUrl,omitempty"`
}
