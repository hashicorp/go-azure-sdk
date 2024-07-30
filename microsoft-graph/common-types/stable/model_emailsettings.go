package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailSettings struct {
	ODataType          *string `json:"@odata.type,omitempty"`
	SenderDomain       *string `json:"senderDomain,omitempty"`
	UseCompanyBranding *bool   `json:"useCompanyBranding,omitempty"`
}
