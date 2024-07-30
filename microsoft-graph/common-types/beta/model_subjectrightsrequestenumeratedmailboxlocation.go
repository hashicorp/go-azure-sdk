package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestEnumeratedMailboxLocation struct {
	ODataType          *string   `json:"@odata.type,omitempty"`
	Upns               *[]string `json:"upns,omitempty"`
	UserPrincipalNames *[]string `json:"userPrincipalNames,omitempty"`
}
