package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationIdentityCreationConfiguration struct {
	ODataType   *string                    `json:"@odata.type,omitempty"`
	UserDomains *[]EducationIdentityDomain `json:"userDomains,omitempty"`
}
