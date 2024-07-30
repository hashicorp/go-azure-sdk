package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyCredentialConfigurationRestrictionType string

const (
	KeyCredentialConfigurationRestrictionType_AsymmetricKeyLifetime KeyCredentialConfigurationRestrictionType = "asymmetricKeyLifetime"
)

func PossibleValuesForKeyCredentialConfigurationRestrictionType() []string {
	return []string{
		string(KeyCredentialConfigurationRestrictionType_AsymmetricKeyLifetime),
	}
}

func (s *KeyCredentialConfigurationRestrictionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKeyCredentialConfigurationRestrictionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKeyCredentialConfigurationRestrictionType(input string) (*KeyCredentialConfigurationRestrictionType, error) {
	vals := map[string]KeyCredentialConfigurationRestrictionType{
		"asymmetrickeylifetime": KeyCredentialConfigurationRestrictionType_AsymmetricKeyLifetime,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KeyCredentialConfigurationRestrictionType(input)
	return &out, nil
}
