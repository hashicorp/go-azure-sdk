package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationRequirementPolicy struct {
	Detail              *string                                             `json:"detail,omitempty"`
	ODataType           *string                                             `json:"@odata.type,omitempty"`
	RequirementProvider *AuthenticationRequirementPolicyRequirementProvider `json:"requirementProvider,omitempty"`
}
