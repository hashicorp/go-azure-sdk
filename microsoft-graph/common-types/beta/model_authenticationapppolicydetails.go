package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAppPolicyDetails struct {
	AdminConfiguration       *AuthenticationAppPolicyDetailsAdminConfiguration       `json:"adminConfiguration,omitempty"`
	AuthenticationEvaluation *AuthenticationAppPolicyDetailsAuthenticationEvaluation `json:"authenticationEvaluation,omitempty"`
	ODataType                *string                                                 `json:"@odata.type,omitempty"`
	PolicyName               *string                                                 `json:"policyName,omitempty"`
	Status                   *AuthenticationAppPolicyDetailsStatus                   `json:"status,omitempty"`
}
