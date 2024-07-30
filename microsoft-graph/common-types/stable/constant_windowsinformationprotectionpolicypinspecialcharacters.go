package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionPolicyPinSpecialCharacters string

const (
	WindowsInformationProtectionPolicyPinSpecialCharacters_Allow             WindowsInformationProtectionPolicyPinSpecialCharacters = "allow"
	WindowsInformationProtectionPolicyPinSpecialCharacters_NotAllow          WindowsInformationProtectionPolicyPinSpecialCharacters = "notAllow"
	WindowsInformationProtectionPolicyPinSpecialCharacters_RequireAtLeastOne WindowsInformationProtectionPolicyPinSpecialCharacters = "requireAtLeastOne"
)

func PossibleValuesForWindowsInformationProtectionPolicyPinSpecialCharacters() []string {
	return []string{
		string(WindowsInformationProtectionPolicyPinSpecialCharacters_Allow),
		string(WindowsInformationProtectionPolicyPinSpecialCharacters_NotAllow),
		string(WindowsInformationProtectionPolicyPinSpecialCharacters_RequireAtLeastOne),
	}
}

func (s *WindowsInformationProtectionPolicyPinSpecialCharacters) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsInformationProtectionPolicyPinSpecialCharacters(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsInformationProtectionPolicyPinSpecialCharacters(input string) (*WindowsInformationProtectionPolicyPinSpecialCharacters, error) {
	vals := map[string]WindowsInformationProtectionPolicyPinSpecialCharacters{
		"allow":             WindowsInformationProtectionPolicyPinSpecialCharacters_Allow,
		"notallow":          WindowsInformationProtectionPolicyPinSpecialCharacters_NotAllow,
		"requireatleastone": WindowsInformationProtectionPolicyPinSpecialCharacters_RequireAtLeastOne,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsInformationProtectionPolicyPinSpecialCharacters(input)
	return &out, nil
}
