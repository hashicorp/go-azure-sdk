package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GovernanceSubject struct {
	DisplayName   *string `json:"displayName,omitempty"`
	Email         *string `json:"email,omitempty"`
	Id            *string `json:"id,omitempty"`
	ODataType     *string `json:"@odata.type,omitempty"`
	PrincipalName *string `json:"principalName,omitempty"`
	Type          *string `json:"type,omitempty"`
}
