package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataSharingConsent struct {
	GrantDateTime      *string `json:"grantDateTime,omitempty"`
	Granted            *bool   `json:"granted,omitempty"`
	GrantedByUpn       *string `json:"grantedByUpn,omitempty"`
	GrantedByUserId    *string `json:"grantedByUserId,omitempty"`
	Id                 *string `json:"id,omitempty"`
	ODataType          *string `json:"@odata.type,omitempty"`
	ServiceDisplayName *string `json:"serviceDisplayName,omitempty"`
	TermsUrl           *string `json:"termsUrl,omitempty"`
}
