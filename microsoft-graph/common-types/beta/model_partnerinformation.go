package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerInformation struct {
	CommerceUrl       *string                        `json:"commerceUrl,omitempty"`
	CompanyName       *string                        `json:"companyName,omitempty"`
	CompanyType       *PartnerInformationCompanyType `json:"companyType,omitempty"`
	HelpUrl           *string                        `json:"helpUrl,omitempty"`
	ODataType         *string                        `json:"@odata.type,omitempty"`
	PartnerTenantId   *string                        `json:"partnerTenantId,omitempty"`
	SupportEmails     *[]string                      `json:"supportEmails,omitempty"`
	SupportTelephones *[]string                      `json:"supportTelephones,omitempty"`
	SupportUrl        *string                        `json:"supportUrl,omitempty"`
}
