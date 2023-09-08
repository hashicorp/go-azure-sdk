package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PasswordCredentialConfigurationRestrictionType string

const (
	PasswordCredentialConfigurationRestrictionTypecustomPasswordAddition PasswordCredentialConfigurationRestrictionType = "CustomPasswordAddition"
	PasswordCredentialConfigurationRestrictionTypepasswordAddition       PasswordCredentialConfigurationRestrictionType = "PasswordAddition"
	PasswordCredentialConfigurationRestrictionTypepasswordLifetime       PasswordCredentialConfigurationRestrictionType = "PasswordLifetime"
	PasswordCredentialConfigurationRestrictionTypesymmetricKeyAddition   PasswordCredentialConfigurationRestrictionType = "SymmetricKeyAddition"
	PasswordCredentialConfigurationRestrictionTypesymmetricKeyLifetime   PasswordCredentialConfigurationRestrictionType = "SymmetricKeyLifetime"
)

func PossibleValuesForPasswordCredentialConfigurationRestrictionType() []string {
	return []string{
		string(PasswordCredentialConfigurationRestrictionTypecustomPasswordAddition),
		string(PasswordCredentialConfigurationRestrictionTypepasswordAddition),
		string(PasswordCredentialConfigurationRestrictionTypepasswordLifetime),
		string(PasswordCredentialConfigurationRestrictionTypesymmetricKeyAddition),
		string(PasswordCredentialConfigurationRestrictionTypesymmetricKeyLifetime),
	}
}

func parsePasswordCredentialConfigurationRestrictionType(input string) (*PasswordCredentialConfigurationRestrictionType, error) {
	vals := map[string]PasswordCredentialConfigurationRestrictionType{
		"custompasswordaddition": PasswordCredentialConfigurationRestrictionTypecustomPasswordAddition,
		"passwordaddition":       PasswordCredentialConfigurationRestrictionTypepasswordAddition,
		"passwordlifetime":       PasswordCredentialConfigurationRestrictionTypepasswordLifetime,
		"symmetrickeyaddition":   PasswordCredentialConfigurationRestrictionTypesymmetricKeyAddition,
		"symmetrickeylifetime":   PasswordCredentialConfigurationRestrictionTypesymmetricKeyLifetime,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PasswordCredentialConfigurationRestrictionType(input)
	return &out, nil
}
