package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestEnumeratedMailboxLocation struct {
	ODataType          *string   `json:"@odata.type,omitempty"`
	UserPrincipalNames *[]string `json:"userPrincipalNames,omitempty"`
}
