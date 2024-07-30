package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermsOfUseContainer struct {
	AgreementAcceptances *[]AgreementAcceptance `json:"agreementAcceptances,omitempty"`
	Agreements           *[]Agreement           `json:"agreements,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	ODataType            *string                `json:"@odata.type,omitempty"`
}
