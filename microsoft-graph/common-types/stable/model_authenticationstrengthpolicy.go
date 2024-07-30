package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationStrengthPolicy struct {
	AllowedCombinations       *AuthenticationStrengthPolicyAllowedCombinations   `json:"allowedCombinations,omitempty"`
	CombinationConfigurations *[]AuthenticationCombinationConfiguration          `json:"combinationConfigurations,omitempty"`
	CreatedDateTime           *string                                            `json:"createdDateTime,omitempty"`
	Description               *string                                            `json:"description,omitempty"`
	DisplayName               *string                                            `json:"displayName,omitempty"`
	Id                        *string                                            `json:"id,omitempty"`
	ModifiedDateTime          *string                                            `json:"modifiedDateTime,omitempty"`
	ODataType                 *string                                            `json:"@odata.type,omitempty"`
	PolicyType                *AuthenticationStrengthPolicyPolicyType            `json:"policyType,omitempty"`
	RequirementsSatisfied     *AuthenticationStrengthPolicyRequirementsSatisfied `json:"requirementsSatisfied,omitempty"`
}
