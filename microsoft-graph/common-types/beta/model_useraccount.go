package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAccount struct {
	DisplayName      *string            `json:"displayName,omitempty"`
	LastSeenDateTime *string            `json:"lastSeenDateTime,omitempty"`
	ODataType        *string            `json:"@odata.type,omitempty"`
	RiskScore        *string            `json:"riskScore,omitempty"`
	Service          *string            `json:"service,omitempty"`
	SigninName       *string            `json:"signinName,omitempty"`
	Status           *UserAccountStatus `json:"status,omitempty"`
}
