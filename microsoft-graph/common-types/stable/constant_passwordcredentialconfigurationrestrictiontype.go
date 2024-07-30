package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PasswordCredentialConfigurationRestrictionType string

const (
	PasswordCredentialConfigurationRestrictionType_CustomPasswordAddition PasswordCredentialConfigurationRestrictionType = "customPasswordAddition"
	PasswordCredentialConfigurationRestrictionType_PasswordAddition       PasswordCredentialConfigurationRestrictionType = "passwordAddition"
	PasswordCredentialConfigurationRestrictionType_PasswordLifetime       PasswordCredentialConfigurationRestrictionType = "passwordLifetime"
	PasswordCredentialConfigurationRestrictionType_SymmetricKeyAddition   PasswordCredentialConfigurationRestrictionType = "symmetricKeyAddition"
	PasswordCredentialConfigurationRestrictionType_SymmetricKeyLifetime   PasswordCredentialConfigurationRestrictionType = "symmetricKeyLifetime"
)

func PossibleValuesForPasswordCredentialConfigurationRestrictionType() []string {
	return []string{
		string(PasswordCredentialConfigurationRestrictionType_CustomPasswordAddition),
		string(PasswordCredentialConfigurationRestrictionType_PasswordAddition),
		string(PasswordCredentialConfigurationRestrictionType_PasswordLifetime),
		string(PasswordCredentialConfigurationRestrictionType_SymmetricKeyAddition),
		string(PasswordCredentialConfigurationRestrictionType_SymmetricKeyLifetime),
	}
}

func (s *PasswordCredentialConfigurationRestrictionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePasswordCredentialConfigurationRestrictionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePasswordCredentialConfigurationRestrictionType(input string) (*PasswordCredentialConfigurationRestrictionType, error) {
	vals := map[string]PasswordCredentialConfigurationRestrictionType{
		"custompasswordaddition": PasswordCredentialConfigurationRestrictionType_CustomPasswordAddition,
		"passwordaddition":       PasswordCredentialConfigurationRestrictionType_PasswordAddition,
		"passwordlifetime":       PasswordCredentialConfigurationRestrictionType_PasswordLifetime,
		"symmetrickeyaddition":   PasswordCredentialConfigurationRestrictionType_SymmetricKeyAddition,
		"symmetrickeylifetime":   PasswordCredentialConfigurationRestrictionType_SymmetricKeyLifetime,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PasswordCredentialConfigurationRestrictionType(input)
	return &out, nil
}
