package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationStrengthRoot struct {
	AuthenticationCombinations *AuthenticationStrengthRootAuthenticationCombinations `json:"authenticationCombinations,omitempty"`
	AuthenticationMethodModes  *[]AuthenticationMethodModeDetail                     `json:"authenticationMethodModes,omitempty"`
	Combinations               *AuthenticationStrengthRootCombinations               `json:"combinations,omitempty"`
	Id                         *string                                               `json:"id,omitempty"`
	ODataType                  *string                                               `json:"@odata.type,omitempty"`
	Policies                   *[]AuthenticationStrengthPolicy                       `json:"policies,omitempty"`
}
